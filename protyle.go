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
	"github.com/88250/lute/ast"
	"github.com/88250/lute/html"
	"github.com/88250/lute/parse"
	"github.com/88250/lute/render"
)

func (lute *Lute) SpinBlockDOM(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	// fmt.Println(ivHTML)
	return ""
}

// 软换行后生成多个块，需要把老 ID 调整到第一个块上

// 空段落块还原

// 使用 Markdown 标记符嵌套行级元素后被还原为纯文本 https://github.com/siyuan-note/siyuan/issues/7637
// 这里需要将混合嵌套（比如 <strong><span a></span></strong>）的行级元素拆分为多个平铺的行级元素（<span strong> 和 <span strong a>）

func (lute *Lute) HTML2BlockDOM(sHTML string) (vHTML string) {
	_ = "STUB: not implemented"
	// fmt.Println(sHTML)
	return ""
}

func (lute *Lute) BlockDOM2HTML(vHTML string) (sHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) BlockDOM2InlineBlockDOM(vHTML string) (vIHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) Md2BlockDOM(markdown string, reserveEmptyParagraph bool) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) Md2BlockDOMWithAutoLink(markdown string, reserveEmptyParagraph bool) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) Md2BlockDOMTree(markdown string, reserveEmptyParagraph bool) (vHTML string, tree *parse.Tree) {
	_ = "STUB: not implemented"
	return "", nil
}

// 先将 TextMark 转换为 Inlines https://github.com/siyuan-note/siyuan/issues/13056

func (lute *Lute) InlineMd2BlockDOM(markdown string) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) BlockDOM2Md(htmlStr string) (kramdown string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) BlockDOM2StdMd(htmlStr string) (markdown string) {
	_ = "STUB: not implemented"
	return ""
}

// DOM 转 AST

// 将 kramdown IAL 节点内容置空

func (lute *Lute) BlockDOM2Text(htmlStr string) (text string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) BlockDOM2TextLen(htmlStr string) int { _ = "STUB: not implemented"; return 0 }

func (lute *Lute) BlockDOM2Content(htmlStr string) (text string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) BlockDOM2EscapeMarkerContent(htmlStr string) (text string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) Tree2BlockDOM(tree *parse.Tree, options *render.Options, parseOptions *parse.Options) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) RenderNodeBlockDOM(node *ast.Node) string { _ = "STUB: not implemented"; return "" }

func (lute *Lute) BlockDOM2Tree(htmlStr string) (ret *parse.Tree) {
	_ = "STUB: not implemented"
	return nil
}

// Improve `inline code` markdown editing https://github.com/siyuan-note/siyuan/issues/9978
// spinBlockDOMTests #212

// 替换结尾空白，否则 HTML 解析会产生冗余节点导致生成空的代码块

// 将字符串解析为 DOM 树

// 调整 DOM 结构

// 将 HTML 树转换为 Markdown AST

// 调整树结构

// 合并代码节点 https://github.com/Vanessa219/vditor/issues/167

func (lute *Lute) MergeSameTextMark(n *ast.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) MergeSameSpan(n *ast.Node) { _ = "STUB: not implemented"; return }

// open marker
// close marker

func (lute *Lute) CancelSuperBlock(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) CancelList(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) CancelBlockquote(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) CancelCallout(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) Blocks2Ps(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) Blocks2Hs(ivHTML, level string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) OL2TL(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) UL2TL(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) TL2OL(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

// task marker

func (lute *Lute) TL2UL(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

// task marker

func (lute *Lute) OL2UL(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) UL2OL(ivHTML string) (ovHTML string) { _ = "STUB: not implemented"; return "" }

func (lute *Lute) Callout2Blockquote(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) Blockquote2Callout(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	return ""
}

// 标记符 >

// 第一个段落 [!TYPE]
// 第一个段落的 IAL

func (lute *Lute) blockDOM2Md(htmlStr string) (markdown string) {
	_ = "STUB: not implemented"
	return ""
}

// 将 AST 进行 Markdown 格式化渲染

func (lute *Lute) genASTByBlockDOM(n *html.Node, tree *parse.Tree) {
	_ = "STUB: not implemented"
	return
}

// Custom dom, which will be omitted when build tree https://github.com/88250/lute/issues/206

// 任务列表

// 嵌入块中存在换行 SQL 语句时会被转换为段落文本 https://github.com/siyuan-note/siyuan/issues/5728

// Allow changing headings with `#` https://github.com/siyuan-note/siyuan/issues/7924

func (lute *Lute) genASTContenteditable(n *html.Node, tree *parse.Tree) {
	_ = "STUB: not implemented"
	return
}

/* 外部内容粘贴到表格中后编辑导致换行丢失  https://github.com/siyuan-note/siyuan/issues/7501 */
/* 表格内存在行级公式时编辑会产生换行 https://github.com/siyuan-note/siyuan/issues/2279 */

// After entering `\` in the table, the next column is merged incorrectly https://github.com/siyuan-note/siyuan/issues/7817

// 仅包含转义字符时转义自身 \

// 合并相邻的代码

// 叠加代码

// 表格单元格中使用代码和 `|` 的问题 https://github.com/siyuan-note/siyuan/issues/4717

// 表格单元格中使用代码和 `|` 的问题 https://github.com/siyuan-note/siyuan/issues/4717

// `<kbd>` 中反斜杠转义问题 https://github.com/siyuan-note/siyuan/issues/2242

// `<kbd>` 无法删除 https://github.com/siyuan-note/siyuan/issues/4162

// 某些情况下复制过来的 DOM 是该情况，这里按纯文本解析

// 给文字和图片同时设置字体格式后图片丢失 https://github.com/siyuan-note/siyuan/issues/6297

// 行级元素前输入转义符 `\` 导致异常 https://github.com/siyuan-note/siyuan/issues/6237

// 开头结尾空格后会形成 * foo * 导致强调、加粗删除线标记失效，这里将空格移到右标记符前后 _*foo*_

// 丢弃没有锚文本的链接

//n.FirstChild.NextSibling.FirstChild.NextSibling

// < 和 > 符号不用转义，可以符合 Markdown 规范 https://github.com/siyuan-note/siyuan/issues/15023

// 转义字符加行级样式后继续输入会出现标记符 https://github.com/siyuan-note/siyuan/issues/6134

// TextMark 节点

// 将硬换行转换为软换行 https://github.com/siyuan-note/siyuan/issues/14481

// 处理结尾换行

// 开头结尾空格后会形成 * foo * 导致强调、加粗删除线标记失效，这里将空格移到右标记符前后 _*foo*_

// 转义符导致的行级元素样式属性暴露 https://github.com/siyuan-note/siyuan/issues/2969

// foo\‸**bar**

// foo\**bar**

// 处理结尾换行

// 处理结尾换行

// 处理结尾换行

// < 和 > 符号不用转义，可以符合 Markdown 规范 https://github.com/siyuan-note/siyuan/issues/15023

func (lute *Lute) setBlockIAL(n *html.Node, node *ast.Node) (ialTokens []byte) {
	_ = "STUB: not implemented"
	return nil
}

// setBlockIAL2 设置块级 IAL，仅用于 h2m 中。
func (lute *Lute) setBlockIAL2(n *html.Node, node *ast.Node) (ialTokens []byte) {
	_ = "STUB: not implemented"
	return nil
}

func processSpanMarkerSpace(n *html.Node, node *ast.Node) { _ = "STUB: not implemented"; return }

func (lute *Lute) removeInnerMarker(n *html.Node, marker string) { _ = "STUB: not implemented"; return }

func (lute *Lute) removeInnerMarker0(n *html.Node, marker string) {
	_ = "STUB: not implemented"
	return
}

// 自定义标签

func (lute *Lute) removeTempMark(dataType string) (ret string) {
	_ = "STUB: not implemented"
	return ""
}
