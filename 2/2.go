package main

import "fmt"

var A = []int{1, 1, 2, 2}
var B = []int{-2, 2, -1, 1}
var m, n int
var rect = [][]int{}

func isOutOfIndex(a, b int) bool {
	if a < 0 || b < 0 {
		return true
	}

	if a >= m || b >= n {
		return true
	}
	return false
}

const INT_MAX = int(^uint(0) >> 1)

var sum = 0
var min = ^INT_MAX

func dfs(k, v int) {
	sum += rect[k][v]
	if k == m-1 {
		if sum < min {
			min = sum
		} else {
			sum -= rect[k][v]
			return
		}
	}

	for i := 0; i < 4; i++ {
		a := k + A[i]
		b := v + B[i]
		if isOutOfIndex(a, b) {
			continue
		}
		dfs(a, b)
	}
	sum -= rect[k][v]

}

func main() {
	for i := 0; i < n; i++ {
		dfs(0, i)
	}
	fmt.Println(min)
}

//最短路没写出来
//附加题 直接拿矩阵数组做全局标记判环,无负环的时候求最短路径方法相同
