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

func (context *Context) parseTable(paragraph *ast.Node) (retParagraph, retTable *ast.Node) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 没有 | 的行依旧归入段落中

// 合并单元格
/* width: 是为了兼容遗留数据 */

// 合并单元格
/* width: 是为了兼容遗留数据 */

func (context *Context) parseTable0(tokens []byte) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// 换行+冒号会被识别为表格 https://github.com/88250/lute/issues/198

// 如果只有两行并且对齐方式是默认对齐且没有 | 时（foo\n---）就和 Setext 标题规则冲突了
// 但在块级解析时显然已经尝试进行解析 Setext 标题，还能走到这里说明 Setetxt 标题解析失败，
// 所以这里也不能当作表进行解析了，返回普通段落

func (context *Context) newTableHead(headRows []*ast.Node) *ast.Node {
	_ = "STUB: not implemented"
	return nil
}

func inInline(tokens []byte, i int, mathOrCodeMarker byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (context *Context) parseTableRow(line []byte, aligns []int, isHead bool) (ret *ast.Node) {
	_ = "STUB: not implemented"
	return nil
}

// 分隔符行定义了表的列数，如果表头列数还大于这个列数，则说明不满足表格式

// 可能需要补全剩余的列

func (context *Context) findTableDelimRow(lines [][]byte) (index int) {
	_ = "STUB: not implemented"
	return 0
}

func (context *Context) parseTableDelimRow(line []byte) (aligns []int) {
	_ = "STUB: not implemented"
	return nil
}

func (context *Context) tableDelimAlign(col []byte) int { _ = "STUB: not implemented"; return 0 }
