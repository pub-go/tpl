package html

import (
	"bufio"
	"io"

	"code.gopub.tech/errors"
	"code.gopub.tech/logs"
	"code.gopub.tech/tpl/exp"
)

// BaseScanner 词法扫描器
type BaseScanner struct {
	reader  *bufio.Reader
	pos     Pos
	preCols int
	ch      rune
	chSize  int
	err     error
	done    bool
}

// NewBaseScanner 构造函数
func NewBaseScanner(r io.Reader) *BaseScanner {
	return (&BaseScanner{
		reader: bufio.NewReader(r),
	}).SetPos(exp.NewPos(1, 1))
}

// SetPos 设置位置
func (b *BaseScanner) SetPos(p Pos) *BaseScanner {
	b.pos = p
	return b
}

// NextRune 读取下一个字符
func (b *BaseScanner) NextRune() error {
	b.ch, b.chSize, b.err = b.reader.ReadRune()
	logs.Trace(bg, "read|ch=%c|size=%v|err=%v", b.ch, b.chSize, b.err)
	if b.err != nil {
		if errors.Is(b.err, io.EOF) {
			b.done = true
		}
		return b.err
	}
	if b.ch == '\n' {
		b.preCols = b.pos.Column
		b.pos.Line++
		b.pos.Column = 1
	} else if b.ch == '\t' {
		b.pos.Column += 4
	} else {
		b.pos.Column++
	}
	return nil
}

// UnRead 回退一个字符
func (b *BaseScanner) UnRead() {
	logs.Trace(bg, "unread|ch=%c", b.ch)
	b.reader.UnreadRune()
	if b.ch == '\n' {
		b.pos.Line--
		b.pos.Column = b.preCols
	} else if b.ch == '\t' {
		b.pos.Column -= 4
	} else {
		b.pos.Column--
	}
}

// Err 记录解析错误
func (b *BaseScanner) Err(format string, a ...any) error {
	b.err = errors.Errorf(format, a...)
	return b.err
}
