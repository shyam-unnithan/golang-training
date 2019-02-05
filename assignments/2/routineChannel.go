package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func getInt(out chan int) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		out <- i
	}
	close(out)
}

func sumInt(in chan int, out chan int) {
	defer wg.Done()
	for v := range in {
		v++
		out <- v
	}
	close(out)
}

func printInt(in chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	wg.Add(3)
	pushInt := make(chan int)
	pushSum := make(chan int)
	go getInt(pushInt)
	go sumInt(pushInt, pushSum)
	go printInt(pushSum)
	wg.Wait()
}
