package main

//https://leetcode.com/problems/magic-squares-in-grid/
func main() {
	grid := [][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}
	inside := numMagicSquaresInside(grid)
	println(inside)
}

func numMagicSquaresInside(grid [][]int) int {
	count := 0
	row := len(grid)
	col := len(grid[0])
	for i := 1; i < row-1; i++ {
		for j := 1; j < col-1; j++ {
			if grid[i][j] != 5 {
				continue
			}
			dupCheck := make([]int, 10, 10)
			flag := true
			for m := i - 1; m <= i+1; m++ {
				for n := j - 1; n <= j+1; n++ {
					if grid[m][n] > 9 || grid[m][n] < 1 {
						flag = false
						break
					}
					if dupCheck[grid[m][n]] > 0 {
						flag = false
						break
					}
					dupCheck[grid[m][n]] = 1
				}
			}
			if !flag {
				continue
			}
			if grid[i][j-1]+grid[i][j]+grid[i][j+1] != 15 {
				continue
			}
			if grid[i-1][j]+grid[i][j]+grid[i+1][j] != 15 {
				continue
			}
			if grid[i-1][j-1]+grid[i][j]+grid[i+1][j+1] != 15 {
				continue
			}
			if grid[i+1][j-1]+grid[i][j]+grid[i-1][j+1] != 15 {
				continue
			}
			if grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1] != 15 {
				continue
			}
			if grid[i-1][j-1]+grid[i][j-1]+grid[i+1][j-1] != 15 {
				continue
			}
			if grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1] != 15 {
				continue
			}
			if grid[i-1][j+1]+grid[i][j+1]+grid[i+1][j+1] != 15 {
				continue
			}
			count += 1
		}
	}
	return count
}
