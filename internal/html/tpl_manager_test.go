package html

import (
	"fmt"
	"os"
	"testing"
)

func TestNewHtmlTemplate(t *testing.T) {
	m := NewTplManager()
	err := m.Glob(os.DirFS("../testdata/"), "*")
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	type Item struct {
		ID int
	}
	tpl, err := m.GetTemplate("index.html")
	if err != nil {
		t.Fatalf("execute template error: %+v", err)
	}
	err = tpl.Execute(os.Stdout, map[string]any{
		"t":         func(input string, args ...any) string { return fmt.Sprintf(input, args...) },
		"name":      "<b>Tom</b>",
		"hideItems": false,
		"items":     []Item{{ID: 999}, {ID: 666}},
		"len":       func(a []Item) int { return len(a) },
	})
	if err != nil {
		t.Fatalf("execute template error: %+v", err)
	}
}
