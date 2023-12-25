package tpl

import (
	"io"
	"strings"
)

type Template interface {
	Execute(w io.Writer, data any) error
}

func Render(tpl Template, data any) (string, error) {
	var sb strings.Builder
	err := tpl.Execute(&sb, data)
	return sb.String(), err
}
