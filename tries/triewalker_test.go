package tries

import "testing"

// TestTrieWalker_New_Empty tests the NewTrieWalker function.
// It creates a trie walker with an empty trie and checks whether it is both at the root and at a leaf.
func TestTrieWalker_New_Empty(t *testing.T) {
	trie := NewTrie()
	walker := NewTrieWalker(trie)
	if !walker.IsAtRoot() {
		t.Error("New trie walker should be at root")
	}
	if !walker.IsAtLeaf() {
		t.Error("New trie walker should be at leaf")
	}
}

// TestTrieWalker_New_nonempty tests the NewTrieWalker function.
// It creates a trie walker and inserts a node into the trie.
// It checks whether the walker is at the root and not at a leaf.
func TestTrieWalker_New_nonempty(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)
	if !walker.IsAtRoot() {
		t.Error("New trie walker should be at root")
	}
	if walker.IsAtLeaf() {
		t.Error("New trie walker should not be at leaf")
	}
}

// TestTrieWalker_Step tests the Step function.
// It creates a new trie walker, adds a node to the trie, and steps past it.
// It checks whether the following:
//   - the Step function returns true up to the node
//   - the Step function returns false past the node
//   - each step moves the walker to the correct node
func TestTrieWalker_Step(t *testing.T) {
	trie := NewTrie()
	trie.Insert("abc")
	walker := NewTrieWalker(trie)

	a := trie.root.children['a']
	b := a.children['b']
	c := b.children['c']

	if !walker.Step('a') {
		t.Error("Step should return true for rune 'a'")
	}
	if walker.current != a {
		t.Error("Step should move walker to correct node")
	}

	if !walker.Step('b') {
		t.Error("Step should return true for rune 'b'")
	}
	if walker.current != b {
		t.Error("Step should move walker to correct node")
	}

	if !walker.Step('c') {
		t.Error("Step should return true for rune 'c'")
	}

	if walker.Step('d') {
		t.Error("Step should return false for rune 'd'")
	}
	if walker.current != c {
		t.Error("Step should not move walker past leaf node")
	}
}