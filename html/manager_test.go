package html

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestNewHtmlTemplate(t *testing.T) {
	m := NewTplManager()
	m.maxIndent=5
	err := m.Parse(os.DirFS("../testdata/"), func(path string) bool { return strings.HasSuffix(path, ".tpl.html") })
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	type Item struct {
		ID int
	}
	tpl, err := m.GetTemplate("index.tpl.html")
	if err != nil {
		t.Fatalf("execute template error: %+v", err)
	}
	var i int
	var buf strings.Builder
	err = tpl.Execute(&buf, map[string]any{
		"t":          func(input string, args ...any) string { return fmt.Sprintf(input, args...) },
		"get":        func() int { return i },
		"incrAndGet": func() int { i++; return i },
		"name":       "<b>Tom</b>",
		"hideItems":  false,
		"items":      []Item{{ID: 999}, {ID: 666}},
		"len":        func(a []Item) int { return len(a) },
	})
	if err != nil {
		t.Fatalf("execute template error: %+v", err)
	}
	result, _ := os.Open("../testdata/index.result.html")
	os.WriteFile("../testdata/index.got.html", []byte(buf.String()), 0777)
	want, _ := io.ReadAll(result)
	got := buf.String()
	if string(want) != got {
		t.Fatalf("want index.result.html got: %s", got)
	}
}
