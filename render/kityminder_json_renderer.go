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

// KityMinderJSONRenderer 描述了 KityMinder JSON 渲染器。
type KityMinderJSONRenderer struct {
	*BaseRenderer
}

// NewKityMinderJSONRenderer 创建一个 KityMinder JSON 渲染器。
func NewKityMinderJSONRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) Renderer {
	_ = "STUB: not implemented"
	return *new(Renderer)
}

func (r *KityMinderJSONRenderer) renderDefault(n *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderBlockQueryEmbed(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderYamlFrontMatter(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderToC(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderMathBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderTable(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderParagraph(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// List.ListItem.Paragraph
// ListItem 下面只有一个段落时不渲染该段落

func (r *KityMinderJSONRenderer) renderBlockquote(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderSuperBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderHeading(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderList(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderListItem(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderThematicBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderCodeBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderKramdownBlockIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) renderDocument(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *KityMinderJSONRenderer) data(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *KityMinderJSONRenderer) openObj() { _ = "STUB: not implemented"; return }

func (r *KityMinderJSONRenderer) closeObj() { _ = "STUB: not implemented"; return }

func (r *KityMinderJSONRenderer) openChildren(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *KityMinderJSONRenderer) closeChildren(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *KityMinderJSONRenderer) comma(node *ast.Node) { _ = "STUB: not implemented"; return }

func headingChildren(heading *ast.Node) (ret []*ast.Node) { _ = "STUB: not implemented"; return nil }
