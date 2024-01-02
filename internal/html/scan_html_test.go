package html_test

import (
	"fmt"
	"strings"
	"testing"

	"code.gopub.tech/logs/pkg/arg"
	"code.gopub.tech/tpl/internal/exp"
	"code.gopub.tech/tpl/internal/html"
)

func ExampleHtmlScanner() {
	s := html.NewHtmlScanner(strings.NewReader(`<!DOCTYPE html>
<html lang = "zh">
<head>
	<meta charset="UTF-8">
    <title>标题</title>
</head>
<body>
	<script>/*<script*/</script>
</body>
</html>`))
	tokens, err := s.GetAllTokens()
	if err != nil {
		panic(err)
	}
	fmt.Printf("tokens: %v", arg.JSON(tokens))
	// Output:
	// tokens: [{"Kind":1,"Value":"\u003c!DOCTYPE html\u003e","Start":{"Line":1,"Column":1},"End":{"Line":1,"Column":16},"Tag":{"Name":"!DOCTYPE","Attrs":[{"Name":"html","NameStart":{"Line":1,"Column":11},"NameEnd":{"Line":1,"Column":15},"Value":null,"ValueStart":{"Line":0,"Column":0},"ValueEnd":{"Line":0,"Column":0},"ValueTokens":null}]}},{"Kind":2,"Value":"\n","Start":{"Line":1,"Column":16},"End":{"Line":2,"Column":1},"Tag":null},{"Kind":1,"Value":"\u003chtml lang = \"zh\"\u003e","Start":{"Line":2,"Column":1},"End":{"Line":2,"Column":19},"Tag":{"Name":"html","Attrs":[{"Name":"lang","NameStart":{"Line":2,"Column":7},"NameEnd":{"Line":2,"Column":11},"Value":"\"zh\"","ValueStart":{"Line":2,"Column":14},"ValueEnd":{"Line":2,"Column":18},"ValueTokens":null}]}},{"Kind":2,"Value":"\n","Start":{"Line":2,"Column":19},"End":{"Line":3,"Column":1},"Tag":null},{"Kind":1,"Value":"\u003chead\u003e","Start":{"Line":3,"Column":1},"End":{"Line":3,"Column":7},"Tag":{"Name":"head","Attrs":null}},{"Kind":2,"Value":"\n\t","Start":{"Line":3,"Column":7},"End":{"Line":4,"Column":5},"Tag":null},{"Kind":1,"Value":"\u003cmeta charset=\"UTF-8\"\u003e","Start":{"Line":4,"Column":5},"End":{"Line":4,"Column":27},"Tag":{"Name":"meta","Attrs":[{"Name":"charset","NameStart":{"Line":4,"Column":11},"NameEnd":{"Line":4,"Column":18},"Value":"\"UTF-8\"","ValueStart":{"Line":4,"Column":19},"ValueEnd":{"Line":4,"Column":26},"ValueTokens":null}]}},{"Kind":2,"Value":"\n    ","Start":{"Line":4,"Column":27},"End":{"Line":5,"Column":5},"Tag":null},{"Kind":1,"Value":"\u003ctitle\u003e","Start":{"Line":5,"Column":5},"End":{"Line":5,"Column":12},"Tag":{"Name":"title","Attrs":null}},{"Kind":2,"Value":"标题","Start":{"Line":5,"Column":12},"End":{"Line":5,"Column":14},"Tag":null},{"Kind":1,"Value":"\u003c/title\u003e","Start":{"Line":5,"Column":14},"End":{"Line":5,"Column":22},"Tag":{"Name":"/title","Attrs":null}},{"Kind":2,"Value":"\n","Start":{"Line":5,"Column":22},"End":{"Line":6,"Column":1},"Tag":null},{"Kind":1,"Value":"\u003c/head\u003e","Start":{"Line":6,"Column":1},"End":{"Line":6,"Column":8},"Tag":{"Name":"/head","Attrs":null}},{"Kind":2,"Value":"\n","Start":{"Line":6,"Column":8},"End":{"Line":7,"Column":1},"Tag":null},{"Kind":1,"Value":"\u003cbody\u003e","Start":{"Line":7,"Column":1},"End":{"Line":7,"Column":7},"Tag":{"Name":"body","Attrs":null}},{"Kind":2,"Value":"\n\t","Start":{"Line":7,"Column":7},"End":{"Line":8,"Column":5},"Tag":null},{"Kind":1,"Value":"\u003cscript\u003e","Start":{"Line":8,"Column":5},"End":{"Line":8,"Column":13},"Tag":{"Name":"script","Attrs":null}},{"Kind":2,"Value":"/*\u003cscript*/","Start":{"Line":8,"Column":13},"End":{"Line":8,"Column":24},"Tag":null},{"Kind":1,"Value":"\u003c/script\u003e","Start":{"Line":8,"Column":24},"End":{"Line":8,"Column":33},"Tag":{"Name":"/script","Attrs":null}},{"Kind":2,"Value":"\n","Start":{"Line":8,"Column":33},"End":{"Line":9,"Column":1},"Tag":null},{"Kind":1,"Value":"\u003c/body\u003e","Start":{"Line":9,"Column":1},"End":{"Line":9,"Column":8},"Tag":{"Name":"/body","Attrs":null}},{"Kind":2,"Value":"\n","Start":{"Line":9,"Column":8},"End":{"Line":10,"Column":1},"Tag":null},{"Kind":1,"Value":"\u003c/html\u003e","Start":{"Line":10,"Column":1},"End":{"Line":10,"Column":8},"Tag":{"Name":"/html","Attrs":null}}]
}

func ptr[T any](a T) *T { return &a }

func TestHtmlScanner_GetAllTokens1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []*html.Token
		wantErr bool
		err     string
	}{
		{
			name: "html",
			input: `<!DOCTYPE html>
<html lang = "zh">
<!-- -->Text<![CDATA[data]]>
<p>Hello, World!</p>
<script>/*<script*/</script>
</html>`,
			want: []*html.Token{
				{
					Kind:  html.TokenKindTag,
					Value: "<!DOCTYPE html>",
					Start: exp.NewPos(1, 1),
					End:   exp.NewPos(1, 16),
					Tag: &html.Tag{
						Name: "!DOCTYPE",
						Attrs: []*html.Attr{
							{
								Name:      "html",
								NameStart: exp.NewPos(1, 11),
								NameEnd:   exp.NewPos(1, 15),
							},
						},
					},
				},
				{
					Kind:  html.TokenKindText,
					Value: "\n",
					Start: exp.NewPos(1, 16),
					End:   exp.NewPos(2, 1),
				},
				{
					Kind:  html.TokenKindTag,
					Value: `<html lang = "zh">`,
					Start: exp.NewPos(2, 1),
					End:   exp.NewPos(2, 19),
					Tag: &html.Tag{
						Name: "html",
						Attrs: []*html.Attr{
							{
								Name:       "lang",
								NameStart:  exp.NewPos(2, 7),
								NameEnd:    exp.NewPos(2, 11),
								Value:      ptr(`"zh"`),
								ValueStart: exp.NewPos(2, 14),
								ValueEnd:   exp.NewPos(2, 18),
							},
						},
					},
				},
				{
					Kind:  html.TokenKindText,
					Value: "\n",
					Start: exp.NewPos(2, 19),
					End:   exp.NewPos(3, 1),
				},
				{
					Kind:  html.TokenKindComment,
					Value: "<!-- -->",
					Start: exp.NewPos(3, 1),
					End:   exp.NewPos(3, 9),
				},
				{
					Kind:  html.TokenKindText,
					Value: "Text",
					Start: exp.NewPos(3, 9),
					End:   exp.NewPos(3, 13),
				},
				{
					Kind:  html.TokenKindCDATA,
					Value: "<![CDATA[data]]>",
					Start: exp.NewPos(3, 13),
					End:   exp.NewPos(3, 29),
				},
				{
					Kind:  html.TokenKindText,
					Value: "\n",
					Start: exp.NewPos(3, 29),
					End:   exp.NewPos(4, 1),
				},
				{
					Kind:  html.TokenKindTag,
					Value: "<p>",
					Start: exp.NewPos(4, 1),
					End:   exp.NewPos(4, 4),
					Tag:   &html.Tag{Name: "p"},
				},
				{
					Kind:  html.TokenKindText,
					Value: "Hello, World!",
					Start: exp.NewPos(4, 4),
					End:   exp.NewPos(4, 17),
				},
				{
					Kind:  html.TokenKindTag,
					Value: "</p>",
					Start: exp.NewPos(4, 17),
					End:   exp.NewPos(4, 21),
					Tag:   &html.Tag{Name: "/p"},
				},
				{
					Kind:  html.TokenKindText,
					Value: "\n",
					Start: exp.NewPos(4, 21),
					End:   exp.NewPos(5, 1),
				},
				{
					Kind:  html.TokenKindTag,
					Value: "<script>",
					Start: exp.NewPos(5, 1),
					End:   exp.NewPos(5, 9),
					Tag: &html.Tag{
						Name: "script",
					},
				},
				{
					Kind:  html.TokenKindText,
					Value: "/*<script*/",
					Start: exp.NewPos(5, 9),
					End:   exp.NewPos(5, 20),
				},
				{
					Kind:  html.TokenKindTag,
					Value: "</script>",
					Start: exp.NewPos(5, 20),
					End:   exp.NewPos(5, 29),
					Tag: &html.Tag{
						Name: "/script",
					},
				},
				{
					Kind:  html.TokenKindText,
					Value: "\n",
					Start: exp.NewPos(5, 29),
					End:   exp.NewPos(6, 1),
				},
				{
					Kind:  html.TokenKindTag,
					Value: "</html>",
					Start: exp.NewPos(6, 1),
					End:   exp.NewPos(6, 8),
					Tag: &html.Tag{
						Name: "/html",
					},
				},
			},
		},
		{
			name:  "text-eof",
			input: "Hello",
			want: []*html.Token{
				{
					Kind:  html.TokenKindText,
					Value: "Hello",
					Start: exp.NewPos(1, 1),
					End:   exp.NewPos(1, 6),
				},
			},
		},
		{
			name:    "tag-eof",
			input:   "<p",
			wantErr: true,
			err:     `read tag "<p" failed (from position 1:1 to 1:3): unecpected EOF`,
		},
		{
			name:    "comment-error",
			input:   "<!-->-->",
			wantErr: true,
			err:     `read comment failed (from position 1:1 to 1:6): comment text must not start with ">" or "->": >`,
		},
		{
			name:    "comment-error2",
			input:   "<!-- <!-- -->",
			wantErr: true,
			err:     `read comment failed (from position 1:1 to 1:14): comment text must not contains "<!--", "-->" or "--!>":  <!--`,
		},
		{
			name:    "comment-error2",
			input:   "<!-- <!--->",
			wantErr: true,
			err:     `read comment failed (from position 1:1 to 1:12): comment text must not end with "<!-":  <!-`,
		},
		{
			name:  "attr-name-only",
			input: `<input checked disabled k=v>`,
			want: []*html.Token{
				{
					Kind:  html.TokenKindTag,
					Value: `<input checked disabled k=v>`,
					Start: exp.NewPos(1, 1),
					End:   exp.NewPos(1, 29),
					Tag: &html.Tag{
						Name: "input",
						Attrs: []*html.Attr{
							{Name: "checked", NameStart: exp.NewPos(1, 8), NameEnd: exp.NewPos(1, 15)},
							{Name: "disabled", NameStart: exp.NewPos(1, 16), NameEnd: exp.NewPos(1, 24)},
							{Name: "k", NameStart: exp.NewPos(1, 25), NameEnd: exp.NewPos(1, 26),
								Value: ptr("v"), ValueStart: exp.NewPos(1, 27), ValueEnd: exp.NewPos(1, 28)},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := html.NewHtmlScanner(strings.NewReader(tt.input))
			got, err := tr.GetAllTokens()
			if (err != nil) != tt.wantErr {
				t.Errorf("HtmlScanner.GetAllTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Logf("err=%v", err)
				if !strings.Contains(err.Error(), tt.err) {
					t.Errorf("HtmlScanner.GetAllTokens() error = %v, want contains %s", err, tt.err)
				}
			}
			if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", tt.want) {
				t.Errorf("HtmlScanner.GetAllTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
