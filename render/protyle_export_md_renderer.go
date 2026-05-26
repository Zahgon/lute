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

	"github.com/88250/lute/ast"
	"github.com/88250/lute/parse"
)

type ProtyleExportMdRenderer struct {
	*BaseRenderer
	NodeWriterStack []*bytes.Buffer
}

func NewProtyleExportMdRenderer(tree *parse.Tree, options *Options, parseOptions *parse.Options) *ProtyleExportMdRenderer {
	_ = "STUB: not implemented"
	return nil
}

func (r *ProtyleExportMdRenderer) renderCallout(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCustomBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderAttributeView(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTextMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 多加一个转义符 Improve the handling of inline-code containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9252

// 填充空格以满足 Markdown 语法 https://ld246.com/article/1597581380183
// https://github.com/siyuan-note/siyuan/issues/6472
// https://github.com/siyuan-note/siyuan/issues/9542
// 这里无法使用零宽空格，只能用空格，否则 Pandoc 导出会有问题

// 通过零宽空格来区隔相邻的 Markdown 标记符

func (r *ProtyleExportMdRenderer) renderMdMarker(node *ast.Node, entering bool) (ret string) {
	_ = "STUB: not implemented"
	return ""
}

// 重新排序，将 a、inline-memo、block-ref、file-annotation-ref、inline-math 放在最前面，将 code 放在最后面

// 过滤掉 text 类型

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227

// 最内层是 code 时，需要在渲染 code 前添加零宽空格，然后再渲染 code 标记符
// Improve exporting inline code markdown element https://github.com/siyuan-note/siyuan/issues/10988

// 最内层是 code 时，需要在渲染 code 标记符后添加零宽空格，然后再渲染其他标记符

func reverse(ss []string) { _ = "STUB: not implemented"; return }

func (r *ProtyleExportMdRenderer) renderMdMarker0(node *ast.Node, currentTextmarkType string, entering bool) (ret string) {
	_ = "STUB: not implemented"
	return ""
}

// < 和 > 符号不用转义，可以符合 Markdown 规范 https://github.com/siyuan-note/siyuan/issues/15023

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227

func (r *ProtyleExportMdRenderer) renderMdMarker1(node *ast.Node, currentTextmarkType string, entering bool) (ret string) {
	_ = "STUB: not implemented"
	return ""
}

func (r *ProtyleExportMdRenderer) renderBr(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderUnderline(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderUnderlineOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderUnderlineCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderKbd(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderKbdOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderKbdCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderVideo(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderAudio(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderIFrame(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderWidget(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderGitConflictCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderGitConflictContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderGitConflictOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderGitConflict(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSuperBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSuperBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSuperBlockLayoutMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSuperBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLinkRefDefBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLinkRefDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTag(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTagOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTagCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderKramdownBlockIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderKramdownSpanIAL(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMark(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMark1OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMark1CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMark2OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMark2CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSup(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSupOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSupCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSub(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSubOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSubCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockQueryEmbedScript(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockQueryEmbed(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockRefID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockRefSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockRefText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockRefDynamicText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFileAnnotationRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFileAnnotationRefID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFileAnnotationRefSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFileAnnotationRefText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderYamlFrontMatterCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderYamlFrontMatterContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderYamlFrontMatterOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderYamlFrontMatter(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderHtmlEntity(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBackslashContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBackslash(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderToC(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFootnotesRef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFootnotesDefBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderFootnotesDef(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmojiAlias(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmojiImg(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmojiUnicode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmoji(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTableCell(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTableRow(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTableHead(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderTable(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 遍历单元格算出最大宽度

// 自动添加空格会导致单元格宽度发生变化

// 遍历字节点，将可能会多出来的空格计算出来

// 空格仅一个字节，可以直接计算长度

func (r *ProtyleExportMdRenderer) renderStrikethrough(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrikethrough1OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrikethrough1CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrikethrough2OpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrikethrough2CloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLinkTitle(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 这里不进行转义输出，仅转义双引号 https://github.com/siyuan-note/siyuan/issues/15023

func (r *ProtyleExportMdRenderer) renderLinkDest(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLinkSpace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLinkText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCloseParen(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderOpenParen(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderGreater(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLess(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCloseBrace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderOpenBrace(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCloseBracket(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderOpenBracket(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBang(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderImage(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderLink(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderInlineHTML(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderDocument(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderParagraph(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// ListItem.Paragraph

// 必须通过列表（而非列表项）上的紧凑标识判断，因为在设置该标识时仅设置了 List.Tight
// 设置紧凑标识的具体实现可参考函数 List.Finalize()

func (r *ProtyleExportMdRenderer) renderText(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"

	// 去掉开头的零宽空格 Exported Markdown inline code no longer contains zero-width spaces after it https://github.com/siyuan-note/siyuan/issues/15328
	// When exporting inline code to Markdown, try to remove the trailing zero-width space https://github.com/siyuan-note/siyuan/issues/15854
	return *new(ast.WalkStatus)
}

// 导出 Markdown 时去掉图片节点左右的零宽空格 Remove zero-width spaces around image nodes when exporting Markdown https://github.com/siyuan-note/siyuan/issues/14263

func (r *ProtyleExportMdRenderer) renderCodeSpan(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeSpanOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeSpanContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeSpanCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderInlineMath(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderInlineMathOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderInlineMathContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// Improve the handling of inline-math containing `|` in the table https://github.com/siyuan-note/siyuan/issues/9227

func (r *ProtyleExportMdRenderer) renderInlineMathCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMathBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMathBlockContent(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMathBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderMathBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeBlockCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeBlockCode(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeBlockInfoMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeBlockOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderCodeBlock(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmphasis(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmAsteriskOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmAsteriskCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmUnderscoreOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderEmUnderscoreCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrong(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrongA6kOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrongA6kCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrongU8eOpenMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderStrongU8eCloseMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderBlockquote(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 已经是根这一层

// 在表格中不能换行，否则会破坏表格的排版 https://github.com/Vanessa219/vditor/issues/368

func (r *ProtyleExportMdRenderer) renderBlockquoteMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderHeading(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderHeadingC8hMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderHeadingID(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderList(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderListItem(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

// 说明该列表项为空 https://github.com/siyuan-note/siyuan/issues/6206

func (r *ProtyleExportMdRenderer) renderTaskListItemMarker(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderThematicBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderHardBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) renderSoftBreak(node *ast.Node, entering bool) ast.WalkStatus {
	_ = "STUB: not implemented"
	return *new(ast.WalkStatus)
}

func (r *ProtyleExportMdRenderer) withoutKramdownBlockIAL(node *ast.Node) bool {
	_ = "STUB: not implemented"
	return false
}
