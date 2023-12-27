package html

import (
	"fmt"
	"os"
	"testing"
)

func TestNewHtmlTemplate(t *testing.T) {
	f, err := os.Open("../testdata/index.html")
	if err != nil {
		t.Errorf("failed to open test file: %v", err)
	}
	h := NewHtmlTemplate()
	err = h.Add("index.html", f)
	if err != nil {
		t.Errorf("add template error: %v", err)
	}
	type Item struct {
		ID int
	}
	err = h.Execute(os.Stdout, map[string]any{
		"t":     func(input string, args ...any) string { return fmt.Sprintf(input, args...) },
		"name":  "<b>Tom</b>",
		"items": []Item{{ID: 999}, {ID: 666}},
		"len":   func(a []Item) int { return len(a) },
	})
	if err != nil {
		t.Errorf("execute template error: %+v", err)
	}
}
