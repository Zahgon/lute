// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"errors"
	"io"

	"github.com/88250/lute/html/atom"
)

// A TokenType is the type of a Token.
type TokenType uint32

const (
	// ErrorToken means that an error occurred during tokenization.
	ErrorToken TokenType = iota
	// TextToken means a text node.
	TextToken
	// A StartTagToken looks like <a>.
	StartTagToken
	// An EndTagToken looks like </a>.
	EndTagToken
	// A SelfClosingTagToken tag looks like <br/>.
	SelfClosingTagToken
	// A CommentToken looks like <!--x-->.
	CommentToken
	// A DoctypeToken looks like <!DOCTYPE x>
	DoctypeToken
)

// ErrBufferExceeded means that the buffering limit was exceeded.
var ErrBufferExceeded = errors.New("max buffer exceeded")

// String returns a string representation of the TokenType.
func (t TokenType) String() string { _ = "STUB: not implemented"; return "" }

// An Attribute is an attribute namespace-key-value triple. Namespace is
// non-empty for foreign attributes like xlink, Key is alphabetic (and hence
// does not contain escapable characters like '&', '<' or '>'), and Val is
// unescaped (it looks like "a<b" rather than "a&lt;b").
//
// Namespace is only used by the parser, not the tokenizer.
type Attribute struct {
	Namespace, Key, Val string
}

// A Token consists of a TokenType and some Data (tag name for start and end
// tags, content for text, comments and doctypes). A tag Token may also contain
// a slice of Attributes. Data is unescaped for all Tokens (it looks like "a<b"
// rather than "a&lt;b"). For tag Tokens, DataAtom is the atom for Data, or
// zero if Data is not a known tag name.
type Token struct {
	Type     TokenType
	DataAtom atom.Atom
	Data     string
	Attr     []*Attribute
}

// tagString returns a string representation of a tag Token's Data and Attr.
func (t Token) tagString() string { _ = "STUB: not implemented"; return "" }

// String returns a string representation of the Token.
func (t Token) String() string { _ = "STUB: not implemented"; return "" }

// span is a range of bytes in a Tokenizer's buffer. The start is inclusive,
// the end is exclusive.
type span struct {
	start, end int
}

// A Tokenizer returns a stream of HTML Tokens.
type Tokenizer struct {
	// r is the source of the HTML text.
	r io.Reader
	// tt is the TokenType of the current token.
	tt TokenType
	// err is the first error encountered during tokenization. It is possible
	// for tt != Error && err != nil to hold: this means that Next returned a
	// valid token but the subsequent Next call will return an error token.
	// For example, if the HTML text input was just "plain", then the first
	// Next call would set z.err to io.EOF but return a TextToken, and all
	// subsequent Next calls would return an ErrorToken.
	// err is never reset. Once it becomes non-nil, it stays non-nil.
	err error
	// readErr is the error returned by the io.Reader r. It is separate from
	// err because it is valid for an io.Reader to return (n int, err1 error)
	// such that n > 0 && err1 != nil, and callers should always process the
	// n > 0 bytes before considering the error err1.
	readErr error
	// buf[raw.start:raw.end] holds the raw bytes of the current token.
	// buf[raw.end:] is buffered input that will yield future tokens.
	raw span
	buf []byte
	// maxBuf limits the data buffered in buf. A value of 0 means unlimited.
	maxBuf int
	// buf[data.start:data.end] holds the raw bytes of the current token's data:
	// a text token's text, a tag token's tag name, etc.
	data span
	// pendingAttr is the attribute key and value currently being tokenized.
	// When complete, pendingAttr is pushed onto attr. nAttrReturned is
	// incremented on each call to TagAttr.
	pendingAttr   [2]span
	attr          [][2]span
	nAttrReturned int
	// rawTag is the "script" in "</script>" that closes the next token. If
	// non-empty, the subsequent call to Next will return a raw or RCDATA text
	// token: one that treats "<p>" as text instead of an element.
	// rawTag's contents are lower-cased.
	rawTag string
	// textIsRaw is whether the current text token's data is not escaped.
	textIsRaw bool
	// convertNUL is whether NUL bytes in the current token's data should
	// be converted into \ufffd replacement characters.
	convertNUL bool
	// allowCDATA is whether CDATA sections are allowed in the current context.
	allowCDATA bool
}

// AllowCDATA sets whether or not the tokenizer recognizes <![CDATA[foo]]> as
// the text "foo". The default value is false, which means to recognize it as
// a bogus comment "<!-- [CDATA[foo]] -->" instead.
//
// Strictly speaking, an HTML5 compliant tokenizer should allow CDATA if and
// only if tokenizing foreign content, such as MathML and SVG. However,
// tracking foreign-contentness is difficult to do purely in the tokenizer,
// as opposed to the parser, due to HTML integration points: an <svg> element
// can contain a <foreignObject> that is foreign-to-SVG but not foreign-to-
// HTML. For strict compliance with the HTML5 tokenization algorithm, it is the
// responsibility of the user of a tokenizer to call AllowCDATA as appropriate.
// In practice, if using the tokenizer without caring whether MathML or SVG
// CDATA is text or comments, such as tokenizing HTML to find all the anchor
// text, it is acceptable to ignore this responsibility.
func (z *Tokenizer) AllowCDATA(allowCDATA bool) { _ = "STUB: not implemented"; return }

// NextIsNotRawText instructs the tokenizer that the next token should not be
// considered as 'raw text'. Some elements, such as script and title elements,
// normally require the next token after the opening tag to be 'raw text' that
// has no child elements. For example, tokenizing "<title>a<b>c</b>d</title>"
// yields a start tag token for "<title>", a text token for "a<b>c</b>d", and
// an end tag token for "</title>". There are no distinct start tag or end tag
// tokens for the "<b>" and "</b>".
//
// This tokenizer implementation will generally look for raw text at the right
// times. Strictly speaking, an HTML5 compliant tokenizer should not look for
// raw text if in foreign content: <title> generally needs raw text, but a
// <title> inside an <svg> does not. Another example is that a <textarea>
// generally needs raw text, but a <textarea> is not allowed as an immediate
// child of a <select>; in normal parsing, a <textarea> implies </select>, but
// one cannot close the implicit element when parsing a <select>'s InnerHTML.
// Similarly to AllowCDATA, tracking the correct moment to override raw-text-
// ness is difficult to do purely in the tokenizer, as opposed to the parser.
// For strict compliance with the HTML5 tokenization algorithm, it is the
// responsibility of the user of a tokenizer to call NextIsNotRawText as
// appropriate. In practice, like AllowCDATA, it is acceptable to ignore this
// responsibility for basic usage.
//
// Note that this 'raw text' concept is different from the one offered by the
// Tokenizer.Raw method.
func (z *Tokenizer) NextIsNotRawText() {
	_ = "STUB: not implemented"

	// Err returns the error associated with the most recent ErrorToken token.
	// This is typically io.EOF, meaning the end of tokenization.
	return
}

func (z *Tokenizer) Err() error { _ = "STUB: not implemented"; return nil }

// readByte returns the next byte from the input stream, doing a buffered read
// from z.r into z.buf if necessary. z.buf[z.raw.start:z.raw.end] remains a contiguous byte
// slice that holds all the bytes read so far for the current token.
// It sets z.err if the underlying reader returns an error.
// Pre-condition: z.err == nil.
func (z *Tokenizer) readByte() byte { _ = "STUB: not implemented"; return 0 }

// Our buffer is exhausted and we have to read from z.r. Check if the
// previous read resulted in an error.

// We copy z.buf[z.raw.start:z.raw.end] to the beginning of z.buf. If the length
// z.raw.end - z.raw.start is more than half the capacity of z.buf, then we
// allocate a new buffer before the copy.

// Adjust the data/attr spans to refer to the same contents after the copy.

// Now that we have copied the live bytes to the start of the buffer,
// we read from z.r into the remainder.

// Buffered returns a slice containing data buffered but not yet tokenized.
func (z *Tokenizer) Buffered() []byte { _ = "STUB: not implemented"; return nil }

// readAtLeastOneByte wraps an io.Reader so that reading cannot return (0, nil).
// It returns io.ErrNoProgress if the underlying r.Read method returns (0, nil)
// too many times in succession.
func readAtLeastOneByte(r io.Reader, b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// skipWhiteSpace skips past any white space.
func (z *Tokenizer) skipWhiteSpace() { _ = "STUB: not implemented"; return }

// No-op.

// readRawOrRCDATA reads until the next "</foo>", where "foo" is z.rawTag and
// is typically something like "script" or "textarea".
func (z *Tokenizer) readRawOrRCDATA() { _ = "STUB: not implemented"; return }

// A textarea's or title's RCDATA can contain escaped entities.

// readRawEndTag attempts to read a tag like "</foo>", where "foo" is z.rawTag.
// If it succeeds, it backs up the input position to reconsume the tag and
// returns true. Otherwise it returns false. The opening "</" has already been
// consumed.
func (z *Tokenizer) readRawEndTag() bool { _ = "STUB: not implemented"; return false }

// The 3 is 2 for the leading "</" plus 1 for the trailing character c.

// readScript reads until the next </script> tag, following the byzantine
// rules for escaping/hiding the closing tag.
func (z *Tokenizer) readScript() { _ = "STUB: not implemented"; return }

// readComment reads the next comment token starting with "<!--". The opening
// "<!--" has already been consumed.
func (z *Tokenizer) readComment() { _ = "STUB: not implemented"; return }

// It's a comment with no data, like <!-->.

// Ignore up to two dashes at EOF.

// readUntilCloseAngle reads until the next ">".
func (z *Tokenizer) readUntilCloseAngle() { _ = "STUB: not implemented"; return }

// readMarkupDeclaration reads the next token starting with "<!". It might be
// a "<!--comment-->", a "<!DOCTYPE foo>", a "<![CDATA[section]]>" or
// "<!a bogus comment". The opening "<!" has already been consumed.
func (z *Tokenizer) readMarkupDeclaration() TokenType {
	_ = "STUB: not implemented"
	return *new(TokenType)
}

// It's a bogus comment.

// readDoctype attempts to read a doctype declaration and returns true if
// successful. The opening "<!" has already been consumed.
func (z *Tokenizer) readDoctype() bool { _ = "STUB: not implemented"; return false }

// Back up to read the fragment of "DOCTYPE" again.

// readCDATA attempts to read a CDATA section and returns true if
// successful. The opening "<!" has already been consumed.
func (z *Tokenizer) readCDATA() bool { _ = "STUB: not implemented"; return false }

// Back up to read the fragment of "[CDATA[" again.

// startTagIn returns whether the start tag in z.buf[z.data.start:z.data.end]
// case-insensitively matches any element of ss.
func (z *Tokenizer) startTagIn(ss ...string) bool { _ = "STUB: not implemented"; return false }

// readStartTag reads the next start tag token. The opening "<a" has already
// been consumed, where 'a' means anything in [A-Za-z].
func (z *Tokenizer) readStartTag() TokenType { _ = "STUB: not implemented"; return *new(TokenType) }

// Several tags flag the tokenizer's next token as raw.

// Look for a self-closing token like "<br/>".

// readTag reads the next tag token and its attributes. If saveAttr, those
// attributes are saved in z.attr, otherwise z.attr is set to an empty slice.
// The opening "<a" or "</a" has already been consumed, where 'a' means anything
// in [A-Za-z].
func (z *Tokenizer) readTag(saveAttr bool) { _ = "STUB: not implemented"; return }

// Read the tag name and attribute key/value pairs.

// Save pendingAttr if saveAttr and that attribute has a non-empty key.

// readTagName sets z.data to the "div" in "<div k=v>". The reader (z.raw.end)
// is positioned such that the first byte of the tag name (the "d" in "<div")
// has already been consumed.
func (z *Tokenizer) readTagName() { _ = "STUB: not implemented"; return }

// readTagAttrKey sets z.pendingAttr[0] to the "k" in "<div k=v>".
// Precondition: z.err == nil.
func (z *Tokenizer) readTagAttrKey() { _ = "STUB: not implemented"; return }

// readTagAttrVal sets z.pendingAttr[1] to the "v" in "<div k=v>".
func (z *Tokenizer) readTagAttrVal() { _ = "STUB: not implemented"; return }

// Next scans the next token and returns its type.
func (z *Tokenizer) Next() TokenType { _ = "STUB: not implemented"; return *new(TokenType) }

// Read everything up to EOF.

// Check if the '<' we have just read is part of a tag, comment
// or doctype. If not, it's part of the accumulated text token.

// We use CommentToken to mean any of "<!--actual comments-->",
// "<!DOCTYPE declarations>" and "<?xml processing instructions?>".

// Reconsume the current character.

// We have a non-text token, but we might have accumulated some text
// before that. If so, we return the text first, and return the non-
// text token on the subsequent call to Next.

// "</>" does not generate a token at all. Generate an empty comment
// to allow passthrough clients to pick up the data using Raw.
// Reset the tokenizer state and start again.

// Raw returns the unmodified text of the current token. Calling Next, Token,
// Text, TagName or TagAttr may change the contents of the returned slice.
func (z *Tokenizer) Raw() []byte { _ = "STUB: not implemented"; return nil }

// convertNewlines converts "\r" and "\r\n" in s to "\n".
// The conversion happens in place, but the resulting slice may be shorter.
func convertNewlines(s []byte) []byte { _ = "STUB: not implemented"; return nil }

var (
	nul         = []byte("\x00")
	replacement = []byte("\ufffd")
)

// Text returns the unescaped text of a text, comment or doctype token. The
// contents of the returned slice may change on the next call to Next.
func (z *Tokenizer) Text() []byte { _ = "STUB: not implemented"; return nil }

// TagName returns the lower-cased name of a tag token (the `img` out of
// `<IMG SRC="foo">`) and whether the tag has attributes.
// The contents of the returned slice may change on the next call to Next.
func (z *Tokenizer) TagName() (name []byte, hasAttr bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// TagAttr returns the lower-cased key and unescaped value of the next unparsed
// attribute for the current tag token and whether there are more attributes.
// The contents of the returned slices may change on the next call to Next.
func (z *Tokenizer) TagAttr() (key, val []byte, moreAttr bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// Token returns the current Token. The result's Data and Attr values remain
// valid after subsequent Next calls.
func (z *Tokenizer) Token() Token { _ = "STUB: not implemented"; return *new(Token) }

// SetMaxBuf sets a limit on the amount of data buffered during tokenization.
// A value of 0 means unlimited.
func (z *Tokenizer) SetMaxBuf(n int) {
	_ = "STUB: not implemented"

	// NewTokenizer returns a new HTML Tokenizer for the given Reader.
	// The input is assumed to be UTF-8 encoded.
	return
}

func NewTokenizer(r io.Reader) *Tokenizer { _ = "STUB: not implemented"; return nil }

// NewTokenizerFragment returns a new HTML Tokenizer for the given Reader, for
// tokenizing an existing element's InnerHTML fragment. contextTag is that
// element's tag, such as "div" or "iframe".
//
// For example, how the InnerHTML "a<b" is tokenized depends on whether it is
// for a <p> tag or a <script> tag.
//
// The input is assumed to be UTF-8 encoded.
func NewTokenizerFragment(r io.Reader, contextTag string) *Tokenizer {
	_ = "STUB: not implemented"
	return nil
}
