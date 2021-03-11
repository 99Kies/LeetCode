package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)

	go work(1, 3)
	go work(4, 6)
	go work2()
	fmt.Println("hello, world")
	//time.Sleep(time.Second)
	wg.Wait()
}

func work(start int, end int) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		fmt.Println(i)
		//time.Sleep(time.Second)
	}
}

func work2() {
	defer wg.Done()
	fmt.Println("hello, goroutine.")
}

//package main
//
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan string)
//
//}
