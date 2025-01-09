package main

type TrieNode struct {
	count int
	links [26]*TrieNode
}

func (t *TrieNode) addWord(word string) {
	current := t
	for _, char := range word {
		if current.links[char-'a'] == nil {
			current.links[char-'a'] = &TrieNode{}
		}
		current = current.links[char-'a']
		current.count++
	}
}

func (t *TrieNode) countPrefix(pref string) int {
	current := t

	for _, char := range pref {
		if current.links[char-'a'] == nil {
			return 0
		}
		current = current.links[char-'a']
	}

	return current.count
}
