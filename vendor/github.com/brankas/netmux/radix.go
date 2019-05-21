package netmux

import (
	"bytes"
)

// radixNode is a radix (patricia) tree node.
type radixNode struct {
	// prefix is the node prefix.
	prefix []byte

	// len is the length of prefix.
	len int

	// next contains any successive radix nodes.
	next map[byte]*radixNode

	// terminal indicates node is also terminal for a prefix.
	terminal bool
}

// newRadixNode creates a new radix (patricia) node for the provided prefixes.
func newRadixNode(prefixes ...[]byte) *radixNode {
	switch len(prefixes) {
	case 0:
		return &radixNode{
			terminal: true,
		}
	case 1:
		return &radixNode{
			prefix:   prefixes[0],
			len:      len(prefixes[0]),
			terminal: true,
		}
	}

	rn := new(radixNode)
	rn.prefix, prefixes = splitPrefixes(prefixes...)

	nodes := make(map[byte][][]byte)
	for _, prefix := range prefixes {
		if len(prefix) == 0 {
			rn.terminal = true
			continue
		}
		nodes[prefix[0]] = append(nodes[prefix[0]], prefix[1:])
	}

	rn.next = make(map[byte]*radixNode)
	for first, prefixes := range nodes {
		rn.next[first] = newRadixNode(prefixes...)
	}

	rn.len = len(rn.prefix)
	return rn
}

// splitPrefixes splits prefixes.
func splitPrefixes(prefixes ...[]byte) ([]byte, [][]byte) {
	switch {
	case len(prefixes) == 0 || len(prefixes[0]) == 0:
		return nil, prefixes
	case len(prefixes) == 1:
		return prefixes[0], nil
	}

	var prefix []byte
	for i := 0; ; i++ {
		var cur byte
		var matched bool
		for j, prefix := range prefixes {
			switch {
			case len(prefix) <= i:
				matched = true
				break

			case j == 0:
				cur = prefix[i]
				continue

			case cur != prefix[i]:
				matched = true
				break
			}
		}

		if matched {
			break
		}

		prefix = append(prefix, cur)
	}

	next := make([][]byte, 0, len(prefixes))
	for _, remaining := range prefixes {
		next = append(next, remaining[len(prefix):])
	}
	return prefix, next
}

// match recursively searches for buf, returning true when the end of buf
// matches a terminal node of the tree or begins with a prefix in the tree.
//
// When prefix is true, and match has reached a node with no children, returns
// true if buf begins with the specified prefix.
//
// Matches prefixes only when prefix is true.
func (rn *radixNode) match(buf []byte, prefix bool) bool {
	n := len(buf)
	switch {
	case rn.len > 0 && !bytes.Equal(buf[:min(rn.len, n)], rn.prefix):
		return false

	case rn.terminal && (prefix || rn.len == n):
		return true

	case rn.len >= n:
		return false
	}

	next, ok := rn.next[buf[rn.len]]
	if !ok {
		return false
	}
	if rn.len == n {
		return next.match(buf[rn.len:rn.len], prefix)
	}
	return next.match(buf[rn.len+1:], prefix)
}

// radixTree is a simple radix tree that handles []byte instead of string
// and cannot be changed after instantiation.
type radixTree struct {
	root *radixNode
	max  int // max depth of the tree.
}

// newRadixTree creates a radix (patricia) tree for the supplied prefixes.
func newRadixTree(prefixes ...[]byte) *radixTree {
	var n int
	for _, prefix := range prefixes {
		n = max(n, len(prefix))
	}
	return &radixTree{
		root: newRadixNode(prefixes...),
		max:  n + 1,
	}
}

// newRadixTreeString creates a radix (patricia) tree for the supplied
// prefixes.
func newRadixTreeString(prefixStrings ...string) *radixTree {
	prefixes := make([][]byte, len(prefixStrings))
	for i, prefix := range prefixStrings {
		prefixes[i] = []byte(prefix)
	}
	return newRadixTree(prefixes...)
}

// match recursively searches for buf, returning true when the end of buf
// matches a terminal node of the tree or begins with a prefix in the tree.
//
// When prefix is true, and match has reached a node with no
// children, returns true if buf begins with the specified prefix.
//
// Matches prefixes only when prefix is true.
func (rt *radixTree) match(buf []byte, prefix bool) bool {
	return rt.root.match(buf[:min(len(buf), rt.max)], prefix)
}

// min returns the minimum of a, b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of a, b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
