package html

// Node 文档树节点
type Node struct {
	Token    *Token  // 节点内容
	End      *Token  // 结束标签(如有)
	Parent   *Node   // 父节点
	Children []*Node // 子节点
}
