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

func ParagraphContinue(p *ast.Node, context *Context) int { _ = "STUB: not implemented"; return 0 }

func paragraphFinalize(p *ast.Node, context *Context) (insertTable bool) {
	_ = "STUB: not implemented"
	return false
}

// 解析链接引用定义

// 列表项下没有子节点，应该挂一个空段落上去，并将当前段落转换为空段落的 IAL 节点

// 尝试解析任务列表项

// 如果是任务列表项则添加任务列表标记符节点

// 暂存于 p 的 IAL 上，最终化列表时会被置空

// 剔除开头的 [ ]、[x] 或者 [X]

// Protyle `Optimize typography` exception in case of task list and heading https://github.com/siyuan-note/siyuan/issues/9035

// Incomplete data when pasting task list nested list https://github.com/siyuan-note/siyuan/issues/9239

// 设置末梢及其状态

// 将该段落节点转成表节点

// 将该段落节点转换成目录节点
