// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package render

import (
	"bytes"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

// RendererFunc 描述了渲染器函数签名。
type RendererFunc func(n *ast.Node, entering bool) ast.WalkStatus

// ExtRendererFunc 描述了用户自定义的渲染器函数签名。
type ExtRendererFunc func(n *ast.Node, entering bool) (string, ast.WalkStatus)

// Renderer 描述了渲染器接口。
type Renderer interface {
	// Render 渲染输出。
	Render() (output []byte)
}

// Options 描述了渲染选项。
type Options struct {
	// SoftBreak2HardBreak 设置是否将软换行（\n）渲染为硬换行（<br />）。
	SoftBreak2HardBreak bool
	// AutoSpace 设置是否对普通文本中的中西文间自动插入空格。
	// https://github.com/sparanoid/chinese-copywriting-guidelines
	AutoSpace bool
	// RenderListStyle 设置在渲染 OL、UL 时是否添加 data-style 属性 https://github.com/88250/lute/issues/48
	RenderListStyle bool
	// CodeSyntaxHighlight 设置是否对代码块进行语法高亮。
	CodeSyntaxHighlight bool
	// CodeSyntaxHighlightDetectLang bool
	CodeSyntaxHighlightDetectLang bool
	// CodeSyntaxHighlightInlineStyle 设置语法高亮是否为内联样式，默认不内联。
	CodeSyntaxHighlightInlineStyle bool
	// CodeSyntaxHighlightLineNum 设置语法高亮是否显示行号，默认不显示。
	CodeSyntaxHighlightLineNum bool
	// CodeSyntaxHighlightStyleName 指定语法高亮样式名，默认为 "github"。
	CodeSyntaxHighlightStyleName string
	// Vditor 所见即所得支持。
	VditorWYSIWYG bool
	// Vditor 即时渲染支持。
	VditorIR bool
	// Vditor 分屏预览支持。
	VditorSV bool
	// Protyle 所见即所得支持。
	ProtyleWYSIWYG bool
	// KramdownBlockIAL 设置是否打开 kramdown 块级内联属性列表支持。 https://kramdown.gettalong.org/syntax.html#inline-attribute-lists
	KramdownBlockIAL bool
	// KramdownSpanIAL 设置是否打开 kramdown 行级内联属性列表支持。
	KramdownSpanIAL bool
	// SuperBlock 设置是否支持超级块。 https://github.com/88250/lute/issues/111
	SuperBlock bool
	// ImageLazyLoading 设置图片懒加载时使用的图片路径，配置该字段后将启用图片懒加载。
	// 图片 src 的值会复制给新属性 data-src，然后使用该参数值作为 src 的值 https://github.com/88250/lute/issues/55
	ImageLazyLoading string
	// ChineseParagraphBeginningSpace 设置是否使用传统中文排版“段落开头空两格”。
	ChineseParagraphBeginningSpace bool
	// Sanitize 设置是否启用 XSS 安全过滤 https://github.com/88250/lute/issues/51
	// 注意：Lute 目前的实现存在一些漏洞，请不要依赖它来防御 XSS 攻击。
	Sanitize bool
	// FixTermTypo 设置是否对普通文本中出现的术语进行修正。
	// https://github.com/sparanoid/chinese-copywriting-guidelines
	// 注意：开启术语修正的话会默认在中西文之间插入空格。
	FixTermTypo bool
	// Terms 将传入的 terms 合并覆盖到已有的 Terms 字典。
	Terms map[string]string
	// ToC 设置是否打开“目录”支持。
	ToC bool
	// HeadingID 设置是否打开“自定义标题 ID”支持。
	HeadingID bool
	// KramdownIALIDRenderName 设置 kramdown 内联属性列表中出现 id 属性时渲染 id 属性用的 name(key) 名称，默认为 "id"。
	// 仅在 HTML 渲染器 HtmlRenderer 中支持。
	KramdownIALIDRenderName string
	// HeadingAnchor 设置是否对标题生成链接锚点。
	HeadingAnchor bool
	// GFMTaskListItemClass 作为 GFM 任务列表项类名，默认为 "vditor-task"。
	GFMTaskListItemClass string
	// DataTask 设置是否渲染任务列表项的 data-task 属性，用于多状态任务列表。
	DataTask bool
	// VditorCodeBlockPreview 设置 Vditor 代码块是否需要渲染预览部分
	VditorCodeBlockPreview bool
	// VditorMathBlockPreview 设置 Vditor 数学公式块是否需要渲染预览部分
	VditorMathBlockPreview bool
	// VditorHTMLBlockPreview 设置 Vditor HTML 块是否需要渲染预览部分
	VditorHTMLBlockPreview bool
	// LinkBase 设置链接、图片、脚注的基础路径。如果用户在链接或者图片地址中使用相对路径（没有协议前缀且不以 / 开头）并且 LinkBase 不为空则会用该值作为前缀。
	// 比如 LinkBase 设置为 http://domain.com/，对于 ![foo](bar.png) 则渲染为 <img src="http://domain.com/bar.png" alt="foo" />
	LinkBase string
	// LinkPrefix 设置链接、图片的路径前缀。一旦设置该值，链接渲染将强制添加该值作为链接前缀，这有别于 LinkBase。
	// 比如 LinkPrefix 设置为 http://domain.com，对于使用绝对路径的 ![foo](/local/path/bar.png) 则渲染为 <img src="http://domain.com/local/path/bar.png" alt="foo" />；
	// 在 LinkBase 和 LinkPrefix 同时设置的情况下，会先处理 LinkBase 逻辑，最后再在 LinkBase 处理结果上加上 LinkPrefix。
	LinkPrefix string
	// NodeIndexStart 用于设置块级节点编号起始值。
	NodeIndexStart int
	// ProtyleContenteditable 设置 Protyle 渲染时标签中的 contenteditable 属性。
	ProtyleContenteditable bool
	// KeepParagraphBeginningSpace 设置是否保留段首空格
	KeepParagraphBeginningSpace bool
	// NetImgMarker 设置 Protyle 是否标记网络图片
	ProtyleMarkNetImg bool
	// Spellcheck 设置是否启用拼写检查
	Spellcheck bool
	// UnorderedListMarker 设置无序列表和任务列表的标记符，默认为 *
	UnorderedListMarker string
	// ImgTag 设置是否使用 <img> 标签渲染图片
	ImgTag bool
	// PreventEncodeLinkSpace 设置是否阻止将链接中对空格编码为 %20
	PreventEncodeLinkSpace bool
	// ExportNormalizeTaskListMarker 设置是否将非标准的任务列表标记符（如 [/]、[>]、[!] 等）统一导出为完成标记 [X]。
	// 开启后 [ ] 和 [X] 保持不变，其余标记符均转换为 [X]，以兼容不支持自定义标记符的 Markdown 解析器。
	ExportNormalizeTaskListMarker bool
}

func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

// BaseRenderer 描述了渲染器结构。
type BaseRenderer struct {
	Options             *Options                         // 渲染选项
	ParseOptions        *parse.Options                   // 解析选项
	RendererFuncs       map[ast.NodeType]RendererFunc    // 渲染器
	DefaultRendererFunc RendererFunc                     // 默认渲染器，在 RendererFuncs 中找不到节点渲染器时会使用该默认渲染器进行渲染
	ExtRendererFuncs    map[ast.NodeType]ExtRendererFunc // 用户自定义的渲染器
	Writer              *bytes.Buffer                    // 输出缓冲
	LastOut             byte                             // 最新输出的一个字节
	Tree                *parse.Tree                      // 待渲染的树
	DisableTags         int                              // 标签嵌套计数器，用于判断不可能出现标签嵌套的情况，比如语法树允许图片节点包含链接节点，但是 HTML <img> 不能包含 <a>
	FootnotesDefs       []*ast.Node                      // 脚注定义集
	RenderingFootnotes  bool                             // 是否正在渲染脚注定义
}

// NormalizedTaskListItemMarker 返回规范化后的任务列表标记符。
// 当 ExportNormalizeTaskListMarker 选项开启时，将非标准标记符（如 /、>、! 等）转为 X。
// 仅用于 Markdown 文本输出场景（format、export_md），不用于 data-task 属性。
func (r *BaseRenderer) NormalizedTaskListItemMarker(node *ast.Node) byte {
	_ = "STUB: not implemented"
	return 0
}

// NormalizedTaskListItemChecked 返回规范化后的任务列表勾选状态。
func (r *BaseRenderer) NormalizedTaskListItemChecked(node *ast.Node) bool {
	_ = "STUB: not implemented"
	return false
}

// NewBaseRenderer 构造一个 BaseRenderer。
func NewBaseRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) *BaseRenderer {
	_ = "STUB: not implemented"
	return nil
}

// Render 从根节点开始遍历并渲染。
func (r *BaseRenderer) Render() (output []byte) { _ = "STUB: not implemented"; return nil }

func (r *BaseRenderer) renderDefault(n *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// WriteByte 输出一个字节 c。
func (r *BaseRenderer) WriteByte(c byte) { _ = "STUB: not implemented"; return }

// Write 输出指定的字节数组 content。
func (r *BaseRenderer) Write(content []byte) { _ = "STUB: not implemented"; return }

// WriteString 输出指定的字符串 content。
func (r *BaseRenderer) WriteString(content string) { _ = "STUB: not implemented"; return }

// Newline 会在最新内容不是换行符 \n 时输出一个换行符。
func (r *BaseRenderer) Newline() { _ = "STUB: not implemented"; return }

func (r *BaseRenderer) TextAutoSpacePrevious(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *BaseRenderer) TextAutoSpaceNext(node *ast.Node) { _ = "STUB: not implemented"; return }

// 优化排版未处理样式文本 https://github.com/siyuan-note/siyuan/issues/6305

func (r *BaseRenderer) LinkTextAutoSpacePrevious(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *BaseRenderer) LinkTextAutoSpaceNext(node *ast.Node) { _ = "STUB: not implemented"; return }

func SubStr(str string, length int) (ret string) { _ = "STUB: not implemented"; return "" }

func HeadingID(heading *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }

func headingID0(heading *ast.Node) { _ = "STUB: not implemented"; return }

func normalizeHeadingID(heading *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }

type Heading struct {
	ID       string     `json:"id"`
	Box      string     `json:"box"`
	Path     string     `json:"path"`
	HPath    string     `json:"hPath"`
	Content  string     `json:"content"`
	Level    int        `json:"level"`
	Children []*Heading `json:"children"`
	parent   *Heading
}

func (r *BaseRenderer) renderToC(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *BaseRenderer) renderToC0(heading *Heading) { _ = "STUB: not implemented"; return }

func (r *BaseRenderer) Tag(name string, attrs [][]string, selfclosing bool) {
	_ = "STUB: not implemented"
	return
}

func (r *BaseRenderer) headings() (ret []*Heading) { _ = "STUB: not implemented"; return nil }

func parentTip(currentHeading, tip *Heading) *Heading { _ = "STUB: not implemented"; return nil }

func headingText(n *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }

func (r *BaseRenderer) setextHeadingLen(node *ast.Node) (ret int) {
	_ = "STUB: not implemented"
	return 0
}

func (r *BaseRenderer) renderListStyle(node *ast.Node, attrs *[][]string) {
	_ = "STUB: not implemented"
	return
}

func (r *BaseRenderer) tagSrc(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *BaseRenderer) replaceSrc(tokens []byte, src string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (r *BaseRenderer) tagSrcPath(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *BaseRenderer) isLastNode(treeRoot, node *ast.Node) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *BaseRenderer) NodeID(node *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }

func (r *BaseRenderer) NodeAttrs(node *ast.Node) (ret [][]string) {
	_ = "STUB: not implemented"
	return nil
}

func (r *BaseRenderer) NodeAttrsStr(node *ast.Node) (ret string) {
	_ = "STUB: not implemented"
	return ""
}

// languagesNoHighlight 中定义的语言不要进行代码语法高亮。这些代码块会在前端进行渲染，比如各种图表。
var languagesNoHighlight = []string{"mermaid", "echarts", "abc", "graphviz", "mindmap", "flowchart", "plantuml", "infographic"}

func NoHighlight(language string) bool { _ = "STUB: not implemented"; return false }

func (r *BaseRenderer) Text(node *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }

func (r *BaseRenderer) ParagraphContainImgOnly(paragraph *ast.Node) (ret bool) {
	_ = "STUB: not implemented"
	return false
}

func (r *BaseRenderer) needUseHTMLTable(table *ast.Node) (ret bool) {
	_ = "STUB: not implemented"
	return false
}

func RenderHeadingText(n *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }
