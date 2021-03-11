package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		c <- "1. Golang"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		fmt.Println("Message: " + <-c)
	}()

	wg.Wait()
}
