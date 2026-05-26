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

var dollar = util.StrToBytes("$")

func (t *Tree) parseInlineMath(ctx *InlineContext) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// Protyle 不允许从行级派生块级

// 块节点

// $ 后面不能紧跟数字

// 中间包含 span 节点的话打断公式，以 span 优先

func (t *Tree) matchInlineMathEnd(tokens []byte) (pos int) { _ = "STUB: not implemented"; return 0 }
