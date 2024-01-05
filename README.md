# tpl
template like thymeleaf

## 模板 Templates
```html
<p :text="${'Hello, ' + name}">Hello, World!</p>
```

### Examples
see sub dir: [example](./example/)

## 代码 Codes

### 基本用法 Basic Usage
```go
import (
	"code.gopub.tech/tpl"
	"code.gopub.tech/tpl/html"
	"code.gopub.tech/tpl/types"
)
// 解析模板文件夹
m := html.NewTplManager()
err := m.ParseWithSuffix(os.DirFS("path/to/tpl"),".html")
// 获取并执行模板
tpl, err := m.GetTemplate("index.html")
err = tpl.Execute(writer, map[string]any{})
```

### 另一种方式 HTMLRender
```go
//go:embed views
var views embed.FS
var hotReload = false // gin.IsDebugging()
// 使用 HTMLRender 接口，支持热加载
r, err := tpl.NewHTMLRender(func() (types.TemplateManager, error) {
    m := html.NewTplManager()
    if hotReload {
        // 使用 os.DirFS 实时读取文件夹
        return m, m.ParseWithSuffix(os.DirFS("views"), ".html")
    }
    // 使用编译时嵌入的 embed.FS 资源
    return m, m.SetSubFS("views").ParseWithSuffix(views, ".html")
}, hotReload)
// 获取并执行模板
render := r.Instance("index.html", types.M{})
err := render.Render(writer)
```

## 如需支持 gettext 国际化(i18n)

### Codes
```go
import (
	"github.com/youthlin/t"
	"github.com/youthlin/t/exp"
)

err = tpl.Execute(writer, withI18n(req, map[string]any{
	"name": name,
}))
render := r.Instance("index.html", withI18n(req, types.M{
	"name": name,
}))

func withI18n(r *http.Request, data types.M) types.M {
	lang := t.GetUserLang(r)
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


```

### Templates
```html
<p :text="${__('Hello, %s', name)}">Hello, World!</p>

```

### Extract
see [xtpl](./cmd/xtpl/)
```bash
go install code.gopub.tech/tpl/cmd/xtpl@latest
xtpl -help
```
