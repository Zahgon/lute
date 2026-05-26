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
)

func EChartsMindmapStr(listContent string) string { _ = "STUB: not implemented"; return "" }

func EChartsMindmap(listContent []byte) []byte { _ = "STUB: not implemented"; return nil }

// echartsMindmap 用于将列表 Markdown 原文转为 ECharts 树图结构，提供给前端渲染脑图。
func echartsMindmap(listContent []byte) []byte { _ = "STUB: not implemented"; return nil }

// 第一个节点如果不是列表的话直接返回

// 移除非列表节点

// 如果根节点下的第一个列表包含多个列表项，则自动生成一个根节点，这些列表项都挂在这个根节点上

// text 返回列表项第一个子节点的文本内容。
func text(listItemFirstChild *ast.Node) (ret string) { _ = "STUB: not implemented"; return "" }

// 遍历到下一个列表或者列表项时退出

func needRoot(root *ast.Node) bool {
	_ = "STUB: not implemented"

	// 检查根节点下是否包含多个列表
	return false
}

// 包含多个列表则需要构建一个 Root 节点

// 没有列表也需要构建一个 Root 节点

// 如果只有一个列表，则检查该列表下是否包含多个列表项

// 包含多个列表项则需要构建一个 Root 节点
