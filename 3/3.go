package main

var perm = [][]int{
	{0, 1, 2},
	{0, 1, 3},
	{0, 1, 4},
	{0, 2, 3},
	{0, 2, 4},
	{0, 3, 4},
	{1, 2, 3},
	{1, 2, 4},
	{1, 3, 4},
	{2, 3, 4},
}

//判断牛
func cow(a []int) (bool, int) {
	for i := 0; i < 10; i++ {
		k := a[perm[i][0]] + a[perm[i][1]] + a[perm[i][2]]
		if k == 10 || k == 20 || k == 30 {
			return true, perm[i][0] + perm[i][1] + perm[i][2]
		}
	}
	return false, 0
}

//转成整数算牛
func convert(v string) []int {
	score := make([]int, 5)

	for i, c := range v {
		if c >= '0' && c <= '9' {
			score[i] = int(c - '0')
		}
		if c == 'T' || c == 'J' || c == 'Q' || c == 'K' {
			score[i] = 10
		} else {
			panic("wrong param")
		}
	}

	return score
}

//转成字节 挨个比大小
func convertChar(b byte) byte {
	if b == 'A' {
		return '9' + 100
	}
	if b == 'T' {
		return '9' + 50
	}
	if b == 'J' {
		return '9' + 97
	}
	if b == 'Q' {
		return '9' + 98
	}
	if b == 'K' {
		return '9' + 99
	}
	if b >= '0' && b <= '9' {
		return b
	}
	panic("wrong param")
}

//假设排好序了 都是普通asc码
func comparestring(a, b string) int {
	s1 := []byte(a)
	s2 := []byte(b)
	for i := 0; i < 5; i++ {
		c1 := convertChar(s1[i])
		c2 := convertChar(s2[i])
		if c1 > c2 {
			return 1
		} else if c1 < c2 {
			return -1
		}

	}
	return 0
}

func compare(v1, v2 string) int {

	if len(v1) != 5 || len(v2) != 5 {
		panic("wrong param")
	}
	score1 := convert(v1)
	score2 := convert(v2)

	sum1 := 0
	sum2 := 0
	for i := 0; i < 5; i++ {
		sum1 += score1[i]
		sum2 += score2[i]
	}

	c1, cow1sum := cow(score1)
	c2, cow2sum := cow(score2)

	cow1sum = (sum1 - cow1sum) % 10
	cow2sum = (sum2 - cow2sum) % 10

	if cow1sum == 0 {
		cow1sum = 100
	}

	if cow2sum == 0 {
		cow2sum = 100
	}
	if c1 && c2 { //都有牛
		if cow1sum == cow2sum {
			return 0
		}
		if cow1sum > cow2sum {
			return 1
		}
		return 2

	} else if !c1 && !c2 { //都没牛
		//比牌
		return comparestring(v1, v2)
	} else if c1 {
		return 1
	}
	// c2 true
	return -1

}

func main() {

}
