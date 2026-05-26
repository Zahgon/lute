// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package parse

import (
	"github.com/88250/lute/ast"
	"github.com/88250/lute/html"
)

func (t *Tree) parseInlineHTML(ctx *InlineContext) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// a closing tag

func (t *Tree) replaceCreateTokens(tokens []byte, caretInTag, caretLeftSpace bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tree) processSpanTag(tags []byte, startTag, endTag string, ctx *InlineContext) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// <span data-type="a">

func (t *Tree) parseCDATA(tokens []byte) (valid bool, remains, content []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (t *Tree) parseDeclaration(tokens []byte) (valid bool, remains, content []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (t *Tree) parseProcessingInstruction(tokens []byte) (valid bool, remains, content []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (t *Tree) parseHTMLComment(tokens []byte) (valid bool, remains, comment []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func TagAttr(tokens []byte) (valid bool, remains, attr, name, val []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil, nil
}

func parseAttrValSpec(tokens []byte) (valid bool, remains, valSpec []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// A double-quoted attribute value consists of ", zero or more characters not including ", and a final ".

// A single-quoted attribute value consists of ', zero or more characters not including ', and a final '.

// An unquoted attribute value is a nonempty string of characters not including whitespace, ", ', =, <, >, or `.

// 大于字符 > 不计入 valSpec

// 属性使用空白分隔

func parseAttrName(tokens []byte) (remains, attrName []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Tree) parseTagName(tokens []byte) (remains, tagName []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetSpanIAL(node *ast.Node, n *html.Node) { _ = "STUB: not implemented"; return }

// 比如设置表格列宽，颜色等

// 设置表格合并单元格

// 合并这两个 IAL

func ContainTextMark(node *ast.Node, dataTypes ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func SetTextMarkNode(node *ast.Node, n *html.Node, options *Options) {
	_ = "STUB: not implemented"
	return
}

// 重新排序，将 a、inline-memo、block-ref、file-annotation-ref、inline-math 放在最前面

// 带有字体样式的公式复制之后内容不正确 https://github.com/siyuan-note/siyuan/issues/6799

// Improve some inline elements Markdown editing https://github.com/siyuan-note/siyuan/issues/9999

// Improve the unescaping of copied block contents https://github.com/siyuan-note/siyuan/issues/16136

// 表格中的代码中带有管道符时使用 HTML 实体替换管道符 Improve the handling of inline-code containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9252

// 下划线中支持包含 Markdown 语法 Improve underline element parsing https://github.com/siyuan-note/siyuan/issues/13768

// 不支持下划线中包含多个元素

func StyleValue(style string) (ret string) { _ = "STUB: not implemented"; return "" }

func startEndBlank(str string) (startBlank, endBlank string) {
	_ = "STUB: not implemented"
	return "", ""
}
