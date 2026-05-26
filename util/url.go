// Lute - 一款结构化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

// 这个文件搬运自标准库，为了减少依赖，生成的 js 更小。

type encoding int

const (
	encodePath encoding = 1 + iota
	encodePathSegment
	encodeHost
	encodeZone
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

type EscapeError string

func (e EscapeError) Error() string { _ = "STUB: not implemented"; return "" }

type InvalidHostError string

func (e InvalidHostError) Error() string { _ = "STUB: not implemented"; return "" }

// PathUnescape does the inverse transformation of PathEscape,
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB. It returns an error if any % is not followed
// by two hexadecimal digits.
//
// PathUnescape is identical to QueryUnescape except that it does not
// unescape '+' to ' ' (space).
func PathUnescape(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// unescape unescapes a string; the mode specifies
// which section of the URL string is being unescaped.
func unescape(s string, mode encoding) (string, error) {
	_ = "STUB: not implemented"
	// Count %, check that they're well-formed.
	return "", nil
}

// Per https://tools.ietf.org/html/rfc3986#page-21
// in the host component %-encoding can only be used
// for non-ASCII bytes.
// But https://tools.ietf.org/html/rfc6874#section-2
// introduces %25 being allowed to escape a percent sign
// in IPv6 scoped-address literals. Yay.

// RFC 6874 says basically "anything goes" for zone identifiers
// and that even non-ASCII can be redundantly escaped,
// but it seems prudent to restrict %-escaped bytes here to those
// that are valid host name bytes in their unescaped form.
// That is, you can use escaping in the zone identifier but not
// to introduce bytes you couldn't just write directly.
// But Windows puts spaces here! Yay.

// PathEscape escapes the string so it can be safely placed
// inside a URL path segment.
func PathEscape(s string) string { _ = "STUB: not implemented"; return "" }

func escape(s string, mode encoding) string { _ = "STUB: not implemented"; return "" }

// Return true if the specified character should be escaped when
// appearing in a URL string, according to RFC 3986.
//
// Please be informed that for now shouldEscape does not check all
// reserved characters correctly. See golang.org/issue/5684.
func shouldEscape(c byte, mode encoding) bool {
	_ = "STUB: not implemented"
	// §2.3 Unreserved characters (alphanum)
	return false
}

// §3.2.2 Host allows
//	sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
// as part of reg-name.
// We add : because we include :port as part of host.
// We add [ ] because we include [ipv6]:port as part of host.
// We add < > because they're the only characters left that
// we could possibly allow, and Parse will reject them if we
// escape them (because hosts can't use %-encoding for
// ASCII bytes).

// §2.3 Unreserved characters (mark)

// §2.2 Reserved characters (reserved)
// Different sections of the URL allow a few of
// the reserved characters to appear unescaped.

// §3.3
// The RFC allows : @ & = + $ but saves / ; , for assigning
// meaning to individual path segments. This package
// only manipulates the path as a whole, so we allow those
// last three as well. That leaves only ? to escape.

// §3.3
// The RFC allows : @ & = + $ but saves / ; , for assigning
// meaning to individual path segments.

// §3.2.1
// The RFC allows ';', ':', '&', '=', '+', '$', and ',' in
// userinfo, so we must escape only '@', '/', and '?'.
// The parsing of userinfo treats ':' as special so we must escape
// that too.

// §3.4
// The RFC reserves (so we must escape) everything.

// §4.1
// The RFC text is silent but the grammar allows
// everything, so escape nothing.

// RFC 3986 §2.2 allows not escaping sub-delims. A subset of sub-delims are
// included in reserved from RFC 2396 §2.2. The remaining sub-delims do not
// need to be escaped. To minimize potential breakage, we apply two restrictions:
// (1) we always escape sub-delims outside of the fragment, and (2) we always
// escape single quote to avoid breaking callers that had previously assumed that
// single quotes would be escaped. See issue #19917.

// Everything else must be escaped.

func ishex(c byte) bool { _ = "STUB: not implemented"; return false }

func unhex(c byte) byte { _ = "STUB: not implemented"; return 0 }
