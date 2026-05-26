// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

// These replacements permit compatibility with old numeric entities that
// assumed Windows-1252 encoding.
// https://html.spec.whatwg.org/multipage/syntax.html#consume-a-character-reference
var replacementTable = [...]rune{
	'\u20AC', // First entry is what 0x80 should be replaced with.
	'\u0081',
	'\u201A',
	'\u0192',
	'\u201E',
	'\u2026',
	'\u2020',
	'\u2021',
	'\u02C6',
	'\u2030',
	'\u0160',
	'\u2039',
	'\u0152',
	'\u008D',
	'\u017D',
	'\u008F',
	'\u0090',
	'\u2018',
	'\u2019',
	'\u201C',
	'\u201D',
	'\u2022',
	'\u2013',
	'\u2014',
	'\u02DC',
	'\u2122',
	'\u0161',
	'\u203A',
	'\u0153',
	'\u009D',
	'\u017E',
	'\u0178', // Last entry is 0x9F.
	// 0x00->'\uFFFD' is handled programmatically.
	// 0x0D->'\u000D' is a no-op.
}

// unescapeEntity reads an entity like "&lt;" from b[src:] and writes the
// corresponding "<" to b[dst:], returning the incremented dst and src cursors.
// Precondition: b[src] == '&' && dst <= src.
// attribute should be true if parsing an attribute value.
func unescapeEntity(b []byte, dst, src int, attribute bool) (dst1, src1 int) {
	_ = "STUB: not implemented"
	// https://html.spec.whatwg.org/multipage/syntax.html#consume-a-character-reference
	return 0, 0
}

// i starts at 1 because we already know that s[0] == '&'.

// We need to have at least "&#.".

// No characters matched.

// Replace characters from Windows-1252 with UTF-8 equivalents.

// Replace invalid characters with the replacement character.

// Consume the maximum number of characters possible, with the
// consumed characters matching one of the named references.

// Lower-cased characters are more common in entities, so we check for them first.

// No-op.

// No-op.

// unescape unescapes b's entities in-place, so that "a&lt;b" becomes "a<b".
// attribute should be true if parsing an attribute value.
func unescape(b []byte, attribute bool) []byte { _ = "STUB: not implemented"; return nil }

// lower lower-cases the A-Z bytes in b in-place, so that "aBc" becomes "abc".
func lower(b []byte) []byte { _ = "STUB: not implemented"; return nil }

const escapedChars = "&'<>\"\r"

func escape(w writer, s string) error { _ = "STUB: not implemented"; return nil }

// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.

// "&#34;" is shorter than "&quot;".

// EscapeString escapes special characters like "<" to become "&lt;". It
// escapes only five such characters: <, >, &, ' and ".
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
// always true.
func EscapeString(s string) string { _ = "STUB: not implemented"; return "" }

// UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a
// larger range of entities than EscapeString escapes. For example, "&aacute;"
// unescapes to "á", as does "&#225;" and "&xE1;".
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
// always true.
func UnescapeString(s string) string { _ = "STUB: not implemented"; return "" }
