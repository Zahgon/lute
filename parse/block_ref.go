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

func (t *Tree) parseBlockRef(ctx *InlineContext) *ast.Node { _ = "STUB: not implemented"; return nil }

// 这里使用 for 是为了简化逻辑，不是为了循环

// 跟空格的话后续尝试锚文本解析

func (context *Context) parseBlockRefID(tokens []byte) (passed, remains, id []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
