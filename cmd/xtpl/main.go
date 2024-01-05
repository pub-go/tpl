package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"code.gopub.tech/tpl/exp"
	"code.gopub.tech/tpl/exp/parser"
	"code.gopub.tech/tpl/html"
	"github.com/antlr4-go/antlr/v4"
	"github.com/youthlin/t"
	"github.com/youthlin/t/translator"
)

//go:embed lang
var embedLangs embed.FS

func main() {
	initTranslation()  // 加载翻译
	arg := parseArgs() // 解析参数
	run(arg)           // 开始运行
}

func initTranslation() {
	path, ok := os.LookupEnv("LANG_PATH")
	if ok {
		t.Load(path)
	} else {
		t.LoadFS(embedLangs)
	}
	t.SetLocale("")
}

func parseArgs() *Context {
	var (
		textTags = flag.String("text_tags", "",
			t.T("set tags that only contains text, default value is %s",
				strings.Join(html.GetDefaultTextTags(), ", ")))
		voidElements = flag.String("void_elements", "",
			t.T("set void elements which can not contains any child value, default value is %s",
				strings.Join(html.GetDefaultVoidElements(), ", ")))
		tagPrefix = flag.String("tag_prefix", "",
			t.T("set tag prefix, default value is %q", html.DefaultTagPrefix))
		attrPrefix = flag.String("attr_prefix", "",
			t.T("set attr prefix, default value is %q", html.DefaultAttrPrefix))
		path    = flag.String("path", ".", t.T("set template files' path"))
		pattern = flag.String("pattern", `.*\.html`, t.T("set template file name pattern"))
		// __("msgid")
		// _n("msgid", "plural")
		// _x("ctx", "msgid")
		// _xn("ctx", "msgid", "plural")
		keywords = flag.String("keywords", `T;N:1,2;N64:1,2;X:1c,2;XN:1c,2,3;XN64:1c,2,3;__;_n:1,2;_x:1c,2;_xn:1c,2,3`, t.T("set extract keywords"))
		output   = flag.String("output", "", t.T("set output path"))
		keys     []Keyword
	)
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), t.T("Usage of %s:", os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
	rexp, err := regexp.Compile(*pattern)
	if err != nil {
		die(t.T("invalid pattern %q: %v", *pattern, err))
	}
	kw, err := parseKeywords(*keywords) // 抽取关键字
	if err != nil {
		die(err.Error())
	}
	keys = kw
	var split = func(s *string) []string {
		a := strings.Split(*s, ",")
		var result = make([]string, 0, len(a))
		for _, i := range a {
			result = append(result, strings.TrimSpace(i))
		}
		return result
	}
	return &Context{
		textTags:     split(textTags),
		voidElements: split(voidElements),
		tagPrefix:    *tagPrefix,
		attrPrefix:   *attrPrefix,
		path:         *path,
		rawpatn:      *pattern,
		pattern:      rexp,
		keyword:      *keywords,
		Keywords:     keys,
		output:       *output,
	}
}

func die(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func run(ctx *Context) {
	m := html.NewTplManager().
		SetTextTags(ctx.textTags).
		SetVoidElements(ctx.voidElements).
		SetTagPrefix(ctx.tagPrefix).
		SetAttrPrefix(ctx.attrPrefix)
	err := m.ParseWithRegexp(os.DirFS(ctx.path), ctx.pattern)
	if err != nil {
		die(t.T("can not parse templates: path=%s, pattern=%s, err=%+v",
			ctx.path, ctx.pattern, err))
	}
	fmt.Println(t.T("path: %s, pattern: %s", ctx.path, ctx.pattern))
	for file, node := range m.Files() {
		fmt.Println(t.T("extract from file: %s", file))
		extract(ctx, file, node)
	}
	Save(ctx)
}

func extract(ctx *Context, file string, node *html.Node) {
	if token := node.Token; token != nil {
		if tag := token.Tag; tag != nil {
			for _, attr := range tag.Attrs {
				for _, vt := range attr.ValueTokens {
					if tree := vt.Tree; tree != nil {
						ctx.entries = append(ctx.entries,
							extractFromTree(ctx, file, vt.Start, tree)...)
					}
				}
			}
		}
	}
	for _, child := range node.Children {
		extract(ctx, file, child)
	}
}

func extractFromTree(ctx *Context, file string, pos exp.Pos, tree parser.IExpressionContext) []*translator.Entry {
	l := &listener{file: file, pos: pos, ctx: ctx}
	antlr.ParseTreeWalkerDefault.Walk(l, tree)
	return l.entries
}

type listener struct {
	parser.BaseGoExpressionListener
	file    string
	pos     exp.Pos
	ctx     *Context
	entries []*translator.Entry
}

func (l *listener) EnterPrimaryExpr(ctx *parser.PrimaryExprContext) {
	// primaryExpr -> operand
	// primaryExpr -> primaryExpr ( arguments | ... )
	// a.f()
	//           primaryExpr
	//            /       \
	//      primaryExpr arguments
	//         /     \       |
	//   primaryExpr field   ()
	//      /         |
	// operand        .f
	//   |
	//   a
	// __("msgid", args...)
	arg := ctx.Arguments()
	if arg == nil {
		return
	}
	left := ctx.PrimaryExpr()
	funName := getFnName(left)
	if funName == "" {
		return
	}
	debug("函数名: %s", funName)
	list := arg.ExpressionList()
	if list == nil {
		return
	}
	count := list.GetChildCount()
	if count == 0 {
		return
	}
	for _, kw := range l.ctx.Keywords {
		l.doExtract(kw, funName, list)
	}
}

func debug(f string, args ...any) {
	// fmt.Printf(f+"\n", args...)
}

func getFnName(ctx parser.IPrimaryExprContext) string {
	if ctx == nil {
		return ""
	}
	// primaryExpr -> operand
	op := ctx.Operand()
	if op != nil {
		if ex := op.Expression(); ex != nil {
			// operand -> '(' expression ')'
			debug("括号表达式: %s", ex.GetText())
			return getFnName(ex.PrimaryExpr())
		}
		// operand -> operandName
		if on := op.OperandName(); on != nil {
			debug("变量名%s", on.GetText())
			return on.GetText()
		}
		// operand -> literal // 字面量不能作为函数名
		return ""
	}
	// primaryExpr -> primaryExpr ( field | index | slice | arguments )
	pri := ctx.PrimaryExpr()
	if pri != nil {
		debug("后缀表达式: %s", ctx.GetText())
		if field := ctx.Field(); field != nil {
			debug("字段名: %s", field.GetText())
			return field.IDENTIFIER().GetText()
		}
		// index 索引不支持
		// t['T'] 可以直接使用 t.T
		// 索引里的表达式一般是字面量，用于从 map 中取含有空格的键
		// 像函数显然不会含有空格，可以直接用 field 表示
	}
	return ""
}

func isIdentifier(primary parser.IPrimaryExprContext) bool {
	op := primary.Operand()
	if op == nil {
		return false
	}
	on := op.OperandName()
	return on != nil
}

func (l *listener) doExtract(kw Keyword, fnName string, list parser.IExpressionListContext) {
	if kw.Name != fnName {
		return
	}
	maxArgs := kw.MaxArgIndex()
	count := list.GetChildCount()
	if maxArgs > count {
		return
	}
	entry := new(translator.Entry)
	if i := kw.MsgCtxt; i > 0 {
		param := list.Expression(i - 1)
		if s, ok := isStringLiteral(param); ok {
			entry.MsgCtxt = s
		}
	}
	if i := kw.MsgID; i > 0 {
		param := list.Expression(i - 1)
		if s, ok := isStringLiteral(param); ok {
			at := param.GetStart()
			entry.MsgID = s
			entry.MsgCmts = append(entry.MsgCmts, fmt.Sprintf("#: %s:%s",
				l.file, l.pos.Add(at.GetLine(), at.GetColumn())))
		}
	}
	if i := kw.MsgID2; i > 0 {
		param := list.Expression(i - 1)
		if s, ok := isStringLiteral(param); ok {
			entry.MsgID2 = s
		}
	}
	l.entries = append(l.entries, entry)
}

func isStringLiteral(tree parser.IExpressionContext) (string, bool) {
	if tree == nil {
		return "", false
	}
	primary := tree.PrimaryExpr()
	if primary == nil {
		return "", false
	}
	op := primary.Operand()
	if op == nil {
		return "", false
	}
	li := op.Literal()
	if li == nil {
		return "", false
	}
	s := li.String_()
	if s == nil {
		return "", false
	}
	return unquote(s.GetText()), true
}

func unquote(text string) string {
	if text[0] == '\'' {
		// Unquote 内部假定单引号括起的都是单个字符
		// 我们允许单引号括起字符串 这里兼容一下 转为双引号
		text = strings.ReplaceAll(text, "\\'", "'")
		text = `"` + text[1:len(text)-1] + `"`
	}
	t, _ := strconv.Unquote(text) // 去除引号
	return t
}

func buildHeader(ctx *Context) *translator.Entry {
	e := new(translator.Entry)
	e.MsgCmts = []string{
		"# SOME DESCRIPTIVE TITLE.",
		"# Copyright (C) YEAR THE PACKAGE'S COPYRIGHT HOLDER",
		"# This file is distributed under the same license as the PACKAGE package.",
		"# FIRST AUTHOR <EMAIL@ADDRESS>, YEAR.",
		"#",
		"#, fuzzy",
	}
	headers := []string{
		"Project-Id-Version: PACKAGE VERSION",
		"Report-Msgid-Bugs-To: ",
		fmt.Sprintf("POT-Creation-Date: %v",
			time.Now().Format("2006-01-02 15:04:05-0700")),
		"PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE",
		"Last-Translator: FULL NAME <EMAIL@ADDRESS>",
		"Language-Team: LANGUAGE <LL@li.org>",
		"Language: ",
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=UTF-8",
		"Content-Transfer-Encoding: 8bit",
		"Plural-Forms: nplurals=INTEGER; plural=EXPRESSION;",
		"X-Created-By: xtpl(https://github.com/pub-go/tpl/tree/main/cmd/xtpl)",
	}
	headers = append(headers, fmt.Sprintf("X-xTpl-Path: %v", ctx.path))
	headers = append(headers, fmt.Sprintf("X-xTpl-Pattern: %v", ctx.rawpatn))
	headers = append(headers, fmt.Sprintf("X-xTpl-Keywords: %v", ctx.keyword))
	headers = append(headers, fmt.Sprintf("X-xTpl-Output: %v", ctx.output))
	headers = append(headers, "")
	e.MsgStr = strings.Join(headers, "\n")
	return e
}

func Save(ctx *Context) {
	pot := new(translator.File)
	pot.AddEntry(buildHeader(ctx))
	entries := ctx.entries
	m := make(map[string]*translator.Entry, len(entries))
	for _, e := range entries {
		key := e.Key()
		if pre, ok := m[key]; ok {
			e.MsgCmts = append(pre.MsgCmts, e.MsgCmts...)
		}
		pot.AddEntry(e)
		m[key] = e
	}
	if name := ctx.output; name != "" {
		out, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			die(t.T("can not open file %q: %v", name, err))
		}
		pot.SaveAsPot(out)
		fmt.Println(t.T("saved to %s", name))
		return
	}
	pot.SaveAsPot(os.Stdout)
	fmt.Println(t.T("default output to stdout, if you want save to a file, please run with `-output <file>`"))
}
