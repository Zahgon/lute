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

// delimiter 描述了强调、链接和图片解析过程中用到的分隔符（[, ![, *, _, ~）相关信息。
type delimiter struct {
	node           *ast.Node  // 分隔符对应的文本节点
	typ            byte       // 分隔符字节 [*_~
	num            int        // 分隔符字节数
	originalNum    int        // 原始分隔符字节数
	canOpen        bool       // 是否是开始分隔符
	canClose       bool       // 是否是结束分隔符
	previous, next *delimiter // 双向链表前后节点

	active            bool
	image             bool
	bracketAfter      bool
	index             int
	previousDelimiter *delimiter
}

// 嵌套强调和链接的解析算法的中文解读可参考这里 https://ld246.com/article/1566893557720

// handleDelim 将分隔符 *_~ 入栈。
func (t *Tree) handleDelim(block *ast.Node, ctx *InlineContext) { _ = "STUB: not implemented"; return }

// 将这个分隔符入栈

// processEmphasis 处理强调、加粗以及删除线。
func (t *Tree) processEmphasis(stackBottom *delimiter, ctx *InlineContext) {
	_ = "STUB: not implemented"
	return
}

// find first closer above stack_bottom:

// move forward, looking for closers, and handling each

// found emphasis closer. now look back for first matching opener:

// calculate actual number of delimiters used from closer

// remove used delimiters from stack elts and inlines

// 优先下标

// 插入起始标记符
// 插入结束标记符

// remove elts between opener and closer in delimiters stack

// if opener has 0 delims, remove it and the inline

// Set lower bound for future searches for openers:

// We can remove a closer that can't be an opener,
// once we've seen there's no matching opener:

// 移除所有分隔符

func (t *Tree) scanDelims(ctx *InlineContext) *delimiter { _ = "STUB: not implemented"; return nil }

// 跳过插入符位置向前看

// Markdown 中 ** 加粗失效问题 https://ld246.com/article/1597581380183

/* ==Mark== 标记使用两个等号 */

/* #Tag# 标记使用一个井号 */

/* ^Sup^ 标记使用一个 ^ */

// 单独处理 ~~~foo~~~ 的情况，即下标嵌套删除线

// ~Sub~ 标记使用一个 ~

func (t *Tree) removeDelimiter(delim *delimiter, ctx *InlineContext) (ret *delimiter) {
	_ = "STUB: not implemented"
	return nil
}

// 栈顶
