// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package lute

import (
	"github.com/88250/lute/html"
	"github.com/88250/lute/parse"
)

// SpinVditorIRDOM 自旋 Vditor Instant-Rendering DOM，用于即时渲染模式下的编辑。
func (lute *Lute) SpinVditorIRDOM(ivHTML string) (ovHTML string) {
	_ = "STUB: not implemented"
	// 替换插入符
	return ""
}

// 替换插入符

// HTML2VditorIRDOM 将 HTML 转换为 Vditor Instant-Rendering DOM，用于即时渲染模式下粘贴。
func (lute *Lute) HTML2VditorIRDOM(sHTML string) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

// VditorIRDOM2HTML 将 Vditor Instant-Rendering DOM 转换为 HTML，用于 Vditor.getHTML() 接口。
func (lute *Lute) VditorIRDOM2HTML(vhtml string) (sHTML string) {
	_ = "STUB: not implemented"
	return ""
}

// Md2VditorIRDOM 将 markdown 转换为 Vditor Instant-Rendering DOM，用于从源码模式切换至即时渲染模式。
func (lute *Lute) Md2VditorIRDOM(markdown string) (vHTML string) {
	_ = "STUB: not implemented"
	return ""
}

// VditorIRDOM2Md 将 Vditor Instant-Rendering DOM 转换为 markdown，用于从即时渲染模式切换至源码模式。
func (lute *Lute) VditorIRDOM2Md(htmlStr string) (markdown string) {
	_ = "STUB: not implemented"
	return ""
}

func (lute *Lute) vditorIRDOM2Md(htmlStr string) (markdown string) {
	_ = "STUB: not implemented"
	// 删掉插入符
	return ""
}

// 替换结尾空白，否则 HTML 解析会产生冗余节点导致生成空的代码块

// 将字符串解析为 DOM 树

// 调整 DOM 结构

// 将 HTML 树转换为 Markdown AST

// 调整树结构

// 合并代码节点 https://github.com/Vanessa219/vditor/issues/167

// 浏览器生成的子列表是 ul.ul 形式，需要将其调整为 ul.li.ul

// 将 AST 进行 Markdown 格式化渲染

// genASTByVditorIRDOM 根据指定的 Vditor IR DOM 节点 n 进行深度优先遍历并逐步生成 Markdown 语法树 tree。
func (lute *Lute) genASTByVditorIRDOM(n *html.Node, tree *parse.Tree) {
	_ = "STUB: not implemented"
	return
}

// 1：浮动工具栏，2：preview 代码块、数学公式块或者不解析的节点

// 处理在结尾 ``` 或者 $$ 后换行的情况
// TODO: 插入符现在已经不可能出现在该位置，确认后移除该段代码

// 处理 FireFox 某些情况下产生的分段

// 不允许在 bq 第一个节点前换行

// 子有序列表第一项必须从 1 开始

// 删掉表格中结尾的 br

// 仅允许 input 出现在任务列表中

// 在任务列表前退格

// ul.li.p.input

// 表格开头输入会导致解析问题，所以插入一个空段落进行分隔

// DOM 后缺少 info span 节点

// 把 ` 后面的字符调整到 info 节点

// kbd 标签由 code 标签构成节点
