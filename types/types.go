package types

import (
	"context"
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
	// 因模板可以 reload, 当初始化时, ctx 是 background; 当 reload 时, ctx 是传入的当次请求的上下文.
	Factory func(ctx context.Context) (TemplateManager, error)

	// Render 借用 gin 的接口 方便集成到 gin 中
	Render interface {
		// Render writes data with custom ContentType.
		Render(http.ResponseWriter) error
		// WriteContentType writes custom ContentType.
		WriteContentType(w http.ResponseWriter)
	}
	// HTMLRender
	HTMLRender interface {
		Instance(context.Context, string, any) Render
		GetTemplate(context.Context, string) (Template, error)
	}
	// ReloadableRender 支持重新解析模板的 HTMLRender
	ReloadableRender interface {
		HTMLRender
		Reload(ctx context.Context) error
	}
)
