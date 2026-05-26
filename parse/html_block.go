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
	"github.com/88250/lute/util"
)

// HtmlBlockStart 判断 HTML 块（<）是否开始。
func HtmlBlockStart(t *Tree, container *ast.Node) int { _ = "STUB: not implemented"; return 0 }

// Protyle WYSIWYG 模式下，只有 <div 开头的块级元素才能被解析为 HTML 块
// Only HTML code wrapped in `<div>` is supported to be parsed into HTML blocks https://github.com/siyuan-note/siyuan/issues/9758

func HtmlBlockContinue(html *ast.Node, context *Context) int { _ = "STUB: not implemented"; return 0 }

// 判断 IAL 打断

func (context *Context) htmlBlockFinalize(html *ast.Node) { _ = "STUB: not implemented"; return }

var (
	htmlBlockTags1      = [][]byte{util.StrToBytes("<script"), util.StrToBytes("<pre"), util.StrToBytes("<style"), util.StrToBytes("<textarea")}
	htmlBlockCloseTags1 = [][]byte{util.StrToBytes("</script>"), util.StrToBytes("</pre>"), util.StrToBytes("</style>"), util.StrToBytes("</textarea>")}
	htmlBlockTags6      = [][]byte{
		util.StrToBytes("<address"), util.StrToBytes("<article"), util.StrToBytes("<aside"), util.StrToBytes("<base"), util.StrToBytes("<basefont"), util.StrToBytes("<blockquote"), util.StrToBytes("<body"), util.StrToBytes("<caption"), util.StrToBytes("<center"), util.StrToBytes("<col"), util.StrToBytes("<colgroup"), util.StrToBytes("<dd"), util.StrToBytes("<details"), util.StrToBytes("<dialog"), util.StrToBytes("<dir"), util.StrToBytes("<div"), util.StrToBytes("<dl"), util.StrToBytes("<dt"), util.StrToBytes("<fieldset"), util.StrToBytes("<figcaption"), util.StrToBytes("<figure"), util.StrToBytes("<footer"), util.StrToBytes("<form"), util.StrToBytes("<frame"), util.StrToBytes("<frameset"), util.StrToBytes("<h1"), util.StrToBytes("<h2"), util.StrToBytes("<h3"), util.StrToBytes("<h4"), util.StrToBytes("<h5"), util.StrToBytes("<h6"), util.StrToBytes("<head"), util.StrToBytes("<header"), util.StrToBytes("<hr"), util.StrToBytes("<html"), util.StrToBytes("<iframe"), util.StrToBytes("<legend"), util.StrToBytes("<li"), util.StrToBytes("<link"), util.StrToBytes("<main"), util.StrToBytes("<menu"), util.StrToBytes("<menuitem"), util.StrToBytes("<nav"), util.StrToBytes("<noframes"), util.StrToBytes("<ol"), util.StrToBytes("<optgroup"), util.StrToBytes("<option"), util.StrToBytes("<p"), util.StrToBytes("<param"), util.StrToBytes("<section"), util.StrToBytes("<source"), util.StrToBytes("<summary"), util.StrToBytes("<table"), util.StrToBytes("<tbody"), util.StrToBytes("<td"), util.StrToBytes("<tfoot"), util.StrToBytes("<th"), util.StrToBytes("<thead"), util.StrToBytes("<title"), util.StrToBytes("<tr"), util.StrToBytes("<track"), util.StrToBytes("<ul"), util.StrToBytes("<video"), util.StrToBytes("<audio"),
		util.StrToBytes("</address"), util.StrToBytes("</article"), util.StrToBytes("</aside"), util.StrToBytes("</base"), util.StrToBytes("</basefont"), util.StrToBytes("</blockquote"), util.StrToBytes("</body"), util.StrToBytes("</caption"), util.StrToBytes("</center"), util.StrToBytes("</col"), util.StrToBytes("</colgroup"), util.StrToBytes("</dd"), util.StrToBytes("</details"), util.StrToBytes("</dialog"), util.StrToBytes("</dir"), util.StrToBytes("</div"), util.StrToBytes("</dl"), util.StrToBytes("</dt"), util.StrToBytes("</fieldset"), util.StrToBytes("</figcaption"), util.StrToBytes("</figure"), util.StrToBytes("</footer"), util.StrToBytes("</form"), util.StrToBytes("</frame"), util.StrToBytes("</frameset"), util.StrToBytes("</h1"), util.StrToBytes("</h2"), util.StrToBytes("</h3"), util.StrToBytes("</h4"), util.StrToBytes("</h5"), util.StrToBytes("</h6"), util.StrToBytes("</head"), util.StrToBytes("</header"), util.StrToBytes("</hr"), util.StrToBytes("</html"), util.StrToBytes("</iframe"), util.StrToBytes("</legend"), util.StrToBytes("</li"), util.StrToBytes("</link"), util.StrToBytes("</main"), util.StrToBytes("</menu"), util.StrToBytes("</menuitem"), util.StrToBytes("</nav"), util.StrToBytes("</noframes"), util.StrToBytes("</ol"), util.StrToBytes("</optgroup"), util.StrToBytes("</option"), util.StrToBytes("</p"), util.StrToBytes("</param"), util.StrToBytes("</section"), util.StrToBytes("</source"), util.StrToBytes("</summary"), util.StrToBytes("</table"), util.StrToBytes("</tbody"), util.StrToBytes("</td"), util.StrToBytes("</tfoot"), util.StrToBytes("</th"), util.StrToBytes("</thead"), util.StrToBytes("</title"), util.StrToBytes("</tr"), util.StrToBytes("</track"), util.StrToBytes("</ul"), util.StrToBytes("</video"), util.StrToBytes("</audio"),
	}
	htmlBlockSinglequote = util.StrToBytes("'")
	htmlBlockDoublequote = util.StrToBytes("\"")
	htmlBlockGreater     = util.StrToBytes(">")
)

func (t *Tree) isHTMLBlockClose(tokens []byte, htmlType int) bool {
	_ = "STUB: not implemented"
	return false
}

// 判断 IAL 打断

func (t *Tree) parseHTML(tokens []byte) (typ int) { _ = "STUB: not implemented"; return 0 }

// at least <? and a newline

func (t *Tree) isOpenTag(tokens []byte) (isOpenTag bool) { _ = "STUB: not implemented"; return false }

// < 后面不能跟空白

// 等号前面空格的情况

func (t *Tree) isCloseTag(tokens []byte) bool { _ = "STUB: not implemented"; return false }
