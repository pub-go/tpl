package main

import (
	"context"
	"embed"
	"net/http"
	"os"

	"code.gopub.tech/logs"
	"code.gopub.tech/tpl"
	"code.gopub.tech/tpl/exp"
	"code.gopub.tech/tpl/html"
	"code.gopub.tech/tpl/types"
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
)

//go:embed views
var views embed.FS

//go:embed lang
var lang embed.FS

func main() {
	logs.Info(context.Background(), "Hello, World")
	t.LoadFS(lang)
	go useGin()
	std()
}

func buildTpl(hotReload bool) (types.HTMLRender, error) {
	return tpl.NewHTMLRender(func(ctx context.Context) (types.TemplateManager, error) {
		logs.Info(ctx, "load templates...")
		m := html.NewTplManager()
		if hotReload {
			// 使用 os.DirFS 实时读取文件夹
			return m, m.ParseWithSuffix(os.DirFS("views"), ".html")
		}
		// 使用编译时嵌入的 embed.FS 资源
		return m, m.SetSubFS("views").ParseWithSuffix(views, ".html")
	}, tpl.WithHotReload(hotReload))
}

func std() {
	m, err := buildTpl(true)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := withI18n(r, types.M{
			"title":  "Welcome",
			"name":   "Tom",
			"number": 2,
		})
		err := m.Instance(ctx, "index.html", data).Render(w)
		// Or:
		// tpl, err := m.GetTemplate(ctx, "index.html")
		// if err == nil {
		// 	err = tpl.Execute(w, data)
		// }
		if err != nil {
			logs.Error(ctx, "err: %+v", err)
		}
	})
	http.ListenAndServe(":9999", nil)
}

func withI18n(r *http.Request, data types.M) types.M {
	lang := t.GetUserLang(r)
	logs.Info(r.Context(), "lang=%v %+v", lang, r.Header)
	t := t.L(lang)
	if data == nil {
		data = types.M{}
	}
	data["t"] = t
	data["__"] = t.T
	data["_x"] = t.X
	data["_n"] = func(msgID, msgIDPlural string, n any, args ...interface{}) string {
		return t.N64(msgID, msgIDPlural, exp.ToNumber[int64](n), args...)
	}
	data["_xn"] = func(msgCtx, msgID, msgIDPlural string, n any, args ...interface{}) string {
		return t.XN64(msgCtx, msgID, msgIDPlural, exp.ToNumber[int64](n), args...)
	}
	return data
}

func useGin() {
	r, err := buildTpl(gin.IsDebugging())
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		ctx := c.Request.Context()
		data := withI18n(c.Request, types.M{
			"title":  "Welcome",
			"name":   "Tom",
			"number": 2,
		})
		tpl := r.Instance(ctx, "index.html", data)
		c.Render(http.StatusOK, tpl)
	})
	engine.Run(":9998")
}
