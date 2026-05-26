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

// SuperBlockStart 判断超级块（{{{ blocks }}}）是否开始。
func SuperBlockStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// 整行过

func SuperBlockContinue(superBlock *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

// 闭合

func (context *Context) superBlockFinalize(superBlock *ast.Node) {
	_ = "STUB: not implemented"
	// 最终化所有子块
	return
}

func (t *Tree) parseSuperBlock() (ok bool, layout []byte) {
	_ = "STUB: not implemented"
	return false, nil
}

func (context *Context) isSuperBlockClose(tokens []byte) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
