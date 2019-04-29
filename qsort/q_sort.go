package main

import "fmt"

var A = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6}

func partition(l, r int) int {
	provt := A[l]
	for l < r {
		for l < r && A[r] > provt {
			r--
		}
		if l < r {
			A[l] = A[r]
			l++
		}
		for l < r && A[l] < provt {
			l++
		}
		if l < r {
			A[r] = A[l]
			r--
		}
	}
	A[l] = provt
	return l
}

func q_sort(l, r int) {
	if l >= r {
		return
	}
	mid := partition(l, r)
	q_sort(l, mid-1)
	q_sort(mid+1, r)
}

func main() {
	q_sort(0, len(A)-1)
	fmt.Println(A)
}
