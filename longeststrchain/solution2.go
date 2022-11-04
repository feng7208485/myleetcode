package main

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	wDic := map[string]bool{}
	for _, word := range words {
		wDic[word] = true
	}
	memorizedWordLen := map[string]int{}
	totalMax := 0
	for _, word := range words {
		maxPreLen := 0
		for i := range word {
			preWord := word[:i] + word[i+1:]
			if wDic[preWord] {
				maxPreLen = max(maxPreLen, max(memorizedWordLen[preWord], 1))
			}
		}
		memorizedWordLen[word] = maxPreLen + 1
		totalMax = max(totalMax, memorizedWordLen[word])
	}
	return totalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
