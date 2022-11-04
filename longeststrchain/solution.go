package main

//https://leetcode.com/problems/longest-string-chain/description/
//func main() {
//	chain := longestStrChain([]string{"a", "b", "ab", "bac"})
//	println(chain)
//}
//
//func longestStrChain(words []string) int {
//	var memorizedWordLen = map[string]int{}
//	var wordLenMap = map[int][]string{}
//	for _, word := range words {
//		if wordLenMap[len(word)] == nil {
//			wordLenMap[len(word)] = make([]string, 0)
//		}
//		wordLenMap[len(word)] = append(wordLenMap[len(word)], word)
//	}
//	maxLen := 0
//	for _, word := range words {
//		maxLen = max(maxLen, max(memorizedWordLen[word], dfs(memorizedWordLen, wordLenMap, word)))
//	}
//	return maxLen
//}
//
//func dfs(memorizedWordLen map[string]int, wordLenMap map[int][]string, word string) int {
//	wLen := len(word)
//	if len(wordLenMap[wLen+1]) == 0 {
//		return 1
//	}
//	if memorizedWordLen[word] > 0 {
//		return memorizedWordLen[word]
//	}
//	word2MaxLen := 0
//	for _, word2 := range wordLenMap[wLen+1] {
//		if isPredecessor(word, word2) {
//			word2MaxLen = max(word2MaxLen, dfs(memorizedWordLen, wordLenMap, word2))
//		}
//	}
//	memorizedWordLen[word] = word2MaxLen + 1
//	return word2MaxLen + 1
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func isPredecessor(word1, word2 string) bool {
//	for idx := range word2 {
//		newW := word2[:idx] + word2[idx+1:]
//		if newW == word1 {
//			return true
//		}
//	}
//	return false
//}
