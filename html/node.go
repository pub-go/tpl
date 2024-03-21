package html

// Node 文档树节点
type Node struct {
	Token      *Token // 节点内容
	End        *Token // 结束标签(如有)
	openIndent bool
	endIndent  bool
	Parent     *Node   // 父节点
	Children   []*Node // 子节点
}

// IsBlankText 当前节点是否是空白文本
func (n *Node) IsBlankText() bool {
	if n != nil && n.Token != nil {
		return n.Token.IsBlankText()
	}
	return false
}

// GetPreviousSibling 获取前一个兄弟节点
func (n *Node) GetPreviousSibling() (pre *Node) {
	if n.Parent == nil {
		return nil
	}
	var p *Node
	for _, child := range n.Parent.Children {
		if child == n {
			pre = p
			break
		}
		p = child
	}
	return
}

func (n *Node) GetLastChild() (node *Node) {
	if len(n.Children) > 0 {
		return n.Children[len(n.Children)-1]
	}
	return
}

// GetPreviousSiblingTag 获取前一个 tag 节点
// 可用于 :else 之前是否 :if 检测
func (n *Node) GetPreviousSiblingTag() (pre *Node) {
	if n.Parent == nil {
		return nil
	}
	var p *Node
	for _, child := range n.Parent.Children {
		if child == n {
			pre = p
			break
		}
		if child.Token != nil && child.Token.Kind == TokenKindTag {
			p = child
		}
	}
	return
}

// GetNextSibling 获取下一个兄弟节点
// 可用于获取 :range 节点之后的空白文本节点
func (n *Node) GetNextSibling() (next *Node) {
	if n.Parent == nil {
		return nil
	}
	var p *Node
	for _, child := range n.Parent.Children {
		if p == n {
			next = child
			break
		}
		p = child
	}
	return
}

// GetChildrenWithoutHeadTailBlankText 获取子节点 但不要第一个和最后一个空白节点(如果不是空白节点就要)
// 用于
//
//	<template :define="name">
//	  <p>blabla</p>
//	</template>
//
// insert 和 replace 时去除多余的空白
func (n *Node) GetChildrenWithoutHeadTailBlankText() (result []*Node) {
	count := len(n.Children)
	for i, child := range n.Children {
		if (i == 0 || i == count-1) && child.IsBlankText() {
			continue
		}
		result = append(result, child)
	}
	return
}
