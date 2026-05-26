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

// Space 会把 tokens 中的中西文之间加上空格。
func (r *BaseRenderer) Space(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

func Space0(text string) (ret string) { _ = "STUB: not implemented"; return "" }

// ing 前不需要空格，如 打码ing https://github.com/88250/lute/issues/9

func addSpaceAtBoundary(prefix string, nextChar rune) string { _ = "STUB: not implemented"; return "" }

// Emoji 1-9
// 在这里处理并不是太合适，应该在 emoji.go 中直接将 Unicode Emoji 解析为节点

func allowSpace(currentChar, nextChar rune) bool { _ = "STUB: not implemented"; return false }

func isCJK(r rune) bool { _ = "STUB: not implemented"; return false }

func isFullWidth(r rune) bool { _ = "STUB: not implemented"; return false }
