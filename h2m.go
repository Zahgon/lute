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

// HTML2Markdown 将 HTML 转换为 Markdown。
func (lute *Lute) HTML2Markdown(htmlStr string) (markdown string, err error) {
	_ = "STUB: not implemented"
	//fmt.Println(htmlStr)
	// 将字符串解析为 DOM 树
	return "", nil
}

// 将 AST 进行 Markdown 格式化渲染

// HTML2Tree 将 HTML 转换为 AST。
func (lute *Lute) HTML2Tree(dom string) (ret *parse.Tree) { _ = "STUB: not implemented"; return nil }

func (lute *Lute) HTMLNode2Tree(n *html.Node) (ret *parse.Tree) {
	_ = "STUB: not implemented"
	return nil
}

// 调整 DOM 结构

// 将 HTML 树转换为 Markdown AST

// 调整树结构

// ul.ul => ul.li.ul

// genASTByDOM 根据指定的 DOM 节点 n 进行深度优先遍历并逐步生成 Markdown 语法树 tree。
func (lute *Lute) genASTByDOM(n *html.Node, tree *parse.Tree) { _ = "STUB: not implemented"; return }

/* 简书代码块 https://github.com/siyuan-note/siyuan/issues/4361 */

// 忽略 Wikipedia [编辑] Do not clip the `Edit` element next to Wikipedia headings https://github.com/siyuan-note/siyuan/issues/11600

// 忽略 Wikipedia 引用中的注释 https://github.com/siyuan-note/siyuan/issues/11640

// 自定义标签

// 链滴剪藏图片时多了长宽显示 https://github.com/siyuan-note/siyuan/issues/10987

// 将 \n空格空格* 转换为\n

// 将 &nbsp; 转换为空格

// 将 \n 转换为空格 https://github.com/siyuan-note/siyuan/issues/6052

// 将其转换为公式块 https://github.com/siyuan-note/siyuan/issues/14360

// 按原文解析，不处理转义

// a 标签锚文本中的标记符不进行转义 https://github.com/siyuan-note/siyuan/issues/14733

// 去掉文本中开头的列表项标记符 https://github.com/siyuan-note/siyuan/issues/14329

// h 下存在 div/p/section 则忽略分块

// 解析 GitHub 语法高亮代码块

// The browser extension supports CSDN formula https://github.com/siyuan-note/siyuan/issues/5624

// The browser extension supports Wikipedia formula clipping https://github.com/siyuan-note/siyuan/issues/11583

// 这里用 for 是为了简化实现

// https://github.com/siyuan-note/siyuan/issues/15457

// span 可能是 TextMark 元素，也可能是公式，其他情况则忽略

// 子有序列表必须从 1 开始

// 删除第一个 code 之前的标签

// 处理 pre.i,ol 的情况，i 标签为“复制代码”或“隐藏代码”，ol 标签为“行号”
// https://github.com/siyuan-note/siyuan/issues/15314

// pre.em,em,code 的情况，这两个 em 是“复制代码”和“隐藏代码” https://github.com/siyuan-note/siyuan/issues/13026

// pre 下只有一个 div，且 div 下只有一个 code，那么将 pre.div 替换为 pre.code https://github.com/siyuan-note/siyuan/issues/11131

// pre 下全是 div，每个 div 为一行代码 https://github.com/siyuan-note/siyuan/issues/14195
// 将其转换为 pre.code， code, ... code，每个 div 为一行代码，然后交由后续处理

// 改进两种 pre.ol.li 的代码块解析 https://github.com/siyuan-note/siyuan/issues/11296
// 第一种：将 pre.ol.li.p.span, span, ... span 转换为 pre.ol.li.p.code, code, ... code，然后交由第二种处理

// 第二种：将 pre.ol.li.p.code, code, ... code 转换为 pre.code, code, ... code，然后交由后续处理

// pre.code code 每个 code 为一行的结构，需要在 code 中间插入换行

// pre.code.span 每个 span 为一行的结构，需要在 span 中间插入换行

// CSDN 代码块：pre.code.ol.li

// CSDN 代码块：pre.code,ul
// 去掉最后一个代码行号子块 https://github.com/siyuan-note/siyuan/issues/5564

// pre.code code 每个 code 为一行的结构，需要在 code 中间插入换行

// 避免下面 util.DomText 把 p 转换为两个换行

// 如果表格中只有一行一列，那么丢弃表格直接使用代码块
// Improve HTML parsing code blocks https://github.com/siyuan-note/siyuan/issues/11068

// 表格中不支持添加块级元素，所以这里只能将其转换为多个行级代码元素

// 在两个相邻的加粗或者斜体之间插入零宽空格，避免标记符重复

// https://github.com/siyuan-note/siyuan/issues/11682

// 如果全部都是 span 子节点，那么直接使用 span 的内容 https://github.com/siyuan-note/siyuan/issues/11281

// 丢弃标题中文本为空的链接，这样的链接是没有锚文本的锚点
// https://github.com/Vanessa219/vditor/issues/359
// https://github.com/siyuan-note/siyuan/issues/11445

// 剪藏时过滤空的超链接 https://github.com/siyuan-note/siyuan/issues/5686

// Wikipedia 链接嵌套图片的情况只保留图片

// 直接使用 alt 值（即 emoji 字符）https://github.com/siyuan-note/siyuan/issues/13342

// 处理知乎动图

// 处理可能存在的预加载情况

// 处理使用 data-original 属性的情况 https://github.com/siyuan-note/siyuan/issues/11826

// 处理使用 srcset 属性的情况

// Wikipedia 使用图片原图 https://github.com/siyuan-note/siyuan/issues/11640

// 如果上一个节点不是块级元素，那么需要先添加一个换行节点 Improve HTML table clipping https://github.com/siyuan-note/siyuan/issues/15307

// 找到最多的 td 数

// 补全 thead 中 tr 的 th

// 补全 thread 节点

// 转换为行级备注 https://github.com/siyuan-note/siyuan/issues/13998

// 转换为行级备注 https://github.com/siyuan-note/siyuan/issues/13998

// Improve inline elements pasting https://github.com/siyuan-note/siyuan/issues/11740

// 简化为只处理第一个类型

// The browser extension supports Zhihu formula https://github.com/siyuan-note/siyuan/issues/5599

// The browser extension supports CSDN formula https://github.com/siyuan-note/siyuan/issues/5624

// 根据最后 4 个换行符分隔公式内容

// 单独处理 <figcaption> 中的 <span leaf> 标签的情况 https://github.com/siyuan-note/siyuan/issues/14507

// 图片标题不包含非文本元素，包含的话走下面的逻辑，单独作为一个段落

func appendInlineMath(tree *parse.Tree, tex string) { _ = "STUB: not implemented"; return }

func appendMathBlock(tree *parse.Tree, tex string) { _ = "STUB: not implemented"; return }

func appendSpace(n *html.Node, tree *parse.Tree, lute *Lute) { _ = "STUB: not implemented"; return }
