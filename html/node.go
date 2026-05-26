// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
	"github.com/88250/lute/html/atom"
)

// A NodeType is the type of a Node.
type NodeType uint32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
	scopeMarkerNode
)

// Section 12.2.4.3 says "The markers are inserted when entering applet,
// object, marquee, template, td, th, and caption elements, and are used
// to prevent formatting from "leaking" into applet, object, marquee,
// template, td, th, and caption elements".
var scopeMarker = Node{Type: scopeMarkerNode}

// A Node consists of a NodeType and some Data (tag name for element nodes,
// content for text) and are part of a tree of Nodes. Element nodes may also
// have a Namespace and contain a slice of Attributes. Data is unescaped, so
// that it looks like "a<b" rather than "a&lt;b". For element nodes, DataAtom
// is the atom for Data, or zero if Data is not a known tag name.
//
// An empty Namespace implies a "http://www.w3.org/1999/xhtml" namespace.
// Similarly, "math" is short for "http://www.w3.org/1998/Math/MathML", and
// "svg" is short for "http://www.w3.org/2000/svg".
type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node

	Type      NodeType
	DataAtom  atom.Atom
	Data      string
	Namespace string
	Attr      []*Attribute
}

// Unlink 用于将节点从树上移除，后一个兄弟节点会接替该节点。
func (n *Node) Unlink() { _ = "STUB: not implemented"; return }

// InsertBefore 在当前节点前插入一个兄弟节点。
func (n *Node) InsertBefore(sibling *Node) { _ = "STUB: not implemented"; return }

// InsertAfter 在当前节点后插入一个兄弟节点。
func (n *Node) InsertAfter(sibling *Node) { _ = "STUB: not implemented"; return }

// InsertChildBefore inserts newChild as a child of n, immediately before oldChild
// in the sequence of n's children. oldChild may be nil, in which case newChild
// is appended to the end of n's children.
//
// It will panic if newChild already has a parent or siblings.
func (n *Node) InsertChildBefore(newChild, oldChild *Node) { _ = "STUB: not implemented"; return }

// AppendChild adds a node c as a child of n.
//
// It will panic if c already has a parent or siblings.
func (n *Node) AppendChild(c *Node) { _ = "STUB: not implemented"; return }

// RemoveChild removes a node c that is a child of n. Afterwards, c will have
// no parent and no siblings.
//
// It will panic if c's parent is not n.
func (n *Node) RemoveChild(c *Node) { _ = "STUB: not implemented"; return }

// reparentChildren reparents all of src's child nodes to dst.
func reparentChildren(dst, src *Node) { _ = "STUB: not implemented"; return }

// clone returns a new node with the same type, data and attributes.
// The clone has no parent, no siblings and no children.
func (n *Node) clone() *Node { _ = "STUB: not implemented"; return nil }

// nodeStack is a stack of nodes.
type nodeStack []*Node

// pop pops the stack. It will panic if s is empty.
func (s *nodeStack) pop() *Node { _ = "STUB: not implemented"; return nil }

// top returns the most recently pushed node, or nil if s is empty.
func (s *nodeStack) top() *Node { _ = "STUB: not implemented"; return nil }

// index returns the index of the top-most occurrence of n in the stack, or -1
// if n is not present.
func (s *nodeStack) index(n *Node) int { _ = "STUB: not implemented"; return 0 }

// contains returns whether a is within s.
func (s *nodeStack) contains(a atom.Atom) bool { _ = "STUB: not implemented"; return false }

// insert inserts a node at the given index.
func (s *nodeStack) insert(i int, n *Node) { _ = "STUB: not implemented"; return }

// remove removes a node from the stack. It is a no-op if n is not present.
func (s *nodeStack) remove(n *Node) { _ = "STUB: not implemented"; return }

type insertionModeStack []insertionMode

func (s *insertionModeStack) pop() (im insertionMode) {
	_ = "STUB: not implemented"
	return *new(insertionMode)
}

func (s *insertionModeStack) top() insertionMode {
	_ = "STUB: not implemented"
	return *new(insertionMode)
}
