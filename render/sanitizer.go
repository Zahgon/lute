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
	"bytes"

	"github.com/88250/lute/html"
)

// 没有实现可扩展的策略，仅过滤不安全的标签和属性。
// 鸣谢 https://github.com/microcosm-cc/bluemonday

var setOfElementsToSkipContent = map[string]interface{}{
	"frame":    nil,
	"frameset": nil,
	//"iframe":   nil,
	"noembed":  nil,
	"noframes": nil,
	"noscript": nil,
	"nostyle":  nil,
	"object":   nil,
	"script":   nil,
	"style":    nil,
	"title":    nil,
}

func Sanitize(str string) string { _ = "STUB: not implemented"; return "" }

func sanitize(tokens []byte) []byte { _ = "STUB: not implemented"; return nil }

// do not escape multiple query parameters

// do not escape multiple query parameters

// not encouraged, but if a policy allows JavaScript we
// should not HTML escape it as that would break the output

// not encouraged, but if a policy allows CSS styles we
// should not HTML escape it as that would break the output

// HTML escape the text

func linkable(elementName string) bool { _ = "STUB: not implemented"; return false }

func writeLinkableBuf(buff *bytes.Buffer, token *html.Token) {
	_ = "STUB: not implemented"
	// do not escape multiple query parameters
	return
}

// re-apply

func sanitizeAttrs(attrs []*html.Attribute) (ret []*html.Attribute) {
	_ = "STUB: not implemented"
	return nil
}

func removeSpace(s string) string { _ = "STUB: not implemented"; return "" }

func allowAttr(attrName string) bool { _ = "STUB: not implemented"; return false }

// HTML 事件属性。https://www.w3schools.com/tags/ref_eventattributes.asp
var eventAttrs = map[string]interface{}{
	// Window
	"onafterprint":   nil,
	"onbeforeprint":  nil,
	"onbeforeunload": nil,
	"onerror":        nil,
	"onhashchange":   nil,
	"onload":         nil,
	"onmessage":      nil,
	"onoffline":      nil,
	"ononline":       nil,
	"onpagehide":     nil,
	"onpageshow":     nil,
	"onpopstate":     nil,
	"onresize":       nil,
	"onstorage":      nil,
	"onunload":       nil,

	// Form
	"onblur":        nil,
	"onchange":      nil,
	"oncontextmenu": nil,
	"onfocus":       nil,
	"oninput":       nil,
	"oninvalid":     nil,
	"onreset":       nil,
	"onsearch":      nil,
	"onselect":      nil,
	"onsubmit":      nil,

	// Keyboard
	"onkeydown":  nil,
	"onkeypress": nil,
	"onkeyup":    nil,

	// Mouse
	"onclick":      nil,
	"ondblclick":   nil,
	"onmousedown":  nil,
	"onmousemove":  nil,
	"onmouseout":   nil,
	"onmouseover":  nil,
	"onmouseleave": nil,
	"onmouseenter": nil,
	"onmouseup":    nil,
	"onmousewheel": nil,
	"onwheel":      nil,

	// Drag
	"ondrag":      nil,
	"ondragend":   nil,
	"ondragenter": nil,
	"ondragleave": nil,
	"ondragover":  nil,
	"ondragstart": nil,
	"ondrop":      nil,
	"onscroll":    nil,

	// Clipboard
	"oncopy":  nil,
	"oncut":   nil,
	"onpaste": nil,

	// Media
	"onabort":          nil,
	"oncanplay":        nil,
	"oncanplaythrough": nil,
	"oncuechange":      nil,
	"ondurationchange": nil,
	"onemptied":        nil,
	"onended":          nil,
	"onloadeddata":     nil,
	"onloadedmetadata": nil,
	"onloadstart":      nil,
	"onpause":          nil,
	"onplay":           nil,
	"onplaying":        nil,
	"onprogress":       nil,
	"onratechange":     nil,
	"onseeked":         nil,
	"onseeking":        nil,
	"onstalled":        nil,
	"onsuspend":        nil,
	"ontimeupdate":     nil,
	"onvolumechange":   nil,
	"onwaiting":        nil,

	// Misc
	"ontoggle": nil,

	// SVG
	"onbegin":  nil,
	"onend":    nil,
	"onrepeat": nil,

	// meta
	"http-equiv": nil,

	// input
	"formaction": nil,
}
