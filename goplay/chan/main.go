package main

import (
	"fmt"

	"github.com/lawrencehuang/goplay/prime"
)

type Request struct {
	args       int
	f          func(int) prime.PrimeNum
	resultChan chan prime.PrimeNum
}

func main() {
	r := []*Request{}
	requests := make(chan *Request)
	//	quit := make(chan bool)

	go handler(requests)
	for i := 0; i < 20; i++ {
		req := &Request{i*10 + 5, prime.NthPrime, make(chan prime.PrimeNum)}
		r = append(r, req)
		requests <- req
		fmt.Println(<-req.resultChan)
	}

	// 	for _, v := range r {
	// 		fmt.Println(<-v.resultChan)
	// 	}
}

func handler(r chan *Request) {
	for req := range r {
		req.resultChan <- req.f(req.args)
	}
}

/*
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 1; i < 10; i++ {
		ch <- i
	}
	quit <- true // true or false does not matter
}

func shower(ch chan int, quit chan bool) {
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-quit:
			break
		}
	}
}
*/
