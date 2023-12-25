package html

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"code.gopub.tech/logs"
	"code.gopub.tech/tpl/internal/exp"
	"code.gopub.tech/tpl/internal/exp/parser"
)

// CodeTokenKind 代码词法单元类型
type CodeTokenKind int

const (
	Error CodeTokenKind = iota
	// "Hello, ${name}"
	BegEnd    // 开始或结束的引号
	Literal   // 字面量
	CodeStart // ${
	CodeValue // 代码内容
	CodeEnd   // }

)

// CodeToken 代码词法单元
type CodeToken struct {
	Kind  CodeTokenKind
	Start Pos
	End   Pos
	Value string
	Tree  parser.IExpressionContext
}

func (t *CodeToken) String() string {
	return fmt.Sprintf("{Kind=%v,Value=%q,Start=%v,End=%v}",
		t.Kind, t.Value, t.Start, t.End)
}

// codeState 状态
type codeState int

const (
	codeInit  codeState = iota // 开始 下一个字符是 ' 或 "
	codeEnd                    // 结束 下一个字符是 ' 或 "
	codeText                   // 读取普通字符串 如 :text="Hello"
	codeBlock                  // 读取代码块 如 :text="${name}"
)

// CodeScanner 代码解析器
type CodeScanner struct {
	*BaseScanner
	state     codeState
	firstCh   rune // 记录开头结尾是 ' 还是 "
	brace     int  // ${} 内 { 的数量
	quote     rune // 字符串使用的引号
	nextToken *CodeToken
	tokens    []*CodeToken
}

// NewCodeScanner 新建代码解析器实例
func NewCodeScanner(start Pos, s string) *CodeScanner {
	return &CodeScanner{
		BaseScanner: &BaseScanner{
			reader: bufio.NewReader(strings.NewReader(s)),
			pos:    start,
		},
	}
}

// GetAllTokens 获取所有解析后的代码词法单元
func (t *CodeScanner) GetAllTokens() ([]*CodeToken, error) {
	for t.err == nil && !t.done {
		t.NextToken()
	}
	if errors.Is(t.err, io.EOF) {
		t.err = nil
	}
	return t.tokens, t.err
}

// NextToken 解析下一个词法单元
func (t *CodeScanner) NextToken() (*CodeToken, error) {
	if t.nextToken != nil {
		tok := t.nextToken
		t.nextToken = nil
		logs.Trace(bg, "return nextToken=%v", tok)
		return t.addToken(tok), nil
	}
	switch t.state {
	case codeInit, codeEnd: // 扫描开头结尾的引号
		return t.scanQuot()
	case codeText: // 扫描 ${} 外的字面文本
		return t.scanLiteral()
	case codeBlock: // 扫描 ${} 内的代码
		return t.scanCode()
	default:
		panic("unexpected state")
	}
}

// scanQuot 读取开头结尾的引号
func (t *CodeScanner) scanQuot() (*CodeToken, error) {
	start := t.pos
	logs.Trace(bg, "scanQuot|start=%v", start)
	if err := t.NextRune(); err != nil {
		logs.Trace(bg, "scanQuot|err=%v", err)
		// not wrap io.EOF
		return nil, t.Err("expected quote (' or \") at position %v: %w", start, err)
	}
	t.firstCh = t.ch
	switch t.ch {
	case '"', '\'':
		if t.state == codeEnd {
			t.done = true
		} else {
			t.state = codeText
		}
		return t.addToken(&CodeToken{
			Kind:  BegEnd,
			Value: string(t.ch),
			Start: start,
			End:   t.pos,
		}), nil
	}
	return nil, t.Err("expected quote (' or \") but got %c(from position %v to %v)",
		t.ch, start, t.pos)
}

// addToken 添加 token 到结果列表中
func (t *CodeScanner) addToken(tok *CodeToken) *CodeToken {
	logs.Trace(bg, "addToken|tok=%v", tok)
	t.tokens = append(t.tokens, tok)
	return tok
}

// scanLiteral 读取 ${} 外的字面字符串
func (t *CodeScanner) scanLiteral() (*CodeToken, error) {
	var (
		start = t.pos
		end   Pos
		buf   strings.Builder
	)
	logs.Trace(bg, "scanLiteral|start=%v", start)
	for {
		end = t.pos // ${ 之前的位置
		if err := t.NextRune(); err != nil {
			return nil, t.Err("scanLiteral failed from position %v to %v: %w",
				start, t.pos, err)
		}
		ch := t.ch
		str := buf.String()
		logs.Trace(bg, "scanLiteral|string=%q", str)
		if (ch == '"' || ch == '\'') && ch == t.firstCh {
			// 读取到了结尾的引号
			t.state = codeEnd
			if str == "" { // 无字面量字符串 直接返回结束符号
				return t.addToken(&CodeToken{
					Kind:  BegEnd,
					Value: string(ch),
					Start: start,
					End:   t.pos,
				}), nil
			}
			t.UnRead() // 下个循环再返回结束符号
			return t.addToken(&CodeToken{
				Kind:  Literal,
				Value: str,
				Start: start,
				End:   t.pos,
			}), nil
		}
		if ch == '$' {
			err := t.NextRune()
			if t.ch == '{' { // ${ 读取到了代码块开始标志
				t.state = codeBlock
				tok := &CodeToken{
					Kind:  CodeStart,
					Value: "${",
					Start: end,
					End:   t.pos,
				}
				if len(str) == 0 { // "${...}" 直接返回 ${
					return t.addToken(tok), nil
				}
				t.nextToken = tok // "xxx${...}" 先返回 xxx 下次返回 ${
				return t.addToken(&CodeToken{
					Kind:  Literal,
					Value: str,
					Start: start,
					End:   end,
				}), nil
			} else if err == nil {
				t.UnRead()
			}
		}
		buf.WriteRune(ch)
	}
}

// scanCode 读取 ${} 内的代码
func (t *CodeScanner) scanCode() (*CodeToken, error) {
	var (
		start = t.pos
		end   Pos
		buf   strings.Builder
	)
	logs.Trace(bg, "scanCode|start=%v", start)
	for {
		end = t.pos
		if err := t.NextRune(); err != nil {
			return nil, t.Err("scanCode failed, read value: %v(from position %v to %v): %w",
				buf.String(), start, t.pos, err)
		}
		var ch = t.ch
		buf.WriteRune(ch)
		switch ch {
		case '{': // ${} 内遇到 {
			t.brace++
		case '}': // 遇到 }
			if t.brace == 0 { // 代码块结束
				t.state = codeText
				t.nextToken = &CodeToken{
					Kind:  CodeEnd,
					Value: "}",
					Start: end,
					End:   t.pos,
				}
				value := buf.String()
				value = value[:len(value)-1] // 去掉 }
				tok := &CodeToken{
					Kind:  CodeValue,
					Value: value,
					Start: start,
					End:   end,
				}
				if err := compile(tok); err != nil {
					return nil, t.Err("invalid code: %w", err)
				}
				return t.addToken(tok), nil
			}
			t.brace--
		case '"', '\'', '`':
			t.quote = ch
			s, err := t.scanString()
			if err != nil {
				return nil, t.Err("scanString failed (from position %v to %v): %w",
					start, t.pos, err)
			}
			buf.WriteString(s)
		}
		continue
	}
}

// scanString 读取字符串
func (t *CodeScanner) scanString() (string, error) {
	var (
		start = t.pos
		raw   strings.Builder
	)
	for {
		if err := t.NextRune(); err != nil {
			return "", err
		}
		ch := t.ch
		raw.WriteRune(ch)
		logs.Trace(bg, "scanString:%c", t.ch)

		if t.quote == '`' { // 不转义
			if ch == '`' {
				// raw 字符串结束
				return raw.String(), nil
			}
		} else { // ' " 需要处理转义
			if ch == '\\' { // 转义
				if err := t.NextRune(); err != nil {
					return raw.String(), fmt.Errorf("read string failed (from position %v to %v): %w",
						start, t.pos, err)
				}
				raw.WriteRune(t.ch) // \"
			} else if ch == t.quote {
				// 结束
				return raw.String(), nil
			}
		}
	}
}

func compile(tok *CodeToken) error {
	tree, err := exp.ParseCode(tok.Value, exp.WithStartPos(tok.Start))
	tok.Tree = tree
	return err
}

var bg = context.Background()
