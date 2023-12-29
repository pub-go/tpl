package types

import "io"

// Template 一个可以执行的模板
type Template interface {
	Execute(w io.Writer, data any) error
}

// TemplateManager 模板管理器
type TemplateManager interface {
	GetTemplate(name string) (Template, error)
}
