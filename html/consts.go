package html

const (
	DefaultTagPrefix  = "t:" // 默认的标签名前缀
	DefaultAttrPrefix = ":"  // 默认的属性名前缀
)

const (
	tagNameBlock = "block"   // <block>只输出子内容的标签</block>
	attrWith     = "with"    // <div :with="a := ${cond}">满足条件时才会输出本节点</div>
	attrIf       = "if"      // <div :if="${cond}">满足条件时才会输出本节点</div>
	attrElse_If  = "else-if" // 前一个tag必须含 if 属性，前一个节点不满足条件时会判断本节点
	attrElseIf   = "elseif"  // 同上
	attrElIf     = "elif"    // 同上
	attrElse     = "else"    // 前一个tag必须含条件属性，前一个节点不满足条件时会判断本节点
	attrRemove   = "remove"  // <ul :remove="all-but-first"></ul> 移除 tag 本身或子内容或两者或仅保留第一个子tag
	attrRange    = "range"   // <li :range="idx, item : items"></li> 循环输出本节点
	attrText     = "text"    // <p :text="Hello, ${name}">placeholder</p> 动态替换文本
	attrRaw      = "raw"     // 同 text 但不对 HTML 转义
	attrDefine   = "define"  // <template :define="name">xxx</template> 定义一个可复用组件(不含 tag 本身)
	attrInsert   = "insert"  // <div :insert="foo">ignored</div> 用指定的组件作为本节点的内容
	attrReplace  = "replace" // <div :replace="foo">ignored</div> 用指定的组件替换本节点
)

const (
	// 移除 tag 本身和子内容
	removeAll1, removeAll2 = `"all"`, `'all'`
	// 保留 tag 本身，移除 tag 的子内容
	removeBody1, removeBody2 = `"body"`, `'body'`
	// 移除 tag 本身，但保留子内容
	removeTag1, removeTag2 = `"tag"`, `'tag'`
	// 保留 tag 本身，保留 tag 的第一个子 tag
	removeAllButFirst1, removeAllButFirst2 = `"all-but-first"`, `'all-but-first'`
)

const textTrue = "true"

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

const (
	noSuchTemplate      = "no such template `%v`"
	attrShouldHaveValue = "attribute `%s` should have value (at position %v)"
)
