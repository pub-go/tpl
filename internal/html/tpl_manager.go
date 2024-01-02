package html

import (
	"fmt"
	"io"
	"io/fs"
	"path"

	"code.gopub.tech/tpl/internal/exp"
	"code.gopub.tech/tpl/internal/types"
)

var (
	// ErrDuplicatedTplName 重复的模板名称
	ErrDuplicatedTplName = fmt.Errorf("duplicated template name")
	// ErrTplNotFound 模板未找到
	ErrTplNotFound = fmt.Errorf("template not found")
	// ErrNilTag 找不到 tag 信息
	ErrNilTag = fmt.Errorf("unexpected nil tag")
)
var _ types.TemplateManager = (*tplManager)(nil)

// tplManager HTML 模板管理器
type tplManager struct {
	textTags     []string
	voidElements []string
	tagPrefix    string
	attrPrefix   string
	globalScope  exp.Scope
	files        map[string]*Node
}

// NewTplManager 新建一个 HTML 模板管理器
func NewTplManager() *tplManager {
	return (&tplManager{
		files: map[string]*Node{},
	}).
		SetTextTags(GetDefaultTextTags()).
		SetVoidElements(GetDefaultVoidElements()).
		SetTagPrefix(DefaultTagPrefix).
		SetAttrPrefix(DefaultAttrPrefix).
		SetGlobalScope(exp.EmptyScope())
}

// SetTextTags 设置只包含文本的标签
// @param textTags 只包含文本的标签, 如 script, title, style, textarea
func (m *tplManager) SetTextTags(textTags []string) *tplManager {
	if textTags != nil {
		m.textTags = textTags
	}
	return m
}

// SetVoidElements 设置空标签
// @param voidElements 空标签不包含闭合斜线也不包含内容, 如 meta, br, img
func (m *tplManager) SetVoidElements(voidElements []string) *tplManager {
	if voidElements != nil {
		m.voidElements = voidElements
	}
	return m
}

// SetTagPrefix 设置标签前缀
// @param prefix 要设置的前缀
func (m *tplManager) SetTagPrefix(prefix string) *tplManager {
	if prefix != "" {
		m.tagPrefix = prefix
	}
	return m
}

// SetAttrPrefix 设置属性前缀
// @param prefix 要设置的前缀
func (m *tplManager) SetAttrPrefix(prefix string) *tplManager {
	if prefix != "" {
		m.attrPrefix = prefix
	}
	return m
}

// SetGlobalScope 设置全局的作用域 比如可以注入一些工具函数等
func (m *tplManager) SetGlobalScope(scope exp.Scope) *tplManager {
	m.globalScope = scope
	return m
}

// Glob 加载文件系统 fsys 中，文件名匹配 pattern 的文件
func (m *tplManager) Glob(fsys fs.FS, pattern string) error {
	return m.Parse(fsys, func(s string) (bool, error) {
		return path.Match(pattern, s)
	})
}

// Parse 加载文件系统 fsys 中，文件名匹配 match 函数的文件
func (m *tplManager) Parse(fsys fs.FS, match func(path string) (bool, error)) error {
	return fs.WalkDir(fsys, ".", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d != nil && !d.IsDir() {
			ok, err := match(filePath)
			if err != nil {
				return err
			}
			if ok {
				file, err := fsys.Open(filePath)
				if err != nil {
					return err
				}
				if err := m.Add(filePath, file); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// Add 添加模板文件
// @param name 文件名
// @param reader 文件内容
func (m *tplManager) Add(name string, reader io.Reader) error {
	if _, ok := m.files[name]; ok {
		return fmt.Errorf("failed to add template `%v`: %w",
			name, ErrDuplicatedTplName)
	}
	tz := NewHtmlScanner(reader).
		SetTextTags(m.textTags).
		SetAttrPrefix(m.attrPrefix)
	tokens, err := tz.GetAllTokens()
	if err != nil {
		return fmt.Errorf("failed to read html tokens: %w", err)
	}
	p := NewParser().SetVoidElements(m.voidElements)
	tree, err := p.ParseTokens(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse html tokens: %w", err)
	}
	m.files[name] = tree
	return m.addDefinedTpl(name, tree)
}

// addDefinedTpl 查找模板文件中定义的可复用模板
func (m *tplManager) addDefinedTpl(fileName string, tree *Node) error {
	var (
		token      = tree.Token
		attrPrefix = m.attrPrefix
	)
	if token != nil {
		switch token.Kind {
		case TokenKindTag:
			// <p :define="name">xxx</p>
			// <div :insert="name">yyy</div>
			// -->
			// <div>xxx</div>
			tag := token.Tag
			if tag == nil {
				return fmt.Errorf("failed to add template in %v at %v %w",
					fileName, token.Start, ErrNilTag)
			}
			attr, ok := tag.AttrMap()[attrPrefix+attrDefine]
			if !ok {
				break
			}
			tplName, err := attr.Evaluate(exp.EmptyScope())
			if err != nil {
				return err
			}
			if _, has := m.files[tplName]; has {
				return fmt.Errorf("failed to add template `%v` defined in %v at %v: %w",
					tplName, fileName, attr.ValueStart, ErrDuplicatedTplName)
			}
			m.files[tplName] = &Node{ // 定义的模板不包括 :define 本身
				Children: tree.GetChildrenWithoutHeadTailBlankText(),
			}
		}
	}
	for _, child := range tree.Children {
		if err := m.addDefinedTpl(fileName, child); err != nil {
			return err
		}
	}
	return nil
}

// Files 返回已解析模板的名称
func (m *tplManager) Files() (file []string) {
	for key := range m.files {
		file = append(file, key)
	}
	return
}

// GetTemplate implements types.TemplateManager. 获取指定的模板
func (m *tplManager) GetTemplate(name string) (types.Template, error) {
	tree := m.files[name]
	if tree == nil {
		return nil, fmt.Errorf(noSuchTemplate+": %w", name, ErrTplNotFound)
	}
	return NewTemplate(m, name, tree), nil
}
