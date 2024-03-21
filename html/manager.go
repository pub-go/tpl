package html

import (
	"fmt"
	"io"
	"io/fs"
	"regexp"
	"strings"

	"code.gopub.tech/errors"
	"code.gopub.tech/tpl/exp"
	"code.gopub.tech/tpl/types"
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
	subFs        string
	textTags     []string
	voidElements []string
	tagPrefix    string
	attrPrefix   string
	indent       string // 缩进字符串 默认="\t"
	maxIndent    int    // 最大缩进层级 默认0=不开启/按原样
	globalScope  exp.Scope
	files        map[string]*Node
	templates    map[string]*Node
}

// NewTplManager 新建一个 HTML 模板管理器
func NewTplManager() *tplManager {
	return (&tplManager{
		files:     map[string]*Node{},
		templates: map[string]*Node{},
	}).
		SetTextTags(GetDefaultTextTags()).
		SetVoidElements(GetDefaultVoidElements()).
		SetTagPrefix(DefaultTagPrefix).
		SetAttrPrefix(DefaultAttrPrefix).
		SetIdent("\t").
		SetGlobalScope(exp.EmptyScope())
}

// SetSubFS 设置文件系统前缀 后续 Parse 时 会首先调用 fs.SubFS
func (m *tplManager) SetSubFS(dir string) *tplManager {
	if dir != "" {
		m.subFs = dir
	}
	return m
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

// SetIdent 设置缩进字符串
// @param s 缩进字符串
func (m *tplManager) SetIdent(s string) *tplManager {
	m.indent = s
	return m
}

// SetMaxIndent 设置最大缩进层级
// @param i 最大缩进层级
func (m *tplManager) SetMaxIndent(i int) *tplManager {
	m.maxIndent = i
	return m
}

// GetIndent 生成缩进字符串
// @param depth 缩进层级
func (m *tplManager) GetIndent(depth int) string {
	if m.maxIndent > 0 {
		depth = depth % m.maxIndent
		return strings.Repeat(m.indent, depth)
	}
	return ""
}

// SetGlobalScope 设置全局的作用域 比如可以注入一些工具函数等
func (m *tplManager) SetGlobalScope(scope exp.Scope) *tplManager {
	m.globalScope = scope
	return m
}

// ParseWithSuffix 使用指定的后缀匹配文件系统中的每个文件
// 匹配的文件将会被解析为模板
func (m *tplManager) ParseWithSuffix(fsys fs.FS, suffix string) error {
	return m.Parse(fsys, func(path string) bool { return strings.HasSuffix(path, suffix) })
}

// ParseWithRegexp 使用正则表达式匹配指定文件系统中的每个文件
// 匹配的文件将会被解析为模板
func (m *tplManager) ParseWithRegexp(fsys fs.FS, pattern *regexp.Regexp) error {
	return m.Parse(fsys, func(path string) bool {
		return pattern.MatchString(path)
	})
}

// Parse 使用指定的函数匹配文件系统中的每个文件
// 匹配的文件将会被解析为模板
func (m *tplManager) Parse(fsys fs.FS, match func(path string) bool) (err error) {
	if m.subFs != "" {
		fsys, err = fs.Sub(fsys, m.subFs)
		if err != nil {
			return err
		}
	}
	return fs.WalkDir(fsys, ".", func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d != nil && !d.IsDir() {
			if match(filePath) {
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
func (m *tplManager) Add(fileName string, reader io.Reader) error {
	if _, ok := m.templates[fileName]; ok {
		return errors.Errorf("failed to add template `%v`: %w",
			fileName, ErrDuplicatedTplName)
	}
	tz := NewHtmlScanner(reader).
		SetTextTags(m.textTags).
		SetAttrPrefix(m.attrPrefix)
	tokens, err := tz.GetAllTokens()
	if err != nil {
		return errors.Errorf("failed to read html tokens in template `%v`: %w", fileName, err)
	}
	p := NewParser().SetVoidElements(m.voidElements)
	tree, err := p.ParseTokens(tokens)
	if err != nil {
		return errors.Errorf("failed to parse html tokens in template `%v`: %w", fileName, err)
	}
	m.files[fileName] = tree
	m.templates[fileName] = tree
	return m.addDefinedTpl(fileName, tree)
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
				return errors.Errorf("failed to add template in %v at %v: %w",
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
			if _, has := m.templates[tplName]; has {
				return errors.Errorf("failed to add template `%v` defined in %v at %v: %w",
					tplName, fileName, attr.ValueStart, ErrDuplicatedTplName)
			}
			m.templates[tplName] = &Node{ // 定义的模板不包括 :define 本身
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

// Files 返回已解析文件的名称
func (m *tplManager) Files() (result map[string]*Node) {
	result = make(map[string]*Node, len(m.files))
	for k, v := range m.files {
		result[k] = v
	}
	return
}

// Files 返回已解析模板的名称（一个文件本身是一个模板，其中还可以定义多个模板）
func (m *tplManager) Templates() (result map[string]*Node) {
	result = make(map[string]*Node, len(m.templates))
	for k, v := range m.templates {
		result[k] = v
	}
	return
}

// GetTemplate implements types.TemplateManager. 获取指定的模板
func (m *tplManager) GetTemplate(name string) (types.Template, error) {
	tree := m.templates[name]
	if tree == nil {
		return nil, errors.Errorf(noSuchTemplate+": %w", name, ErrTplNotFound)
	}
	return NewTemplate(m, name, tree), nil
}
