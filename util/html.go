// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"bytes"

	"github.com/88250/lute/html"
	"github.com/88250/lute/html/atom"
)

func ParseHTML(htmlStr string) *html.Node { _ = "STUB: not implemented"; return nil }

// doc.html.body

func GetTextMarkTextDataWithoutEscapeQuote(n *html.Node) (content string) {
	_ = "STUB: not implemented"
	return ""
}

// 粘贴 Markdown 时行级元素中的双引号不再转换为实体 https://github.com/siyuan-note/siyuan/issues/14503

func GetTextMarkTextData(n *html.Node) (content string) { _ = "STUB: not implemented"; return "" }

func GetTextMarkInlineMemoData(n *html.Node) (content string) { _ = "STUB: not implemented"; return "" }

func GetTextMarkAData(n *html.Node) (href, title string) { _ = "STUB: not implemented"; return "", "" }

// < 和 > 符号不用转义，可以符合 Markdown 规范 https://github.com/siyuan-note/siyuan/issues/15023

func GetTextMarkInlineMathData(n *html.Node) (content string) { _ = "STUB: not implemented"; return "" }

func GetTextMarkBlockRefData(n *html.Node) (id, subtype string) {
	_ = "STUB: not implemented"
	return "", ""
}

func GetTextMarkFileAnnotationRefData(n *html.Node) (id string) {
	_ = "STUB: not implemented"
	return ""
}

func GetFormula(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func DomChildByTypeAndClass(n *html.Node, dataAtom atom.Atom, class ...string) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func DomChildrenByType(n *html.Node, dataAtom atom.Atom) (ret []*html.Node) {
	_ = "STUB: not implemented"
	// 递归遍历所有子节点
	return nil
}

func DomExistChildByType(n *html.Node, dataAtom ...atom.Atom) bool {
	_ = "STUB: not implemented"
	return false
}

func domChildByType(n *html.Node, dataAtom atom.Atom) *html.Node {
	_ = "STUB: not implemented"
	return nil
}

func DomHTML(n *html.Node) []byte { _ = "STUB: not implemented"; return nil }

func DomTexhtml(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func domTexhtml0(n *html.Node, buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

func escapeMathSymbol(s string) string {
	_ = "STUB: not implemented"
	// 转义 Tex 公式中的符号，比如 _ ^ { }
	return ""
}

func DomText(n *html.Node) string { _ = "STUB: not implemented"; return "" }

func domText0(n *html.Node, buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

// 可能是自定义标签

// Protyle 中的搜索高亮标记需要保留 https://github.com/siyuan-note/siyuan/issues/9821

func IsTempMarkSpan(n *html.Node) bool { _ = "STUB: not implemented"; return false }

func SetDomAttrValue(n *html.Node, attrName, attrVal string) { _ = "STUB: not implemented"; return }

func RemoveDomAttr(n *html.Node, attrName string) { _ = "STUB: not implemented"; return }

func RemoveDomAttrs(n *html.Node) { _ = "STUB: not implemented"; return }

func DomAttrValue(n *html.Node, attrName string) string { _ = "STUB: not implemented"; return "" }

func ExistDomAttr(n *html.Node, attrName string) bool { _ = "STUB: not implemented"; return false }

func DomCustomAttrs(n *html.Node) (ret map[string]string) { _ = "STUB: not implemented"; return nil }

func DomAttrValuesWithPrefix(n *html.Node, prefix string) (ret map[string]string) {
	_ = "STUB: not implemented"
	return nil
}
