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

// FootnotesStart 判断脚注定义（[^label]）是否开始。
func FootnotesStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

func FootnotesContinue(footnotesDef *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

func (t *Tree) FindFootnotesDef(label []byte) (pos int, def *ast.Node) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (t *Tree) ExistFootnotesDef() (ret bool) { _ = "STUB: not implemented"; return false }
