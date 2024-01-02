package exp

import (
	"fmt"
	"strings"

	"code.gopub.tech/tpl/internal/exp/parser"
	"github.com/antlr4-go/antlr/v4"
)

// ParseCode 解析代码为语法树
func ParseCode(input string, opts ...parseCodeOpt) (tree parser.IExpressionContext, err error) {
	opt := &parseCodeOption{
		start: NewPos(1, 1),
	}
	for _, f := range opts {
		f(opt)
	}
	errListener := NewErrorListener(opt.start)

	lexer := parser.NewGoLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	p := parser.NewGoExpression(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	return p.Expression(), errListener.Err()
}

type parseCodeOpt func(*parseCodeOption)

type parseCodeOption struct {
	start Pos
}

// WithStartPos 设置代码的起始位置 默认是起始于第 1 行第 1 列
func WithStartPos(pos Pos) parseCodeOpt {
	return func(pco *parseCodeOption) { pco.start = pos }
}

type errorListener struct {
	*antlr.DefaultErrorListener
	start  Pos
	errors []string
}

// NewErrorListener 新建词法/语法分析错误监听器
func NewErrorListener(start Pos) *errorListener {
	return &errorListener{start: start}
}

// Err 获取词法/语法错误
func (e *errorListener) Err() error {
	if len(e.errors) > 0 {
		return fmt.Errorf("compile error: %s", strings.Join(e.errors, "; "))
	}
	return nil
}

// SyntaxError implements antlr.ErrorListener.
// 出现错误时回调
func (e *errorListener) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol interface{}, line int, column int,
	msg string, ex antlr.RecognitionException) {
	// 语法错误
	e.errors = append(e.errors, fmt.Sprintf("[SyntaxError] %v (position: %v)",
		msg, e.start.Add(line, column)))
}
