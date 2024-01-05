package html

import (
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestParser_ParseTokens(t *testing.T) {
	f, err := os.Open("../testdata/index.tpl.html")
	if err != nil {
		t.Errorf("failed to open test file: %v", err)
	}
	tz := NewHtmlScanner(f)
	tokens, err := tz.GetAllTokens()
	if err != nil {
		t.Errorf("tokenize error: %v", err)
	}
	t.Logf("tokens=%v", tokens)
	p := NewParser()
	node, err := p.ParseTokens(tokens)
	t.Logf("node=%v, %s, err=%v", node, n2s(node), err)
	slog.Info("test end")
}

func n2s(node *Node) string {
	var sb strings.Builder
	var node2s func(sb *strings.Builder, n *Node)
	node2s = func(sb *strings.Builder, n *Node) {
		if n.Token != nil {
			sb.WriteString(n.Token.Value)
		}
		for _, c := range n.Children {
			node2s(sb, c)
		}
		if n.End != nil {
			sb.WriteString(n.End.Value)
		}
	}
	node2s(&sb, node)
	return sb.String()
}
