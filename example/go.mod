module code.gopub.tech/tpl/example

go 1.21.0

replace code.gopub.tech/tpl => ../

// 因为第三方项目依赖了 github.com/gin-gonic/gin v1.4.0
//     (可以使用 `go mod graph` 查看)
// 但是这个版本的 gin 依赖了 github.com/ugorji/go v1.1.4
//     (github.com/ugorji/go 路径下有 go.mod github.com/ugorji/go/codec 是他的子包)
// 而新版 gin 依赖的是 	github.com/ugorji/go/codec v1.2.11
//     (github.com/ugorji/go/codec 路径下有 go.mod 他是单独的包)
// 导致编译时 github.com/ugorji/go/codec 不知道该用哪个
//     ambiguous import: found package github.com/ugorji/go/codec in multiple modules:
//         github.com/ugorji/go v1.1.4
//         github.com/ugorji/go/codec v1.2.12
// 所以这里直接强制替换 gin 老版本指向新版本
replace github.com/gin-gonic/gin v1.4.0 => github.com/gin-gonic/gin v1.9.1

require (
	code.gopub.tech/logs v0.0.5
	code.gopub.tech/tpl v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.9.1
	github.com/youthlin/t v0.0.7
)

require (
	github.com/Xuanwo/go-locale v1.0.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20210803070921-b358b509191a // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cockroachdb/errors v1.8.6 // indirect
	github.com/cockroachdb/logtags v0.0.0-20190617123548-eb05cc24525f // indirect
	github.com/cockroachdb/redact v1.1.1 // indirect
	github.com/cockroachdb/sentry-go v0.6.1-cockroachdb.2 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
