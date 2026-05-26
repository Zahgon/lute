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

func (context *Context) parseLinkRefDef(tokens []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (context *Context) parseLinkTitle(tokens []byte) (validTitle bool, passed, remains, title []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil
}

func (context *Context) parseBlockRefText(tokens []byte) (validTitle bool, passed, remains, title []byte, subtype string) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil, ""
}

func (context *Context) parseLinkTitleMatch(opener, closer byte, tokens []byte) (validTitle bool, passed, remains, title []byte) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil
}

func (context *Context) parseLinkDest(tokens []byte) (ret, remains, destination []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// <autolink>

// [label](/url)

func (context *Context) parseLinkDest2(tokens []byte) (ret, remains, destination []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (context *Context) parseLinkDest1(tokens []byte) (ret, remains, destination []byte) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (context *Context) parseLinkLabel(tokens []byte) (n int, remains, label []byte) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}
