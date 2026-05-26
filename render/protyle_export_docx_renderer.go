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

type ProtyleExportDocxRenderer struct {
	*BaseRenderer
}

func NewProtyleExportDocxRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) *ProtyleExportDocxRenderer {
	_ = "STUB: not implemented"
	return nil
}

func (r *ProtyleExportDocxRenderer) Render() (output []byte) { _ = "STUB: not implemented"; return nil }

func (r *ProtyleExportDocxRenderer) renderCallout(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCustomBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderAttributeView(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTextMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHTMLTag0(node *ast.Node, currentTextMarkType string, entering bool) {
	_ = "STUB: not implemented"
	return
}

func (r *ProtyleExportDocxRenderer) getTextMarkTextContent(node *ast.Node) (ret string) {
	_ = "STUB: not implemented"
	return ""
}

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227

func (r *ProtyleExportDocxRenderer) renderBr(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderUnderline(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderUnderlineOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderUnderlineCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderKbd(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderKbdOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderKbdCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderVideo(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderAudio(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderIFrame(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderWidget(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderGitConflictCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderGitConflictContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderGitConflictOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderGitConflict(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSuperBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSuperBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSuperBlockLayoutMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSuperBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLinkRefDefBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLinkRefDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTag(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTagOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTagCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderKramdownBlockIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderKramdownSpanIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMark1OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMark1CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMark2OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMark2CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSup(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSupOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSupCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSub(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSubOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSubCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockQueryEmbed(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockQueryEmbedScript(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockRefID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockRefSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockRefText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockRefDynamicText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFileAnnotationRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFileAnnotationRefID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFileAnnotationRefSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFileAnnotationRefText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderYamlFrontMatterCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderYamlFrontMatterContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderYamlFrontMatterOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderYamlFrontMatter(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHtmlEntity(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBackslashContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBackslash(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderToC(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFootnotesRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFootnotesDefBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderFootnotesDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"

	// r.WriteString("<li id=\"footnotes-def-" + node.FootnotesRefId + "\">")
	// 在 li 上带 id 后，Pandoc HTML 转换 Docx 会有问题
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeBlockCode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeBlockInfoMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmojiAlias(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmojiImg(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmojiUnicode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmoji(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderInlineMathCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderInlineMathContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderInlineMathOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderInlineMath(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMathBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMathBlockContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMathBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderMathBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTableCell(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTableRow(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTableHead(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTable(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrikethrough(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrikethrough1OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrikethrough1CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrikethrough2OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrikethrough2CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLinkTitle(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLinkDest(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLinkSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLinkText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCloseBrace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderOpenBrace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCloseParen(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderOpenParen(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLess(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderGreater(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCloseBracket(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderOpenBracket(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBang(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderImage(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderLink(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderInlineHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderDocument(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderParagraph(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeSpan(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeSpanOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeSpanContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderCodeSpanCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmphasis(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmAsteriskOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmAsteriskCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmUnderscoreOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderEmUnderscoreCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrong(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrongA6kOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrongA6kCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrongU8eOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderStrongU8eCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockquote(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderBlockquoteMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHeading(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHeadingC8hMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHeadingID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderList(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderListItem(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderTaskListItemMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderThematicBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderHardBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) renderSoftBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportDocxRenderer) spanNodeAttrs(node *ast.Node, attrs *[][]string) {
	_ = "STUB: not implemented"
	return
}
