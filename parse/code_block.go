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
	"github.com/88250/lute/util"
)

// FenceCodeBlockStart 判断围栏代码块（```）是否开始。
func FenceCodeBlockStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// IndentCodeBlockStart 判断缩进代码块（    code）是否开始。
func IndentCodeBlockStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

func CodeBlockContinue(codeBlock *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

// 跳过围栏标记符之前可能存在的空格

// 缩进代码块

func (context *Context) codeBlockFinalize(codeBlock *ast.Node) { _ = "STUB: not implemented"; return }

// 缩进代码块

var codeBlockBacktick = util.StrToBytes("`")

func (t *Tree) parseFencedCode() (ok bool, fenceChar byte, fenceLen int, fenceOffset int, openFence, codeBlockInfo []byte) {
	_ = "STUB: not implemented"
	return false, 0, 0, 0, nil, nil
}

// info 部分不能包含 `

func (context *Context) isFencedCodeClose(tokens []byte, openMarker byte, num int) (ok bool, closeFence []byte) {
	_ = "STUB: not implemented"
	return false, nil
}
