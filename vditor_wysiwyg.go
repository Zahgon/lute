// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package lute

import (
	"bytes"

	"github.com/88250/lute/html"
	"github.com/88250/lute/html/atom"
	"github.com/88250/lute/parse"
)

// Md2HTML 将 markdown 转换为标准 HTML，用于源码模式预览。
func (lute *Lute) Md2HTML(markdown string) (sHTML string) { _ = "STUB: not implemented"; return "" }

// SpinVditorDOM 自旋 Vditor DOM，用于所见即所得模式下的编辑。
func (lute *Lute) SpinVditorDOM(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

// HTML2VditorDOM 将 HTML 转换为 Vditor DOM，用于所见即所得模式下粘贴。
func (lute *Lute) HTML2VditorDOM(sHTML string) (vHTML string) { _ = "STUB: not implemented"; return "" }

// VditorDOM2HTML 将 Vditor DOM 转换为 HTML，用于 Vditor.getHTML() 接口。
func (lute *Lute) VditorDOM2HTML(vhtml string) (sHTML string) { _ = "STUB: not implemented"; return "" }

// Md2VditorDOM 将 markdown 转换为 Vditor DOM，用于从源码模式切换至所见即所得模式。
func (lute *Lute) Md2VditorDOM(markdown string) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

// VditorDOM2Md 将 Vditor DOM 转换为 markdown，用于从所见即所得模式切换至源码模式。
func (lute *Lute) VditorDOM2Md(htmlStr string) (markdown string) {
	_ = "STUB: not implemented"
	return ""
}

// RenderEChartsJSON 用于渲染 ECharts JSON 格式数据。
func (lute *Lute) RenderEChartsJSON(markdown string) (json string) {
	_ = "STUB: not implemented"
	return ""
}

// RenderKityMinderJSON 用于渲染 KityMinder JSON 格式数据。
func (lute *Lute) RenderKityMinderJSON(markdown string) (json string) {
	_ = "STUB: not implemented"
	return ""
}

// HTML2Md 用于将 HTML 转换为 markdown。
func (lute *Lute) HTML2Md(html string) (markdown string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) vditorDOM2Md(htmlStr string) (markdown string) {
	_ = "STUB: not implemented"
	// 删掉插入符
	return ""
}

// 替换结尾空白，否则 HTML 解析会产生冗余节点导致生成空的代码块

// 将字符串解析为 DOM 树

// 调整 DOM 结构

// 将 HTML 树转换为 Markdown AST

// 调整树结构

// 合并代码节点 https://github.com/Vanessa219/vditor/issues/167

// 浏览器生成的子列表是 ul.ul 形式，需要将其调整为 ul.li.ul

// 将 AST 进行 Markdown 格式化渲染

func (lute *Lute) adjustVditorDOM(root *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) adjustBlockInTable(n *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) adjustTag(n *html.Node) { _ = "STUB: not implemented"; return }

// 将某些自定义标签转换为标准标签

// 将非表格内的文本节点中的 <br> 转换为 \n https://github.com/siyuan-note/siyuan/issues/15373

func (lute *Lute) adjustNoscriptImg(n *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) adjustMath(n *html.Node) { _ = "STUB: not implemented"; return }

// 根据最后 4 个换行符分隔公式内容，适配 CSDN 公式剪藏

func (lute *Lute) adjustTableCode(n *html.Node) { _ = "STUB: not implemented"; return }

// 表格类型的代码块进行预处理 https://github.com/siyuan-note/siyuan/issues/11540
// 移除 <td class="gutter">
// td class="code"> 下的 <div class="container"> 改为 <pre>

// 移除 <span class="lnt"> 格式的行号 https://github.com/siyuan-note/siyuan/issues/13242

func (lute *Lute) mergeSameStrong(n *html.Node) { _ = "STUB: not implemented"; return }

// adjustVditorDOMListList 用于将 ul.ul 调整为 ul.li.ul。
func (lute *Lute) adjustVditorDOMListList(n *html.Node) { _ = "STUB: not implemented"; return }

// 规范化换行时 li 的结构，对调 ZWSP 和 <br> 的位置

func (lute *Lute) removeHighlightJSSpans(node *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) hljsSpans(n *html.Node, spans *[]*html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) removeEmptyNodes(node *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) searchEmptyNodes(n *html.Node, emptyNodes *[]*html.Node) {
	_ = "STUB: not implemented"
	return
}

// 前节点或者后节点是行级节点的话保留该空白

// 前节点是加粗节点的话去掉 ZWSP https://github.com/siyuan-note/siyuan/issues/16896

// 浏览器剪藏扩展列表下方段落缩进成为子块 https://github.com/siyuan-note/siyuan/issues/6289

// 如果行级标记节点最后一个子节点是 <br>，则将该 <br> 移动到该行级标记节点的后面
// 表格内多个连续的超链接无法换行显示 https://github.com/siyuan-note/siyuan/issues/5966

// 将嵌套在临时标记中的节点提升到临时标记节点之前

func (lute *Lute) removeWbr(n *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) searchWbr(n *html.Node, wbrNodes *[]*html.Node) {
	_ = "STUB: not implemented"
	return
}

func (lute *Lute) mergeVditorDOMList0(n *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) adjustVditorDOMListTight0(n *html.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) adjustVditorDOMListItemInP(n *html.Node) { _ = "STUB: not implemented"; return }

// li 换行时 id 重复需要重新生成

// 松散 li 换行时和上一个 li.last id 重复

// id 为空的话是行级节点，列表项行级排版自动换行问题 https://github.com/siyuan-note/siyuan/issues/379

// 在 li 下的每个非容器块节点用 p 包裹

func (lute *Lute) removeCodeCode(n *html.Node) { _ = "STUB: not implemented"; return }

// code.code 重复嵌套，则不处理外层 code

func (lute *Lute) adjustVditorDOMCodeA(n *html.Node) {
	_ = "STUB: not implemented"
	// https://github.com/siyuan-note/siyuan/issues/11370
	return
}

// code.a 的情况将 a 移到 code 外层，即 a.code

// forwardNextBlock 向前移动至下一个块级节点，即跳过行级节点。
func (lute *Lute) forwardNextBlock(spanNode *html.Node) (spans []*html.Node, nextBlock *html.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lute *Lute) listItemEnter(li *html.Node) bool { _ = "STUB: not implemented"; return false }

func (lute *Lute) isTightList(list *html.Node) string { _ = "STUB: not implemented"; return "" }

// genASTByVditorDOM 根据指定的 Vditor DOM 节点 n 进行深度优先遍历并逐步生成 Markdown 语法树 tree。
func (lute *Lute) genASTByVditorDOM(n *html.Node, tree *parse.Tree) {
	_ = "STUB: not implemented"
	return
}

// 1：浮动工具栏，2：preview 代码块、数学公式块

// 去掉列表项标记符 1.

// 子有序列表第一项必须从 1 开始

// 处理结尾换行

// 开头结尾空格后会形成 * foo * 导致强调、加粗删除线标记失效，这里将空格移到右标记符前后 _*foo*_

// 处理结尾换行

// 处理结尾换行

// 处理结尾换行

// 所见即所得行级 HTML 解析 https://github.com/Vanessa219/vditor/issues/1156

// 删掉表格中结尾的 br

// 仅允许 input 出现在任务列表中

// 在任务列表前退格

// ul.li.input

// ul.li.p.input

// 处理结尾换行

// 处理结尾换行

// 图片引用风格 ![text][label]

func (lute *Lute) hasAttr(n *html.Node, attrName string) bool {
	_ = "STUB: not implemented"
	return false
}

func (lute *Lute) domParent(n *html.Node, dataAtom atom.Atom) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func (lute *Lute) domChild(n *html.Node, dataAtom atom.Atom) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func (lute *Lute) domChild0(n *html.Node, dataAtom atom.Atom) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func (lute *Lute) setDOMAttrValue(n *html.Node, attrName, attrVal string) {
	_ = "STUB: not implemented"
	return
}

func (lute *Lute) removeDOMAttr(n *html.Node, attrName string) { _ = "STUB: not implemented"; return }

func (lute *Lute) domCode(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func (lute *Lute) domCode0(n *html.Node, buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (lute *Lute) parentIs(n *html.Node, parentTypes ...atom.Atom) bool {
	_ = "STUB: not implemented"
	return false
}

func (lute *Lute) getParent(n *html.Node, parentType atom.Atom) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func (lute *Lute) isCaret(n *html.Node) (isCaret, isEmptyText bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (lute *Lute) isEmptyText(n *html.Node) bool { _ = "STUB: not implemented"; return false }

func (lute *Lute) startsWithNewline(n *html.Node) bool { _ = "STUB: not implemented"; return false }

func (lute *Lute) isInline(n *html.Node) bool { _ = "STUB: not implemented"; return false }

func (lute *Lute) prefixSpaces(text string) (ret string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) suffixSpaces(text string) (ret string) { _ = "STUB: not implemented"; return "" }
