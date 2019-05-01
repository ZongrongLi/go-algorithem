package main

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

var cnt sync.Map

func inittask(tasks chan task, r chan int, p int) {

	l := 0
	begin := l
	for l < p {
		end := Min(begin+10, p)
		//fmt.Println("======================", begin, end)

		t := task{
			begin:  begin,
			end:    end,
			result: r,
		}

		tasks <- t

		begin = end
		if begin == p {
			break
		}
	}
	fmt.Println("======================clsoe")
	close(tasks)
}

func do(t *task) {
	sum := 0
	for i := t.begin; i < t.end; i++ {
		k, ok := cnt.Load(i)
		if !ok {
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", i)
			continue
		}
		tmp := k.(int)
		tmp++
		cnt.Store(i, tmp)
		sum += i
	}

	t.result <- sum

}

func process(tasks chan task, wait *sync.WaitGroup, index int) {
	for t := range tasks {
		fmt.Println(index)
		do(&t)
	}
	wait.Done()
}

func getresult(result chan int) {
	sum := 0
	for k := range result {
		sum += k
	}
	fmt.Println(sum)

}
func distributeTask(numbers int, result chan int, tasks chan task) {
	wait := sync.WaitGroup{}

	for i := 0; i < numbers; i++ {
		wait.Add(1)
		go process(tasks, &wait, i)
	}
	wait.Wait()
	close(result)
	fmt.Println("======================clsoe1")

}

func main() {
	for i := 0; i <= 200; i++ {
		cnt.Store(i, 0)
	}

	tasks := make(chan task, 11)
	result := make(chan int, 5)
	go inittask(tasks, result, 101)

	go distributeTask(5, result, tasks)

	getresult(result)

	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(cnt.Load(i))
	// }

}
