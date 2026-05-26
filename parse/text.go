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

func (t *Tree) parseText(ctx *InlineContext) *ast.Node { _ = "STUB: not implemented"; return nil }

// 遇到潜在的标记符时需要跳出该文本节点，回到行级解析主循环

// isMarker 判断 token 是否是潜在的 Markdown 标记符。
func (t *Tree) isMarker(token byte) bool { _ = "STUB: not implemented"; return false }

var backslash = util.StrToBytes("\\")

func (t *Tree) parseBackslash(block *ast.Node, ctx *InlineContext) *ast.Node {
	_ = "STUB: not implemented"
	return nil
}

// 表格单元格内存在多行时末尾输入转义符 `\` 导致 `<br />` 暴露 https://github.com/siyuan-note/siyuan/issues/7725

// 处理 \‸x 情况，插入符后的字符才是待转义的

// 表格单元格内存在多行时末尾输入转义符 `\` 导致 `<br />` 暴露 https://github.com/siyuan-note/siyuan/issues/7725

// Protyle WYSIWYG 模式下插入符移到转义符节点前面

func (t *Tree) parseNewline(block *ast.Node, ctx *InlineContext) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// 检查前一个节点的结尾空格，如果大于等于两个则说明是硬换行

// Protyle 中的换行符都是软换行 Improve soft line break paste parsing https://github.com/siyuan-note/siyuan/issues/14481

func (t *Tree) MergeText() { _ = "STUB: not implemented"; return }

// mergeText 合并 node 中所有（包括子节点）连续的文本节点。
// 合并后顺便进行中文排版优化以及 GFM 自动邮件链接识别。
func (t *Tree) mergeText(node *ast.Node) { _ = "STUB: not implemented"; return }

// 逐个合并后续兄弟节点

// 递归处理子节点
