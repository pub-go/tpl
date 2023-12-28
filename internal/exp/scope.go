package exp

import (
	"errors"
)

var defaultScope = NewScope(map[string]any{
	"true":  true,
	"false": false,
})
var _ Scope = (*scope)(nil)

// Scope 表示一个作用域
type Scope interface {
	// Get 从该作用域内获取 name 对应的值
	// 返回 ErrNoSuchValue 表示该作用域中无 name 对应的值
	Get(name string) (result any, err error)
}

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
