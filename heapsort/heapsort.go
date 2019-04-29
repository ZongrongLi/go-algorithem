package main

import "fmt"

var A = []int{0, 3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6}

func adjust(l int) {
	ni := l/2 + 1
	for i := ni; i >= 1; i-- {
		key := A[i]
		k := i
		j := k * 2
		for j < l {
			if j+1 < l && A[j] > A[j+1] {
				j++
			}
			if A[j] < A[k] {
				A[k] = A[j]
				k = j
				j *= 2
			} else {
				break
			}

		}
		A[k] = key
	}
}

func main() {
	adjust(len(A))
	for i := len(A) - 1; i >= 1; i-- {
		fmt.Print(A[1], " ")
		A[1] = A[i]
		adjust(i)
	}
}
