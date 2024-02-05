package exp

import (
	"fmt"

	"code.gopub.tech/errors"
	"code.gopub.tech/tpl/exp/parser"
	"github.com/antlr4-go/antlr/v4"
)

var ErrInputTooLong = fmt.Errorf("input too long")

// ParseCode 解析代码为语法树
func ParseCode(input string, opts ...parseCodeOpt) (tree parser.IExpressionContext, err error) {
	opt := &parseCodeOption{
		start: NewPos(1, 1),
	}
	for _, f := range opts {
		f(opt)
	}
	// 收集错误
	errListener := NewErrorListener(opt.start)

	// 词法分析
	lexer := parser.NewGoLexer(antlr.NewInputStream(input))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	// 语法分析
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGoExpression(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)

	tree = p.Expression()
	index := stream.Index()
	last := stream.Get(index)
	if last.GetTokenType() != antlr.TokenEOF {
		err = errors.Wrapf(ErrInputTooLong, "at index=%v token=%v", index, last)
		errListener.errors = append(errListener.errors, err)
	}
	return tree, errListener.Err()
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
	errors []error
}

// NewErrorListener 新建词法/语法分析错误监听器
func NewErrorListener(start Pos) *errorListener {
	return &errorListener{start: start}
}

// Err 获取词法/语法错误
func (e *errorListener) Err() error {
	return errors.Join(e.errors...)
}

// SyntaxError implements antlr.ErrorListener.
// 出现错误时回调
func (e *errorListener) SyntaxError(recognizer antlr.Recognizer,
	offendingSymbol interface{}, line int, column int,
	msg string, ex antlr.RecognitionException) {
	// 语法错误
	e.errors = append(e.errors, errors.Errorf("[SyntaxError] %v (position: %v)",
		msg, e.start.Add(line, column)))
}
