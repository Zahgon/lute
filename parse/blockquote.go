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

// BlockquoteStart 判断引述（>）是否开始。
func BlockquoteStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// > 后面的空格是可选的

func BlockquoteContinue(blockquote *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

func (context *Context) blockquoteFinalize(blockquote *ast.Node) { _ = "STUB: not implemented"; return }
