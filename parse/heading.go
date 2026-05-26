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

// ATXHeadingStart 判断 ATX 标题（#）是否开始。
func ATXHeadingStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// SetextHeadingStart 判断 Setext 标题（- =）是否开始。
func SetextHeadingStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// 尝试解析表，因为可能出现如下情况：
//
//   0
//   -:
//   -
//
// 前两行可以解析出一个只有一个单元格的表。
// Empty list following GFM Table makes table broken https://github.com/b3log/lute/issues/9

// 将该段落节点转成表节点

// 解析链接引用定义

func (t *Tree) parseATXHeading() (ok bool, markers, content []byte, level int) {
	_ = "STUB: not implemented"
	return false, nil, nil, 0
}

func (t *Tree) parseSetextHeading() (level int) { _ = "STUB: not implemented"; return 0 }
