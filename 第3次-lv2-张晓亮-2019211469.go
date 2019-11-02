package main

import (
	"fmt"
	"sync"
)
var (
	myres = make(map[int]int, 20)
	mu sync.Mutex
)
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i

	}

	myres[n] = res

}
func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 20; i++ {
		go factorial(i)
		select {
		case i = <- ch:
		case ch <- i:
		}
	}
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}

}