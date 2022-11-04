package main

// https://leetcode.com/problems/decode-string/solutions/
//type MyTreeNode struct {
//	val  string
//	rNum int
//	subs []*MyTreeNode
//	p    *MyTreeNode
//}
//
//func decodeString(s string) string {
//	root := &MyTreeNode{
//		subs: []*MyTreeNode{},
//	}
//	cur := root
//	preIdx := 0
//	for i, c := range s {
//		if c == ']' {
//			str := s[preIdx:i]
//			cur.subs = append(cur.subs, &MyTreeNode{
//				val:  str,
//				subs: []*MyTreeNode{},
//				p:    cur,
//			})
//			cur = cur.p
//			preIdx = i + 1
//			continue
//		}
//		if c == '[' {
//			k := i - 1
//			for ; k >= preIdx; k-- {
//				if !isDigit(s[k]) {
//					break
//				}
//			}
//			cur.subs = append(cur.subs, &MyTreeNode{
//				val:  string(s[preIdx : k+1]),
//				subs: []*MyTreeNode{},
//				p:    cur,
//			})
//			n, _ := strconv.ParseInt(s[k+1:i], 10, 64)
//			newNode := &MyTreeNode{
//				subs: []*MyTreeNode{},
//				p:    cur,
//				rNum: int(n),
//			}
//			cur.subs = append(cur.subs, newNode)
//			cur = newNode
//			preIdx = i + 1
//			continue
//		}
//	}
//	root.subs = append(cur.subs, &MyTreeNode{
//		val: string(s[preIdx:len(s)]),
//		p:   root,
//	})
//	travels(root)
//	return root.val
//}
//
//func travels(node *MyTreeNode) {
//	if node == nil {
//		return
//	}
//	var builder strings.Builder
//	for _, n := range node.subs {
//		if n.rNum >= 1 {
//			travels(n)
//		}
//		builder.WriteString(n.val)
//	}
//	node.val = builder.String()
//	if node.rNum > 1 {
//		builder.Reset()
//		for i := 0; i < node.rNum; i++ {
//			builder.WriteString(node.val)
//		}
//		node.val = builder.String()
//	}
//}
//
//func isDigit(c uint8) bool {
//	return c >= '0' && c <= '9'
//}
//
//func main() {
//	println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
//}
