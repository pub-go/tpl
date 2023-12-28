package html

// Node 文档树节点
type Node struct {
	Token    *Token  // 节点内容
	End      *Token  // 结束标签(如有)
	Parent   *Node   // 父节点
	Children []*Node // 子节点
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

// GetPreviousSiblingTag 获取前一个 tag 节点
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
