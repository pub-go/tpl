package html

import (
	"fmt"
	"strings"
)

// Parser 文档树解析器
type Parser struct {
	VoidElements []string
}

// NewParser 新建一个解析器实例
func NewParser() *Parser {
	return new(Parser).SetVoidElements(GetDefaultTextTags())
}

// SetVoidElements 设置空标签
func (p *Parser) SetVoidElements(s []string) *Parser {
	if s != nil {
		p.VoidElements = s
	}
	return p
}

// isVoidElement 判断是否是空标签
func (p *Parser) isVoidElement(tagName string) bool {
	tagName = strings.ToLower(tagName)
	for _, name := range p.VoidElements {
		if strings.ToLower(name) == tagName {
			return true
		}
	}
	return false
}

// ParseTokens 将 tokens 解析为文档树
func (p *Parser) ParseTokens(tokens []*Token) (*Node, error) {
	doc := &Node{}
	node := doc
	for _, token := range tokens {
		switch token.Kind {
		case TokenKindTag:
			tag := token.Tag
			if tag == nil {
				return nil, fmt.Errorf("tag is nil: %v", token)
			}
			isVoid := p.isVoidElement(tag.Name)
			switch {
			case tag.IsClose() || isVoid: // 是闭合标签或空标签
				if tag.IsSelfClose() || isVoid { // 是自闭合标签或空标签
					// <!doctype>, <name />
					// 直接往 node 的子标签列表中添加即可
					node.Children = append(node.Children, &Node{
						Token:  token,
						Parent: node,
					})
				} else { // 是闭合标签 但不是自闭合标签 </div>
					node.End = token   // node 结束
					node = node.Parent // 将 node 指向父节点
					// 示例 <html><head></head><body> 当前处理 </head>
					// </head> 是闭合标签，这时需要让 node 指向 html
					// 以便下一次循环将 body 添加到 html 的字标签列表里
				}
			default: // Tag Open
				n := &Node{
					Token:  token,
					Parent: node,
				}
				node.Children = append(node.Children, n)
				node = n // 进入子节点
				// 示例 <div><span>xxx</span> 当前处理 <span>
				// <span> 是开始标签 让 node 指向它，
				// 以便下一次循环将 xxx 添加到 span 的子节点列表
			}
		case TokenKindComment:
			fallthrough
		case TokenKindCDATA:
			fallthrough
		case TokenKindText:
			node.Children = append(node.Children, &Node{
				Token:  token,
				Parent: node,
			}) // 文本、注释、cdata 直接添加到当前 node 的字标签列表中即可
		default:
			return nil, fmt.Errorf("unexpected error token: %v", token)
		}
	}
	return doc, nil
}
