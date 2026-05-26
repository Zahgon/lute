// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

// Package lute 提供了一款结构化的 Markdown 引擎，支持 Go 和 JavaScript。
package lute

import (
	"sync"

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
	"github.com/88250/lute/render"
	"github.com/gopherjs/gopherjs/js"
)

const Version = "1.7.6"

// Lute 描述了 Lute 引擎的顶层使用入口。
type Lute struct {
	ParseOptions  *parse.Options  // 解析选项
	RenderOptions *render.Options // 渲染选项

	HTML2MdRendererFuncs          map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 HTML2Md 渲染器函数
	HTML2VditorDOMRendererFuncs   map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 HTML2VditorDOM 渲染器函数
	HTML2VditorIRDOMRendererFuncs map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 HTML2VditorIRDOM 渲染器函数
	HTML2BlockDOMRendererFuncs    map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 HTML2BlockDOM 渲染器函数
	HTML2VditorSVDOMRendererFuncs map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 HTML2VditorSVDOM 渲染器函数
	Md2HTMLRendererFuncs          map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 Md2HTML 渲染器函数
	Md2VditorDOMRendererFuncs     map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 Md2VditorDOM 渲染器函数
	Md2VditorIRDOMRendererFuncs   map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 Md2VditorIRDOM 渲染器函数
	Md2BlockDOMRendererFuncs      map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 Md2BlockDOM 渲染器函数
	Md2VditorSVDOMRendererFuncs   map[ast.NodeType]render.ExtRendererFunc // 用户自定义的 Md2VditorSVDOM 渲染器函数
}

// New 创建一个新的 Lute 引擎。
//
// 默认启用的解析选项：
//   - GFM 支持
//   - 脚注
//   - 标题自定义 ID
//   - Emoji 别名替换，比如 :heart: 替换为 ❤️
//   - YAML Front Matter
//
// 默认启用的渲染选项：
//   - 软换行转硬换行
//   - 代码块语法高亮
//   - 中西文间插入空格
//   - 修正术语拼写
//   - 标题自定义 ID
func New(opts ...ParseOption) (ret *Lute) { _ = "STUB: not implemented"; return nil }

// Markdown 将 markdown 文本字节数组处理为相应的 html 字节数组。name 参数仅用于标识文本，比如可传入 id 或者标题，也可以传入 ""。
func (lute *Lute) Markdown(name string, markdown []byte) (html []byte) {
	_ = "STUB: not implemented"
	return nil
}

// MarkdownStr 接受 string 类型的 markdown 后直接调用 Markdown 进行处理。
func (lute *Lute) MarkdownStr(name, markdown string) (html string) {
	_ = "STUB: not implemented"
	return ""
}

// Format 将 markdown 文本字节数组进行格式化。
func (lute *Lute) Format(name string, markdown []byte) (formatted []byte) {
	_ = "STUB: not implemented"
	return nil
}

// FormatStr 接受 string 类型的 markdown 后直接调用 Format 进行处理。
func (lute *Lute) FormatStr(name, markdown string) (formatted string) {
	_ = "STUB: not implemented"
	return ""
}

// TextBundle 将 markdown 文本字节数组进行 TextBundle 处理。
func (lute *Lute) TextBundle(name string, markdown []byte, linkPrefixes []string) (textbundle []byte, originalLinks []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TextBundleStr 接受 string 类型的 markdown 后直接调用 TextBundle 进行处理。
func (lute *Lute) TextBundleStr(name, markdown string, linkPrefixes []string) (textbundle string, originalLinks []string) {
	_ = "STUB: not implemented"
	return "", nil
}

// HTML2Text 将指定的 HTMl dom 转换为文本。
func (lute *Lute) HTML2Text(dom string) string { _ = "STUB: not implemented"; return "" }

// RenderJSON 用于渲染 JSON 格式数据。
func (lute *Lute) RenderJSON(markdown string) (json string) { _ = "STUB: not implemented"; return "" }

// Space 用于在 text 中的中西文之间插入空格。
func (lute *Lute) Space(text string) string { _ = "STUB: not implemented"; return "" }

// IsValidLinkDest 判断 str 是否为合法的链接地址。
func (lute *Lute) IsValidLinkDest(str string) bool { _ = "STUB: not implemented"; return false }

func (lute *Lute) GetLinkDest(str string) string { _ = "STUB: not implemented"; return "" }

// GetEmojis 返回 Emoji 别名和对应 Unicode 字符的字典列表。
func (lute *Lute) GetEmojis() (ret map[string]string) { _ = "STUB: not implemented"; return nil }

// PutEmojis 将指定的 emojiMap 合并覆盖已有的 Emoji 字典。
func (lute *Lute) PutEmojis(emojiMap map[string]string) { _ = "STUB: not implemented"; return }

// RemoveEmoji 用于删除 str 中的 Emoji Unicode。
func (lute *Lute) RemoveEmoji(str string) string { _ = "STUB: not implemented"; return "" }

// GetTerms 返回术语字典。
func (lute *Lute) GetTerms() map[string]string { _ = "STUB: not implemented"; return nil }

// PutTerms 将制定的 termMap 合并覆盖已有的术语字典。
func (lute *Lute) PutTerms(termMap map[string]string) { _ = "STUB: not implemented"; return }

var (
	formatRendererSync = render.NewFormatRenderer(nil, nil, nil)
	formatRendererLock = sync.Mutex{}
)

func FormatNodeSync(node *ast.Node, parseOptions *parse.Options, renderOptions *render.Options) (ret string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

var (
	protyleExportMdRendererSync = render.NewProtyleExportMdRenderer(nil, nil, nil)
	protyleExportMdRendererLock = sync.Mutex{}
)

func ProtyleExportMdNodeSync(node *ast.Node, parseOptions *parse.Options, renderOptions *render.Options) (ret string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ProtylePreview 使用指定的 options 渲染 tree 为 Protyle 预览 HTML。
func (lute *Lute) ProtylePreview(tree *parse.Tree, options *render.Options, parseOptions *parse.Options) string {
	_ = "STUB: not implemented"
	return ""
}

// Tree2HTML 使用指定的 options 渲染 tree 为标准 HTML。
func (lute *Lute) Tree2HTML(tree *parse.Tree, options *render.Options, parseOptions *parse.Options) string {
	_ = "STUB: not implemented"
	return ""
}

// ParseOption 描述了解析选项设置函数签名。
type ParseOption func(lute *Lute)

// 以下 Setters 主要是给 JavaScript 端导出方法用。

func (lute *Lute) SetGFMTable(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetGFMTaskListItem(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetArbitraryTaskListItemMarker(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetGFMTaskListItemClass(class string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetDataTask(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetExportNormalizeTaskListMarker(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetGFMStrikethrough(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetGFMStrikethrough1(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetGFMAutoLink(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSoftBreak2HardBreak(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetCodeSyntaxHighlight(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetCodeSyntaxHighlightDetectLang(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetCodeSyntaxHighlightInlineStyle(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetCodeSyntaxHighlightLineNum(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetCodeSyntaxHighlightStyleName(name string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetFootnotes(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetToC(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetHeadingID(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetAutoSpace(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetFixTermTypo(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetEmoji(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetEmojis(emojis map[string]string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetEmojiSite(emojiSite string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetHeadingAnchor(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetTerms(terms map[string]string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetVditorWYSIWYG(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetProtyleWYSIWYG(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetVditorIR(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetVditorSV(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetInlineMath(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetInlineMathAllowDigitAfterOpenMarker(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetLinkPrefix(linkPrefix string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetLinkBase(linkBase string) { _ = "STUB: not implemented"; return }

func (lute *Lute) GetLinkBase() string { _ = "STUB: not implemented"; return "" }

func (lute *Lute) SetVditorCodeBlockPreview(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetVditorMathBlockPreview(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetVditorHTMLBlockPreview(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetRenderListStyle(b bool) { _ = "STUB: not implemented"; return }

// SetSanitize 设置为 true 时表示对输出进行 XSS 过滤。
// 注意：Lute 目前的实现存在一些漏洞，请不要依赖它来防御 XSS 攻击。
func (lute *Lute) SetSanitize(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetImageLazyLoading(dataSrc string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetChineseParagraphBeginningSpace(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetYamlFrontMatter(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSetext(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetBlockRef(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetFileAnnotationRef(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetMark(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetKramdownIAL(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetKramdownBlockIAL(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetKramdownSpanIAL(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetKramdownIALIDRenderName(name string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetTag(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetImgPathAllowSpace(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSuperBlock(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSup(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSub(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetInlineAsterisk(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetInlineUnderscore(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetGitConflict(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetLinkRef(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetIndentCodeBlock(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetDataImage(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetTextMark(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSpin(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetHTML2MarkdownAttrs(attrs []string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetHTMLTag2TextMark(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetParagraphBeginningSpace(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetProtyleMarkNetImg(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetSpellcheck(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetUnorderedListMarker(marker string) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetImgTag(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetPreventEncodeLinkSpace(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetCallout(b bool) { _ = "STUB: not implemented"; return }

func (lute *Lute) SetJSRenderers(options map[string]map[string]*js.Object) {
	_ = "STUB: not implemented"
	return
}

// 稍微进行一点格式校验

// https://go.dev/blog/loopvar-preview
