package exp

import "fmt"

// Pos 行列位置
type Pos struct {
	Line   int // 行
	Column int // 列
}

func NewPos(line, column int) Pos {
	return Pos{Line: line, Column: column}
}

func (p Pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line, p.Column)
}

func (p Pos) Add(line, column int) Pos {
	return Pos{p.Line + line - 1, p.Column + column}
}
