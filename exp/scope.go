package exp

import (
	"errors"
	"fmt"
)

var defaultScope = NewScope(map[string]any{
	"true":    true,
	"false":   false,
	"string":  ToString,
	"bytes":   ToBytes,
	"runes":   ToRunes,
	"int":     ToNumber[int],
	"int8":    ToNumber[int8],
	"int16":   ToNumber[int16],
	"int32":   ToNumber[int32],
	"int64":   ToNumber[int64],
	"uint":    ToNumber[uint],
	"uint8":   ToNumber[uint8],
	"uint16":  ToNumber[uint16],
	"uint32":  ToNumber[uint32],
	"uint64":  ToNumber[uint64],
	"float32": ToNumber[float32],
	"float64": ToNumber[float64],
	"isNil":   func(a any) bool { return a == nil },
	"printf":  func(format string, args ...any) string { return fmt.Sprintf(format, args...) },
	"println": func(a ...any) string { return fmt.Sprintln(a...) },
})
var _ Scope = (*scope)(nil)

// Scope 表示一个作用域
type Scope interface {
	// Get 从该作用域内获取 name 对应的值
	// 返回 ErrNoSuchValue 表示该作用域中无 name 对应的值
	Get(name string) (result any, err error)
}

// EmptyScope 构造一个没有任何值的作用域
func EmptyScope() Scope { return NewScope(map[string]any{}) }

// NewScope 从结构体/map/数组/切片新建作用域
func NewScope(data any) Scope {
	return &scope{data: data}
}

// scope 作用域接口的默认实现
type scope struct {
	data any
}

// Get implements Scope.
func (s *scope) Get(name string) (any, error) {
	return getValue(name, s.data)
}

// WithDefaultScope 组合默认作用域
func WithDefaultScope(s Scope) Scope {
	return Combine(s, defaultScope)
}

// Combine 组合两个作用域 当子作用域中找不到时 从父作用域找
func Combine(s, p Scope) Scope {
	return &combineScope{
		parent: p,
		scope:  s,
	}
}

// combineScope 组合作用域接口的默认实现
type combineScope struct {
	parent Scope
	scope  Scope
}

// Get implements Scope.
func (s *combineScope) Get(name string) (any, error) {
	v, e := s.scope.Get(name)
	if errors.Is(e, ErrNoSuchValue) {
		return s.parent.Get(name)
	}
	return v, e
}
