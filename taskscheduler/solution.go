package main

//
////https://leetcode.com/problems/task-scheduler/description/
//type taskInfo struct {
//	count   int
//	nextPos int
//}
//
//func leastInterval(tasks []byte, n int) int {
//	infoMap := map[byte]*taskInfo{}
//	for _, b := range tasks {
//		if infoMap[b] == nil {
//			infoMap[b] = &taskInfo{}
//		}
//		infoMap[b].count += 1
//	}
//	unitCnt := 0
//	unfinishedTaskCnt := len(tasks)
//	for i := unfinishedTaskCnt; i > 0; {
//		for _, tInfo := range infoMap {
//			if tInfo.count == 0 || tInfo.nextPos > unitCnt {
//				continue
//			}
//			tInfo.count -= 1
//			tInfo.nextPos += n + 1
//			i--
//			break
//		}
//		unitCnt += 1
//	}
//	return unitCnt
//}
