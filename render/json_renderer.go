// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package render

import (
	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

type JSONRenderer struct {
	*BaseRenderer
}

func NewJSONRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) Renderer {
	_ = "STUB: not implemented"
	// 渲染器剔除语法树块级 IAL 节点
	return *new(Renderer)
}

func (r *JSONRenderer) renderNode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 去掉结尾的 }

func ial2Map(ial [][]string) (ret map[string]string) { _ = "STUB: not implemented"; return nil }
