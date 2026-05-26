// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

//go:build !javascript
// +build !javascript

package render

import (
	"github.com/88250/lute/ast"
)

func (r *HtmlRenderer) renderCodeBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 缩进代码块处理

// renderCodeBlockCode 进行代码块 HTML 渲染，实现语法高亮。
func (r *HtmlRenderer) renderCodeBlockCode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// Go 代码块自动格式化 https://github.com/b3log/lute/issues/37

func highlightChroma(codeNode *ast.Node, tokens []byte, language string, r *HtmlRenderer) (rendered bool) {
	_ = "STUB: not implemented"
	return false
}

func isGo(language string) bool { _ = "STUB: not implemented"; return false }

// github.com/src-d/enry/v2 不怎么准确
//
//var candidateLangs = []string{
//	"bash", "csharp", "cpp", "css", "go", "html", "xml", "java", "js", "json", "kotlin", "less", "lua", "makefile", "markdown",
//	"nginx", "objc", "php", "properties", "python", "ruby", "rust", "scss", "sql", "shell", "toml", "ts", "yaml", "swift",
//	"dart", "gradle", "julia", "matlab",
//}
//
//func detectLanguage(code []byte) (language string) {
//	language, _ = enry.GetLanguageByClassifier(code, candidateLangs)
//	return
//}

func detectLanguage(code []byte) string { _ = "STUB: not implemented"; return "" }
