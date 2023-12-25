package html

import (
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
	err = h.Execute(os.Stdout, nil)
	if err != nil {
		t.Errorf("execute template error: %v", err)
	}
}
