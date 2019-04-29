package main

import "fmt"

var A = []int{3, 1, 4, 1, 5}
var B = []int{3, 1, 4, 1, 5}

// func f(A, B []int) []int {
// 	m := make(map[int]int)

// 	ans := make([]int, 0)
// 	for _, k := range A {
// 		m[k] = 1
// 	}
// 	for _, k := range B {
// 		if _, ok := m[k]; ok {
// 			ans = append(ans, k)
// 		}
// 	}

// 	return ans
// }

func f(A, B []int) []int {
	ans := make([]int, 0)

	i1 := 0
	i2 := 0
	for {
		if A[i1] == B[i2] {
			ans = append(ans, A[i1])
			i1++
			i2++
		} else if A[i1] < B[i2] {
			i1++
		} else {
			i2++
		}

		if i1 >= len(A) || i2 >= len(B) {
			break
		}

	}
	return ans
}

func main() {
	fmt.Print(f(A, B))
}
