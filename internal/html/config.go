package html

const (
	TagStart = '<'  // 标签开头
	NewLine  = '\n' // 换行
)
const (
	DefaultTagPrefix  = "t:" // 默认的标签名前缀
	DefaultAttrPrefix = ":"  // 默认的属性名前缀
)

// GetDefaultTextTags 默认的只包含文本的标签
func GetDefaultTextTags() []string {
	return []string{
		"script", "style", "textarea", "title",
	}
}

// GetDefaultVoidElements 默认的空标签 不含闭合斜线 不含内容
func GetDefaultVoidElements() []string {
	return []string{
		"!doctype", // 特殊
		"area", "base", "br", "col", "embed", "hr", "img",
		"input", "link", "meta", "source", "track", "wbr",
	}
}
