package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/youthlin/t"
	"github.com/youthlin/t/translator"
)

type Context struct {
	// args
	textTags     []string
	voidElements []string
	tagPrefix    string
	attrPrefix   string
	path         string
	rawpatn      string
	pattern      *regexp.Regexp
	keyword      string
	Keywords     []Keyword
	output       string
	// results
	entries []*translator.Entry
}

// Keyword gettext keyword
type Keyword struct {
	Name    string // 函数名称
	MsgCtxt int    // 上下文参数是第几个
	MsgID   int    // 翻译文本是第几个参数
	MsgID2  int    // 复数是第几个参数
}

// MaxArgIndex 最大参数位置
func (kw Keyword) MaxArgIndex() int {
	var i = kw.MsgCtxt
	if kw.MsgID > i {
		i = kw.MsgID
	}
	if kw.MsgID2 > i {
		i = kw.MsgID2
	}
	return i
}

// parseKeywords gettext;__:1;_n:1,2;_x:1c,2;_xn:1c,2,3
// 解析函数抽取位置，多个位置用分号链接。
// 每个位置的格式是 <函数名>:<xc,y,z>
// 冒号及之后可以省略，表示第一个参数就是翻译文本，如 gettext; __
// 冒号之后可以只有一个位置，表示那个位置就是翻译文本，如 gettext:1; __:1
// 冒号之后可以有两个位置，用逗号分割，分别表示翻译文本和复数文本，如 _n:1,2
// 上下文文本的位置用 'c' 字符后缀表示
// 冒号之后可以有两个位置，用逗号分割，分别表示上下文文本和翻译文本，如 _x:1c,2
// 冒号之后可以有三个位置，用逗号分割，分别表示上下文文本、翻译文本和复数文本，如 _xn:1c,2,3
func parseKeywords(str string) (result []Keyword, err error) {
	kw := strings.Split(str, ";")
	for _, key := range kw {
		// T
		// T:1
		// N:1,2
		// X:1c,2
		// XN:1c,2,3
		nameIndex := strings.Split(key, ":")
		if len(nameIndex) == 1 {
			name := nameIndex[0]
			if name == "" {
				return nil, fmt.Errorf("invalid keywords %q: empty function name", str)
			}
			result = append(result, Keyword{Name: name, MsgID: 1})
			continue
		}
		if len(nameIndex) != 2 {
			return nil, fmt.Errorf("invalid keywords %q: <FunctionName>:<positions>", str)
		}
		name := nameIndex[0]
		if name == "" {
			return nil, fmt.Errorf("invalid keywords %q: empty function name", str)
		}
		k := Keyword{
			Name: nameIndex[0],
		}
		index := strings.Split(nameIndex[1], ",")
		switch len(index) {
		case 1:
			i, err := strconv.ParseInt(index[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf(t.T("invalid keywords: %q. msg id index is not a number %q: %w"),
					str, index[0], err)
			}
			k.MsgID = int(i)
		case 2:
			i1 := index[0]
			i2 := index[1]
			i1c := strings.HasSuffix(i1, "c")
			if i1c {
				c := i1[:len(i1)-1]
				cIndex, err := strconv.ParseInt(c, 10, 64)
				if err != nil {
					return nil, fmt.Errorf(t.T("invalid keywords: %q. context index is not a number %q: %w"),
						str, c, err)
				}
				k.MsgCtxt = int(cIndex)

				index, err := strconv.ParseInt(i2, 10, 64)
				if err != nil {
					return nil, fmt.Errorf(t.T("invalid keywords: %q. msg id index is not a number %q: %w"),
						str, i2, err)
				}
				k.MsgID = int(index)
			} else {
				index, err := strconv.ParseInt(i1, 10, 64)
				if err != nil {
					return nil, fmt.Errorf(t.T("invalid keywords: %q. msg id index is not a number %q: %w"),
						str, i1, err)
				}
				k.MsgID = int(index)

				index, err = strconv.ParseInt(i2, 10, 64)
				if err != nil {
					return nil, fmt.Errorf(t.T("invalid keywords: %q. msg plural index is not a number %q: %w"),
						str, i2, err)
				}
				k.MsgID2 = int(index)
			}
		case 3:
			i1 := index[0]
			i2 := index[1]
			i3 := index[2]
			if !strings.HasSuffix(i1, "c") {
				return nil, fmt.Errorf(t.T("invalid keywords: %q. context index must end with 'c': %q"),
					str, i1)
			}
			c := i1[:len(i1)-1]
			index, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				return nil, fmt.Errorf(t.T("invalid keywords: %q. context index is not a number %q: %w"),
					str, c, err)
			}
			k.MsgCtxt = int(index)

			index, err = strconv.ParseInt(i2, 10, 64)
			if err != nil {
				return nil, fmt.Errorf(t.T("invalid keywords: %q. msg id index is not a number %q: %w"),
					str, i2, err)
			}
			k.MsgID = int(index)

			index, err = strconv.ParseInt(i3, 10, 64)
			if err != nil {
				return nil, fmt.Errorf(t.T("invalid keywords: %q. msg id index is not a number %q: %w"),
					str, i3, err)
			}
			k.MsgID2 = int(index)
		default:
			return nil, fmt.Errorf(t.T("invalid keywords: %q. tow much keyword index"), str)
		}
		result = append(result, k)
	}
	return
}
