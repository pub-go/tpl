package tpl

import (
	"bytes"
	"context"
	"net/http"
	"strings"

	"code.gopub.tech/tpl/types"
)

// RenderToString 执行一个模板 并将结果输出为字符串
func RenderToString(tpl types.Template, data any) (string, error) {
	var sb strings.Builder
	err := tpl.Execute(&sb, data)
	return sb.String(), err
}

// RenderToString 执行一个模板 并将结果输出为字节数组
func RenderToBytes(tpl types.Template, data any) ([]byte, error) {
	var sb bytes.Buffer
	err := tpl.Execute(&sb, data)
	return sb.Bytes(), err
}

// NewHTMLRender 新建一个 HTML 渲染器
//
//	//go:embed views
//	var views embed.FS
//
//	var hotReload = gin.IsDebugging()
//	// NewHTMLRender
//	r, err := tpl.NewHTMLRender(func(context.Context) (types.TemplateManager, error) {
//		m := html.NewTplManager()
//		if hotReload {
//			// 使用 os.DirFS 实时读取文件夹
//			return m, m.ParseWithSuffix(os.DirFS("views"), ".html")
//		}
//		// 使用编译时嵌入的 embed.FS 资源
//		return m, m.SetSubFS("views").ParseWithSuffix(views, ".html")
//	}, tpl.WithHotReload(hotReload))
//
//	ginS.Get("/", func(c *gin.Context) {
//		// Instance
//		c.Render(http.StatusOK, r.Instance(c, "index.html", gin.H{}))
//	})
//	ginS.Run()
func NewHTMLRender(builder types.Factory, opts ...NewHTMLRenderOpt) (types.ReloadableRender, error) {
	opt := &newHtmlRenderOpt{}
	for _, setter := range opts {
		setter(opt)
	}
	r := &htmlRender{
		hotReload: opt.hotReload,
		builder:   builder,
	}
	return r, r.Reload(context.Background()) // 首次 build 时无 ctx
}

type newHtmlRenderOpt struct {
	hotReload bool
}
type NewHTMLRenderOpt func(*newHtmlRenderOpt)

// WithHotReload 当需要热加载时 每次渲染都会重新解析模板.
// 出于性能考虑 应当仅在调试阶段开启;
// 正式环境如需重新解析模板 可以通过下方 Reload 方法触发.
func WithHotReload(hotReload bool) NewHTMLRenderOpt {
	return func(o *newHtmlRenderOpt) { o.hotReload = hotReload }
}

type htmlRender struct {
	hotReload bool
	builder   types.Factory
	manager   types.TemplateManager
}

// Reload implements types.ReloadableRender.
// 重新解析模板文件.
func (h *htmlRender) Reload(ctx context.Context) error {
	m, err := h.builder(ctx)
	if err != nil {
		return err
	}
	h.manager = m
	return nil
}

// Instance implements types.HTMLRender.
// 获取 Render 实例.
func (h *htmlRender) Instance(ctx context.Context, tplName string, data any) types.Render {
	tpl, err := h.GetTemplate(ctx, tplName)
	if err != nil {
		return &render{
			err: err,
		}
	}
	return &render{
		tpl: tpl,
		obj: data,
	}
}

// GetTemplate implements types.HTMLRender.
// 获取一个模板实例.
func (h *htmlRender) GetTemplate(ctx context.Context, tplName string) (types.Template, error) {
	var (
		m   = h.manager
		err error
	)
	if h.hotReload {
		m, err = h.builder(ctx)
	}
	if err != nil {
		return nil, err
	}
	return m.GetTemplate(tplName)
}

type render struct {
	err error
	tpl types.Template
	obj any
}

// Render implements types.Render.
func (r *render) Render(w http.ResponseWriter) error {
	if r.err != nil {
		return r.err
	}
	return r.tpl.Execute(w, r.obj)
}

var htmlContentType = []string{"text/html; charset=utf-8"}

// WriteContentType implements types.Render.
func (*render) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = htmlContentType
	}
}
