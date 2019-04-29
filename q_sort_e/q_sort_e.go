//快排的三色问题
package main

import "fmt"

var A = []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8, 4, 6}

func partition(l, r int) (int, int) {
	provt := A[l]
	m := l - 1
	for l < r {
		for l < r && A[r] > provt {
			r--
		}
		if l < r {
			if A[r] == provt {
				A[l] = A[r]
			} else {
				m++
				A[l] = provt
				A[m] = A[r]
			}
			l++
		}
		for l < r {
			if A[l] < provt {
				m++
				tmp := A[l]
				A[l] = provt
				A[m] = tmp
				l++
			} else if A[l] == provt {
				l++
			} else {
				break
			}

		}
		if l < r {
			A[r] = A[l]
			r--
		}
	}
	A[l] = provt
	return m + 1, l
}

func q_sort(l, r int) {
	if l >= r {
		return
	}
	midl, midh := partition(l, r)
	//fmt.Println(A)

	q_sort(l, midl-1)
	q_sort(midh+1, r)
}

func main() {
	q_sort(0, len(A)-1)
	fmt.Println(A)
}
