// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package html

import (
	"unicode/utf8"
)

func UnescapeBytes(tokens []byte) (ret []byte) { _ = "STUB: not implemented"; return nil }

func HtmlUnescapeString(s string) string {
	_ = "STUB: not implemented"
	// 鸣谢 https://gitlab.com/golang-commonmark
	return ""
}

func parseEntity(s string) (string, int) { _ = "STUB: not implemented"; return "", 0 }

// initial state

// &#

// &q

// &#x

// &#0

// &#x0

const BadEntity = string(utf8.RuneError)

func isValidEntityCode(c int64) bool { _ = "STUB: not implemented"; return false }

// never used

// control codes
