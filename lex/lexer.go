// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package lex

// Lexer 描述了词法分析器结构。
type Lexer struct {
	input  []byte // 输入的文本字节数组
	length int    // 输入的文本字节数组的长度
	offset int    // 当前读取字节位置
	width  int    // 最新一个字符的长度（字节数）
}

// NewLexer 创建一个词法分析器。
func NewLexer(input []byte) (ret *Lexer) { _ = "STUB: not implemented"; return nil }

// 以 \n 结尾预处理

// NextLine 返回下一行。
func (l *Lexer) NextLine() (ret []byte) { _ = "STUB: not implemented"; return nil }

// \r\n
// 移除 \r，依靠下一个的 \n 切行
// 重新计算总长
// \rX
// 将 \r 替换为 \n

// \rEOF
// 将 \r 替换为 \n

// 将 \u0000 替换为 \uFFFD

// \uFFFD 的 UTF-8 编码为 \xEF\xBF\xBD 共三个字节

// 重新计算总长

// 说明占用多个字节
