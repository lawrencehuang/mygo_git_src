// Package main provides ...
package main

import "fmt"

func main() {
	//	call_clojure()
	callFibGen()
}

func callFibGen() {
	n := 15
	for x := range fibGen(n) {
		fmt.Println(x)
	}
}

func fibGen(n int) chan int {
	ch := make(chan int)
	go func(n int) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			ch <- a
		}
		close(ch)
	}(n)
	return ch
}

// Call fib clusure
func callClosure() {
	f1 := fibClosure()
	for i := 0; i < 10; i++ {
		fmt.Println(f1())
	}
}
func fibClosure() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// Call fib with chan
func callChan1() {
	c := make(chan int)
	quit := make(chan bool)
	go fibChan1(c, quit)
	for i := 0; i < 10; i++ {
		c <- i
	}
	quit <- true
}
func fibChan1(c chan int, quit chan bool) {
	a, b := 0, 1
	for {
		select {
		case <-c:
			a, b = b, a+b
			fmt.Println(a)

		case <-quit:
			break
		}
	}
}

func callChan2() {
	c := make(chan int)
	go fibChan2(c)
	for i := 1; i < 20; i++ {
		c <- i
	}
	close(c)
}

func fibChan2(c chan int) {
	a, b := 0, 1
	for i := range c {
		a, b = b, a+b
		fmt.Println(i, a)
	}
}
