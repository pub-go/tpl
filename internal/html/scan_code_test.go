package html_test

import (
	"reflect"
	"strings"
	"testing"

	"code.gopub.tech/logs"
	"code.gopub.tech/tpl/internal/exp"
	"code.gopub.tech/tpl/internal/html"
)

func TestCodeScanner_GetAllTokens(t *testing.T) {
	logs.SetDefault(logs.NewLogger(logs.NewHandler(
	//logs.WithLevel(logs.LevelTrace),
	)))
	tests := []struct {
		name    string
		input   string
		want    []*html.CodeToken
		wantErr bool
		err     string
	}{
		{
			name:  "literal-empty",
			input: `""`,
			want: []*html.CodeToken{
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 1), End: exp.NewPos(1, 2)},
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 2), End: exp.NewPos(1, 3)},
			},
		},
		{
			name:  "literal-$",
			input: `"$"`,
			want: []*html.CodeToken{
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 1), End: exp.NewPos(1, 2)},
				{Kind: html.Literal, Value: `$`, Start: exp.NewPos(1, 2), End: exp.NewPos(1, 3)},
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 3), End: exp.NewPos(1, 4)},
			},
		},
		{
			name:  "code-literal-code-id",
			input: `"Hello, ${name}"`,
			want: []*html.CodeToken{
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 1), End: exp.NewPos(1, 2)},
				{Kind: html.Literal, Value: `Hello, `, Start: exp.NewPos(1, 2), End: exp.NewPos(1, 9)},
				{Kind: html.CodeStart, Value: `${`, Start: exp.NewPos(1, 9), End: exp.NewPos(1, 11)},
				{Kind: html.CodeValue, Value: `name`, Start: exp.NewPos(1, 11), End: exp.NewPos(1, 15)},
				{Kind: html.CodeEnd, Value: `}`, Start: exp.NewPos(1, 15), End: exp.NewPos(1, 16)},
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 16), End: exp.NewPos(1, 17)},
			},
		},
		{
			name:  "code-literal-string",
			input: `"Hello, ${。{name}+'}AAA'}"`,
			want: []*html.CodeToken{
				{Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 1), End: exp.NewPos(1, 2)},
				{Kind: html.Literal, Value: `Hello, `, Start: exp.NewPos(1, 2), End: exp.NewPos(1, 9)},
				{Kind: html.CodeStart, Value: `${`, Start: exp.NewPos(1, 9), End: exp.NewPos(1, 11)},
				// {Kind: html.CodeValue, Value: `{。name}+'}AAA'`, Start: exp.NewPos(1, 11), End: exp.NewPos(1, 25)},
				// {Kind: html.CodeEnd, Value: `}`, Start: exp.NewPos(1, 25), End: exp.NewPos(1, 26)},
				// {Kind: html.BegEnd, Value: `"`, Start: exp.NewPos(1, 26), End: exp.NewPos(1, 27)},
			},
			wantErr: true,
			err:     "[SyntaxError]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := html.NewCodeScanner(exp.NewPos(1, 1), tt.input)
			got, err := tr.GetAllTokens()
			if (err != nil) != tt.wantErr {
				t.Errorf("CodeScanner.GetAllTokens() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				t.Logf("err=%v", err)
				if !strings.Contains(err.Error(), tt.err) {
					t.Errorf("CodeScanner.GetAllTokens() error = %v, want contains %s", err, tt.err)
				}
			}
			for _, i := range got {
				i.Tree = nil // for test DeepEqual
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CodeScanner.GetAllTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
