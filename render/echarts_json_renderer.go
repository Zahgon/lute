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
	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

// EChartsJSONRenderer 描述了 ECharts JSON 渲染器。
type EChartsJSONRenderer struct {
	*BaseRenderer
}

// NewEChartsJSONRenderer 创建一个 ECharts JSON 渲染器。
func NewEChartsJSONRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) Renderer {
	_ = "STUB: not implemented"
	return *new(Renderer)
}

func (r *EChartsJSONRenderer) renderKramdownBlockIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderKramdownSpanIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderSup(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderSub(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderBlockQueryEmbed(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderBlockRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderFileAnnotationRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderDefault(n *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderYamlFrontMatter(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderHtmlEntity(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderBackslashContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderBackslash(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderToC(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderFootnotesRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderFootnotesDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderInlineMath(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderMathBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderEmojiImg(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderEmojiUnicode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderEmojiAlias(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderEmoji(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderTableCell(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderTableRow(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderTableHead(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderTable(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderStrikethrough(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderImage(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderLink(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderInlineHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderDocument(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderParagraph(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderCodeSpan(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderEmphasis(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderStrong(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderBlockquote(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderHeading(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderList(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderListItem(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderTaskListItemMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderThematicBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderHardBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderSoftBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) renderCodeBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *EChartsJSONRenderer) leaf(val string, node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) val(val string, node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) openObj() { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) closeObj(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) openChildren(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) closeChildren(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) comma() { _ = "STUB: not implemented"; return }

func (r *EChartsJSONRenderer) ignore(node *ast.Node) bool { _ = "STUB: not implemented"; return false }

// 以下类型的节点不进行渲染，否则图画出来节点太多
