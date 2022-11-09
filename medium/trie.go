package main

func main() {
	obj := Constructor()
	obj.Insert("apple")

}

type Trie struct {
	children [26]*Trie
	end      bool
}

func Constructor() Trie {
	node := Trie{}

	return node
}

func (this *Trie) Insert(word string) {
	node := this

	for _, v := range word {

		index := v - 'a'

		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}

		node = node.children[index]
	}

	node.end = true
}

func (this *Trie) Search(word string) bool {
	node := this

	for _, v := range word {

		index := v - 'a'

		if node.children[index] == nil {
			return false
		}

		node = node.children[index]
	}

	return node.end
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this

	for _, v := range prefix {

		index := v - 'a'

		if node.children[index] == nil {
			return false
		}

		node = node.children[index]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// SOLUTION USING MAP
//  type Trie struct {
// 	IsTerminated bool
// 	Children     map[rune]*Trie
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	m := make(map[rune]*Trie)
// 	return Trie{IsTerminated: false, Children: m}
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string) {
// 	parent := this
// 	for _, ch := range word {
// 		if child, ok := parent.Children[ch]; ok {
// 			parent = child
// 		} else {
// 			newChild := &Trie{IsTerminated: false, Children: make(map[rune]*Trie)}
// 			parent.Children[ch] = newChild
// 			parent = newChild
// 		}
// 	}
// 	parent.IsTerminated = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	parent := this
// 	for _, ch := range word {
// 		if child, ok := parent.Children[ch]; ok {
// 			parent = child
// 			continue
// 		}
// 		return false
// 	}
// 	return parent.IsTerminated
// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	parent := this
// 	for _, ch := range prefix {
// 		if child, ok := parent.Children[ch]; ok {
// 			parent = child
// 			continue
// 		}
// 		return false
// 	}
// 	return true
// }
