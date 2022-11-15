package main

func leastInterval(tasks []byte, n int) int {
	taskCnt := make([]int, 26)
	taskMinPos := make([]int, 26)
	for _, b := range tasks {
		taskCnt[b-'A']++
	}
	unitCnt := 0
	unfinishedTaskCnt := len(tasks)
	for i := unfinishedTaskCnt; i > 0; {
		for j := 0; j <= 25; j++ {
			if taskCnt[j] == 0 || taskMinPos[j] > unitCnt {
				continue
			}
			taskCnt[j]--
			taskMinPos[j] += n + 1
			i--
			break
		}
		unitCnt++
	}
	return unitCnt
}
