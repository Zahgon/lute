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

// ListStart 判断列表、列表项（* - + 1.）或者任务列表项是否开始。
func ListStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// 修正有序列表项序号

func ListItemContinue(listItem *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

// 列表项后面是空的

func (context *Context) listFinalize(list *ast.Node) { _ = "STUB: not implemented"; return }

// 检查子列表项之间是否包含空行，包含的话说明该列表是非紧凑的，即松散的

var items1 = util.StrToBytes("1")

// parseListMarker 用于解析泛列表（列表、列表项或者任务列表）标记符。
func (t *Tree) parseListMarker(container *ast.Node) (data *ast.ListData, ial [][]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 默认无序列表
// 默认紧凑模式
// 设置前置相对缩进
// 假设有序列表起始为 -1，后面会进行计算赋值

// 有序列表

// 列表项标记符后必须是空白字符

// 如果要打断段落，则列表项内容部分不能为空

// 到这里说明满足列表规则，开始解析并计算内部缩进空格数
// 把起始下标移动到标记符起始位置
// 把结束下标移动到标记符结束位置

// 判断是否是任务列表项

// 至少需要 [ ] 或者 [x] 3 个字符

func (t *Tree) parseOrderedListMarker(tokens []byte) (marker []byte, delimiter byte) {
	_ = "STUB: not implemented"
	return nil, 0
}

// endsWithBlankLine 判断块节点 block 是否是空行结束。如果 block 是列表或者列表项则迭代下降进入判断。
func endsWithBlankLine(block *ast.Node) bool { _ = "STUB: not implemented"; return false }
