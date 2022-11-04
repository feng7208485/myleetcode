package main

import (
	"fmt"
)

func main() {
	climb(4, []int{})
	isPredecessor("", "abcd")
}

func isPredecessor(word1, word2 string) bool {
	for idx, _ := range word2 {
		newW := word2[:idx] + word2[idx+1:]
		println(newW)
	}
	return false
}

func climb(rest int, path []int) {
	if rest == 0 {
		fmt.Println(path)
		return
	}
	if rest >= 1 {
		path = append(path, 1)
		climb(rest-1, path)
		path = path[:len(path)-1]
	}
	if rest >= 2 {
		path = append(path, 2)
		climb(rest-2, path)
		path = path[:len(path)-1]
	}
}
