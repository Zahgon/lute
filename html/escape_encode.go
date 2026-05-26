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

var (
	amp  = []byte("&amp;")
	lt   = []byte("&lt;")
	gt   = []byte("&gt;")
	quot = []byte("&quot;")
)

func UnescapeAttrVal(v string) string { _ = "STUB: not implemented"; return "" }

func EscapeAttrVal(v string) (ret string) { _ = "STUB: not implemented"; return "" }

func UnescapeHTMLStr(h string) string { _ = "STUB: not implemented"; return "" }

func EscapeHTMLStr(h string) string { _ = "STUB: not implemented"; return "" }

func UnescapeHTML(h []byte) (ret []byte) { _ = "STUB: not implemented"; return nil }

func EscapeHTML(html []byte) (ret []byte) { _ = "STUB: not implemented"; return nil }

// 通过延迟初始化减少内存分配，下同

// EncodeDestination percent-encodes rawurl, avoiding double encoding.
// It doesn't touch:
// - alphanumeric characters ([0-9a-zA-Z]);
// - percent-encoded characters (%[0-9a-fA-F]{2});
// - excluded characters ([;/?:@&=+$,-_.!~*'()#]).
// Invalid UTF-8 sequences are replaced with U+FFFD.
func EncodeDestination(rawurl []byte) (ret []byte) {
	_ = "STUB: not implemented"
	// 鸣谢 https://gitlab.com/golang-commonmark/mdurl
	return nil
}

// DecodeDestination decodes a percent-encoded URL.
// Invalid percent-encoded sequences are left as is.
// Invalid UTF-8 sequences are replaced with U+FFFD.
func DecodeDestination(rawurl []byte) []byte {
	_ = "STUB: not implemented"
	// 鸣谢 https://gitlab.com/golang-commonmark/mdurl
	return nil
}

func unhex(b byte) byte { _ = "STUB: not implemented"; return 0 }

func advance(s []byte, pos int) (byte, int) { _ = "STUB: not implemented"; return 0, 0 }
