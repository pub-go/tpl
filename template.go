package tpl

import (
	"strings"

	"code.gopub.tech/tpl/internal/types"
)

type Template = types.Template
type TemplateManager = types.TemplateManager

// Render 执行一个模板 并将结果输出为字符串
func Render(tpl Template, data any) (string, error) {
	var sb strings.Builder
	err := tpl.Execute(&sb, data)
	return sb.String(), err
}
