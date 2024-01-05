package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"

	"code.gopub.tech/logs"
	"code.gopub.tech/tpl"
	"code.gopub.tech/tpl/exp"
	"code.gopub.tech/tpl/html"
	"code.gopub.tech/tpl/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
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
	return tpl.NewHTMLRender(func() (types.TemplateManager, error) {
		m := html.NewTplManager()
		if hotReload {
			// 使用 os.DirFS 实时读取文件夹
			return m, m.ParseWithSuffix(os.DirFS("views"), ".html")
		}
		// 使用编译时嵌入的 embed.FS 资源
		return m, m.SetSubFS("views").ParseWithSuffix(views, ".html")
	}, hotReload)
}

func std() {
	m, err := buildTpl(true)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := m.Instance("index.html", withI18n(r, types.M{
			"title":  "Welcome",
			"name":   "Tom",
			"number": 2,
		})).Render(w)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
	})
	http.ListenAndServe(":9999", nil)
}

func withI18n(r *http.Request, data types.M) types.M {
	lang := t.GetUserLang(r)
	fmt.Printf("lang=%v %+v\n", lang, r.Header)
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

type htmlRender struct {
	types.HTMLRender
}

func (h htmlRender) Instance(s string, a any) render.Render {
	return h.HTMLRender.Instance(s, a)
}

func useGin() {
	r, err := buildTpl(gin.IsDebugging())
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	engine.HTMLRender = htmlRender{HTMLRender: r}
	engine.GET("/", func(c *gin.Context) {
		data := withI18n(c.Request, types.M{
			"title":  "Welcome",
			"name":   "Tom",
			"number": 2,
		})
		c.HTML(http.StatusOK, "index.html", data)
		// c.Render(http.StatusOK, r.Instance("index.html", data))
	})
	engine.Run(":9998")
}
