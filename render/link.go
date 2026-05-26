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

func (r *BaseRenderer) EncodeLinkSpace(dest string) string { _ = "STUB: not implemented"; return "" }

// Improve export of Markdown hyperlink spaces and markers https://github.com/siyuan-note/siyuan/issues/9792

func (r *BaseRenderer) LinkPath(dest []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *BaseRenderer) PrefixPath(dest []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *BaseRenderer) RelativePath(dest []byte) []byte { _ = "STUB: not implemented"; return nil }

func (r *BaseRenderer) isRelativePath(dest []byte) bool { _ = "STUB: not implemented"; return false }

// 检查特定协议前缀
