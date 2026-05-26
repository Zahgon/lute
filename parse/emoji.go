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

// emoji 将 node 下文本节点和链接文本节点中的 Emoji 别名替换为原生 Unicode 字符。
func (t *Tree) emoji(node *ast.Node) { _ = "STUB: not implemented"; return }

// 递归处理子节点

var EmojiSitePlaceholder = util.StrToBytes("${emojiSite}")
var emojiDot = util.StrToBytes(".")

func (t *Tree) emoji0(node *ast.Node) { _ = "STUB: not implemented"; return }

// 先清空，后面逐个添加或者添加 Tokens 或者 Emoji 兄弟节点

// 有的 Emoji 是图片链接，需要单独处理

// 自定义 Emoji 路径用 . 判断，包含 . 的认为是图片路径

// 在 Emoji 节点后插入一个内容为空的文本节点，留作下次迭代

// 丢弃空的文本节点

func (t *Tree) EmojiImgTokens(alias, src string) []byte { _ = "STUB: not implemented"; return nil }
