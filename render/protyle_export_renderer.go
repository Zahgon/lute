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
	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

type ProtyleExportRenderer struct {
	*BaseRenderer
}

func NewProtyleExportRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) *ProtyleExportRenderer {
	_ = "STUB: not implemented"
	return nil
}

func (r *ProtyleExportRenderer) renderCallout(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCustomBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderAttributeView(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTextMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBr(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderUnderline(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderUnderlineOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderUnderlineCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderKbd(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderKbdOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderKbdCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockQueryEmbed(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockQueryEmbedScript(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderVideo(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderAudio(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderWidget(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderIFrame(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) escapeRefText(refText string) string {
	_ = "STUB: not implemented"
	return ""
}

func (r *ProtyleExportRenderer) renderBlockRefID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockRefSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockRefText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockRefDynamicText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderFileAnnotationRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderFileAnnotationRefID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderFileAnnotationRefSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderFileAnnotationRefText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderGitConflictCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderGitConflictContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderGitConflictOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderGitConflict(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTag(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTagOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTagCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSuperBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSuperBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSuperBlockLayoutMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSuperBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLinkRefDefBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLinkRefDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderKramdownBlockIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderKramdownSpanIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMark1OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMark1CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMark2OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMark2CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSup(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSupOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSupCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSub(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSubOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSubCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderYamlFrontMatterCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderYamlFrontMatterContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderYamlFrontMatterOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderYamlFrontMatter(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderHtmlEntity(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBackslashContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBackslash(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderToC(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderFootnotesDefBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderFootnotesDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"

	// r.WriteString("<li id=\"footnotes-def-" + node.FootnotesRefId + "\">")
	// 在 li 上带 id 后，Pandoc HTML 转换 Docx 会有问题
	return *new(ast.WalkStatus)
}

// 回跳 ID 会重复就先不考虑了
//gotoRef := " <a href=\"#footnotes-ref-" + node.FootnotesRefId + "\" class=\"vditor-footnotes__goto-ref\">↩</a>"
//link := &ast.Node{Type: ast.NodeInlineHTML, Tokens: util.StrToBytes(gotoRef)}
//node.FirstChild.AppendChild(link)

func (r *ProtyleExportRenderer) renderFootnotesRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeBlockInfoMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeBlockCode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmojiAlias(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmojiImg(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmojiUnicode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmoji(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderInlineMath(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderInlineMathOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderInlineMathContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderInlineMathCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMathBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMathBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMathBlockContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderMathBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTableCell(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTableRow(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTableHead(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTable(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrikethrough(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrikethrough1OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrikethrough1CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrikethrough2OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrikethrough2CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLinkTitle(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLinkDest(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLinkSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLinkText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCloseParen(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderOpenParen(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderLess(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderGreater(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCloseBrace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderOpenBrace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCloseBracket(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderOpenBracket(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBang(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderImage(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 手动设置了位置

func (r *ProtyleExportRenderer) renderLink(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderInlineHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderDocument(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderParagraph(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeSpan(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeSpanOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeSpanContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderCodeSpanCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmphasis(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmAsteriskOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmAsteriskCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmUnderscoreOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderEmUnderscoreCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrong(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrongA6kOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrongA6kCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrongU8eOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderStrongU8eCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockquote(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderBlockquoteMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderHeading(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderHeadingC8hMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderHeadingID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderList(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderListItem(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderTaskListItemMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderThematicBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderHardBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) renderSoftBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportRenderer) spanNodeAttrs(node *ast.Node, attrs *[][]string) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportRenderer) blockNodeAttrs(node *ast.Node, attrs *[][]string, class string) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportRenderer) nodeClass(node *ast.Node, attrs *[][]string, class string) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportRenderer) nodeDataType(node *ast.Node, attrs *[][]string) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportRenderer) nodeID(node *ast.Node, attrs *[][]string) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportRenderer) spellcheck(attrs *[][]string) { _ = "STUB: not implemented"; return }

func (r *ProtyleExportRenderer) contenteditable(node *ast.Node, attrs *[][]string) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportRenderer) renderIAL(node *ast.Node) { _ = "STUB: not implemented"; return }

func (r *ProtyleExportRenderer) renderTextMarkAttrs(node *ast.Node) (attrs [][]string) {
	_ = "STUB: not implemented"
	return nil
}

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227
