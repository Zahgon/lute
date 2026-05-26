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

func (t *Tree) parseCodeSpan(block *ast.Node, ctx *InlineContext) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// 如果首尾是空格并且整行不是空行时剔除首尾的一个空格

// 表格中的代码中带有管道符的处理 https://github.com/88250/lute/issues/63

// Improve `inline code` markdown editing https://github.com/siyuan-note/siyuan/issues/9978

// HTML 转换 Markdown 时需要转义 HTML 实体

// Code span 内容在 Protyle 下会再次调用 Inline 进行行级解析。
// 这里拷贝一份 ParseOption，只对这次二次解析关闭 autoLink，
// 避免反引号内的内容被自动转换，同时不影响外层普通文本的 autoLink。

func (t *Tree) matchCodeSpanEnd(tokens []byte, num int) (pos int) {
	_ = "STUB: not implemented"
	return 0
}
