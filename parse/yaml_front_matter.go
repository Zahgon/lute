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
	"github.com/88250/lute/editor"
	"github.com/88250/lute/util"
)

// 判断 YAML Front Matter（---）是否开始。
func YamlFrontMatterStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

func YamlFrontMatterContinue(node *ast.Node, context *Context) int {
	_ = "STUB: not implemented"
	return 0
}

var YamlFrontMatterMarker = util.StrToBytes("---")
var YamlFrontMatterMarkerNewline = util.StrToBytes("---\n")
var YamlFrontMatterMarkerCaret = util.StrToBytes("---" + editor.Caret)
var YamlFrontMatterMarkerCaretNewline = util.StrToBytes("---" + editor.Caret + "\n")

func (context *Context) yamlFrontMatterFinalize(node *ast.Node) { _ = "STUB: not implemented"; return }

// 剔除开头的 ---\n

// 剔除结尾的 ---‸

// 把 Vditor 插入符移动到内容末尾

// 剔除结尾的 ---

func (t *Tree) parseYamlFrontMatter() bool { _ = "STUB: not implemented"; return false }

func isYamlFrontMatterClose(context *Context) bool { _ = "STUB: not implemented"; return false }

// 判断 IAL 打断
