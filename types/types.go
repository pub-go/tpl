package types

import (
	"io"
	"net/http"
)

type (
	M = map[string]any
	// Template 一个可以执行的模板
	Template interface {
		Execute(w io.Writer, data any) error
	}

	// TemplateManager 模板管理器
	TemplateManager interface {
		GetTemplate(name string) (Template, error)
	}

	// Factory 模板管理器工厂
	Factory func() (TemplateManager, error)

	// HTMLRender interface is to be implemented by HTMLProduction and HTMLDebug.
	// see gin doc
	HTMLRender interface {
		// Instance returns an HTML instance.
		Instance(string, any) Render
		GetTemplate(string) (Template, error)
	}
	ReloadableRender interface {
		HTMLRender
		Reload() error
	}

	// Render interface is to be implemented by JSON, XML, HTML, YAML and so on.
	// see gin doc
	Render interface {
		// Render writes data with custom ContentType.
		Render(http.ResponseWriter) error
		// WriteContentType writes custom ContentType.
		WriteContentType(w http.ResponseWriter)
	}
)
