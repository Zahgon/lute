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

const (
	ItemBacktick       = byte('`')
	ItemTilde          = byte('~')
	ItemBang           = byte('!')
	ItemCrosshatch     = byte('#')
	ItemAsterisk       = byte('*')
	ItemOpenParen      = byte('(')
	ItemCloseParen     = byte(')')
	ItemHyphen         = byte('-')
	ItemUnderscore     = byte('_')
	ItemPlus           = byte('+')
	ItemEqual          = byte('=')
	ItemTab            = byte('\t')
	ItemOpenBracket    = byte('[')
	ItemCloseBracket   = byte(']')
	ItemDoublequote    = byte('"')
	ItemSinglequote    = byte('\'')
	ItemLess           = byte('<')
	ItemGreater        = byte('>')
	ItemSpace          = byte(' ')
	ItemNewline        = byte('\n')
	ItemCarriageReturn = byte('\r')
	ItemBackslash      = byte('\\')
	ItemSlash          = byte('/')
	ItemDot            = byte('.')
	ItemColon          = byte(':')
	ItemQuestion       = byte('?')
	ItemAmpersand      = byte('&')
	ItemSemicolon      = byte(';')
	ItemPipe           = byte('|')
	ItemDollar         = byte('$')
	ItemCaret          = byte('^')
	ItemOpenBrace      = byte('{')
	ItemCloseBrace     = byte('}')
)

// IsWhitespace 判断 token 是否是空白。
func IsWhitespace(token byte) bool { _ = "STUB: not implemented"; return false }

// IsUnicodeWhitespace 判断 token 是否是 Unicode 空白。
func IsUnicodeWhitespace(r rune) bool { _ = "STUB: not implemented"; return false }

// IsDigit 判断 token 是否为数字 0-9。
func IsDigit(token byte) bool { _ = "STUB: not implemented"; return false }

// IsHexDigit 判断 token 是否是十六进制数字。
func IsHexDigit(token byte) bool { _ = "STUB: not implemented"; return false }

// TokenToUpper 将 token 转为大写。
func TokenToUpper(token byte) byte { _ = "STUB: not implemented"; return 0 }

// IsASCIIPunct 判断 token 是否是一个 ASCII 标点符号。
func IsASCIIPunct(token byte) bool { _ = "STUB: not implemented"; return false }

// IsASCIILetter 判断 token 是否是一个 ASCII 字母。
func IsASCIILetter(token byte) bool { _ = "STUB: not implemented"; return false }

// IsASCIILetterNum 判断 token 是否是一个 ASCII 字母或数字。
func IsASCIILetterNum(token byte) bool { _ = "STUB: not implemented"; return false }

// IsASCIILetterNums 判断 tokens 是否是 ASCII 字母或数字组成。
func IsASCIILetterNums(tokens []byte) bool { _ = "STUB: not implemented"; return false }

// IsASCIILetterNumHyphen 判断 token 是否是一个 ASCII 字母、数字或者横线 -。
func IsASCIILetterNumHyphen(token byte) bool { _ = "STUB: not implemented"; return false }

// IsControl 判断 token 是否是一个控制字符。
func IsControl(token byte) bool { _ = "STUB: not implemented"; return false }

// IsBlank 判断 Tokens 是否都为空格。
func IsBlank(tokens []byte) bool { _ = "STUB: not implemented"; return false }

func Split(tokens []byte, separator byte) (ret [][]byte) { _ = "STUB: not implemented"; return nil }

// SplitWithoutBackslashEscape 使用 separator 作为分隔符将 Tokens 切分为多个子串，被反斜杠 \ 转义的字符不会计入切分。
func SplitWithoutBackslashEscape(tokens []byte, separator byte) (ret [][]byte) {
	_ = "STUB: not implemented"
	return nil
}

//if ItemPipe == token && inInlineMath(tokens, i) {
//	line = append(line, token)
//	continue
//}

// ReplaceAll 会将 Tokens 中的所有 old 使用 new 替换。
func ReplaceAll(tokens []byte, old, new byte) []byte { _ = "STUB: not implemented"; return nil }

// ReplaceNewlineSpace 会将 Tokens 中的所有 "\n " 替换为 "\n"。
func ReplaceNewlineSpace(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

func TrimWhitespace(tokens []byte) (ret []byte) { _ = "STUB: not implemented"; return nil }

func Trim(tokens []byte) (leftWhitespaces, rightWhitespaces, remains []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func TrimRight(tokens []byte) (whitespaces, remains []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TrimLeft(tokens []byte) (whitespaces, remains []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Accept(tokens []byte, token byte) (pos int) { _ = "STUB: not implemented"; return 0 }

func AcceptTokenss(tokens []byte, someTokenss [][]byte) (pos int) {
	_ = "STUB: not implemented"
	return 0
}

func AcceptTokens(remains, someTokens []byte) (pos int) { _ = "STUB: not implemented"; return 0 }

func IsBlankLine(tokens []byte) bool { _ = "STUB: not implemented"; return false }

func SplitWhitespace(tokens []byte) (ret [][]byte) { _ = "STUB: not implemented"; return nil }

// IsBackslashEscapePunct 判断 Tokens 中 pos 所指的值是否是由反斜杠 \ 转义的 ASCII 标点符号。
func IsBackslashEscapePunct(tokens []byte, pos int) bool { _ = "STUB: not implemented"; return false }

func StatWhitespace(tokens []byte) (newlines, spaces, tabs int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func Spnl(tokens []byte) (ret bool, passed, remains []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func Peek(tokens []byte, pos int) byte { _ = "STUB: not implemented"; return 0 }

// BytesShowLength 获取字节数组展示为 UTF8 字符串时的长度。
func BytesShowLength(bytes []byte) int { _ = "STUB: not implemented"; return 0 }

// 按位与 11000000 为 10000000 则表示为 UTF8 字节首位

func RepeatBackslashBeforePipe(content string) string { _ = "STUB: not implemented"; return "" }

func EscapeCommonMarkers(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

func EscapeProtyleMarkers(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

func IsCommonInlineMarker(token byte) bool { _ = "STUB: not implemented"; return false }

func IsProtyleInlineMarker(token byte) bool { _ = "STUB: not implemented"; return false }

func IsMarker(token byte) bool { _ = "STUB: not implemented"; return false }
