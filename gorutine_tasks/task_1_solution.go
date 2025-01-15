package main

import (
	"fmt"
	"sync"
)

// merge - соединяет каналы в один
func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, c := range cs {
		wg.Add(1)
		go func(r <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// fillChan - заполняет канал числами от 0 до n-1
func fillChan(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
