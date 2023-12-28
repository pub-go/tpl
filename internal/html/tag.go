package html

import (
	"sort"
	"strings"
)

// Tag <标签> 如果是结束标签则 Name 以 / 开头
type Tag struct {
	Name   string
	Attrs  []*Attr
	sorted int
	attrs  map[string]*Attr
}

// AttrMap 构造属性 Map 用于快速查找
func (t *Tag) AttrMap() map[string]*Attr {
	if t.attrs == nil || len(t.attrs) != len(t.Attrs) {
		m := make(map[string]*Attr, len(t.Attrs))
		for _, attr := range t.Attrs {
			m[attr.Name] = attr
		}
		t.attrs = m
	}
	return t.attrs
}

func (t *Tag) SortedAttr(prefix string) []*Attr {
	if t.sorted != len(t.Attrs) {
		weight := map[string]int{
			"if":     -3,
			"range":  -2,
			"remove": -1,
		}
		sort.SliceStable(t.Attrs, func(i, j int) bool {
			x, y := t.Attrs[i].Name, t.Attrs[j].Name
			xw, yw := 1, 1
			if strings.HasPrefix(x, prefix) {
				x = strings.TrimPrefix(x, prefix)
				xw = 0
			}
			if strings.HasPrefix(y, prefix) {
				y = strings.TrimPrefix(y, prefix)
				yw = 0
			}
			if xw < yw {
				return true // x 有 prefix, y 没有 prefix, 则 x 排在前
			}
			if xw == 1 && yw == 1 {
				return false // 不是指令属性 保持原来顺序
			}
			return weight[x] < weight[y]
		})
		t.sorted = len(t.Attrs)
	}
	return t.Attrs
}

// IsClose 是否是闭合标签
// 包含两种：正常的闭合标签 </name>, 自闭合标签 <name />
func (t *Tag) IsClose() bool {
	// </name>
	if strings.HasPrefix(t.Name, "/") {
		return true
	}
	return t.IsSelfClose()
}

// IsSelfClose 是否是自闭合标签
// 注意 <meta> 这种 void 标签不属于自闭合标签 只有 <name/> 这种以 / 结尾的才算
// 包含三种：无属性的 <name/>; 无属性值的 <name />; 有属性值的 <name k=v/>
func (t *Tag) IsSelfClose() bool {
	as := len(t.Attrs)
	if as == 0 {
		// <xx/>
		return strings.HasSuffix(t.Name, "/")
	}
	a := t.Attrs[as-1]
	if a.Value == nil {
		// <xx />
		// <xx n/>
		return strings.HasSuffix(a.Name, "/")
	}
	// <xx k=v/>
	return strings.HasSuffix(*a.Value, "/")
}

func (t *Tag) String() string {
	var sb strings.Builder
	sb.WriteString("<")
	sb.WriteString(t.Name)
	for _, attr := range t.Attrs {
		sb.WriteString(" ")
		sb.WriteString(attr.String())
	}
	sb.WriteString(">")
	return sb.String()
}
