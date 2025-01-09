package main

func prefixCount(words []string, pref string) int {
	trie := &TrieNode{}

	for _, word := range words {
		trie.addWord(word)
	}

	return trie.countPrefix(pref)
}
