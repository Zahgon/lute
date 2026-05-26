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
	"github.com/88250/lute/editor"
	"github.com/88250/lute/util"
)

// MathBlockStart 判断数学公式块（$$）是否开始。
func MathBlockStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

func MathBlockContinue(mathBlock *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

// 跳过 $ 之前可能存在的空格

var MathBlockMarker = util.StrToBytes("$$")
var MathBlockMarkerCaret = util.StrToBytes("$$" + editor.Caret)

func (context *Context) mathBlockFinalize(mathBlock *ast.Node) { _ = "STUB: not implemented"; return }

/*
	- foo

	    $$
	bar
	$$
*/

// 剔除开头的 $$

// 剔除结尾的 $$‸

// 把 Vditor 插入符移动到内容末尾

// 剔除结尾的 $$

// 行级元素转换为块级元素 https://ld246.com/article/1730804245164

func (t *Tree) parseMathBlock() (ok bool, mathBlockDollarOffset int) {
	_ = "STUB: not implemented"
	return false, 0
}

func (context *Context) isMathBlockClose(tokens []byte) bool {
	_ = "STUB: not implemented"
	return false
}

// 判断 IAL 打断
