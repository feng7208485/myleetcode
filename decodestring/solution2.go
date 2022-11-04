package main

import "strings"

func decodeString(s string) string {
	stack := make([]interface{}, 0)
	var curStr strings.Builder
	var tmpStr strings.Builder
	curNum := 0
	for _, c := range s {
		if c == '[' {
			stack = append(stack, curStr.String())
			curStr.Reset()
			stack = append(stack, curNum)
			curNum = 0
			continue
		}
		if c == ']' {
			num := stack[len(stack)-1].(int)
			preStr := stack[len(stack)-2].(string)
			stack = stack[:len(stack)-2]
			tmpStr.WriteString(preStr)
			for i := 0; i < num; i++ {
				tmpStr.WriteString(curStr.String())
			}
			curStr.Reset()
			curStr.WriteString(tmpStr.String())
			tmpStr.Reset()
			continue
		}
		if isDigit(c) {
			curNum = curNum*10 + (int(c) - int('0'))
			continue
		}
		curStr.WriteString(string(c))
	}
	return curStr.String()
}

func isDigit(c int32) bool {
	return c >= '0' && c <= '9'
}

func main() {
	println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
}
