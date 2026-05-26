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
)

// NestedInlines2FlattedSpansHybrid 将嵌套的行级节点转换为平铺的文本标记节点。
// 该函数不会移除转义节点。
func NestedInlines2FlattedSpansHybrid(tree *Tree, isExportMd bool) {
	_ = "STUB: not implemented"
	return
}

// 超链接嵌套图片情况下，图片子节点移到超链接节点前面

// case ast.NodeBackslash: Spin 过程中存在转义节点，转义节点本身就是嵌套的，所以需要在这里排除处理

// 粘贴 Markdown 时行级元素中的双引号不再转换为实体 https://github.com/siyuan-note/siyuan/issues/14503

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227

// 合并相邻的链接文本节点

// https://github.com/siyuan-note/siyuan/issues/14788#issuecomment-2882248074

// NestedInlines2FlattedSpans 将嵌套的行级节点转换为平铺的文本标记节点。
// 该函数会移除转义节点。
func NestedInlines2FlattedSpans(tree *Tree, isExportMd bool) { _ = "STUB: not implemented"; return }

// 超链接嵌套图片情况下，图片子节点移到超链接节点前面

// 不再需要反斜杠转义节点

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227

// 合并相邻的链接文本节点

func processNestedNode(n *ast.Node, tag string, tags *[]string, unlinks *[]*ast.Node, entering bool) {
	_ = "STUB: not implemented"
	return
}

func TextMarks2Inlines(tree *Tree) { _ = "STUB: not implemented"; return }
