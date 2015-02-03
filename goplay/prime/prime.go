//package main
package prime

import "fmt"

type PrimeNum struct {
	order int
	val   int
}

//func main() {
//	for p := range Primes(15) {
//		fmt.Println(p)
//	}
//	fmt.Println(NthPrime(1e8))
//}

func (p PrimeNum) String() string {
	return fmt.Sprintf("%d th prime = %d", p.order, p.val)
}

// find Nth Prime
func NthPrime(n int) PrimeNum {
	val := 0
	counter := 0
	for i := 2; n > 0; i++ {
		if Prime(i) {
			val = i
			n--
			counter++
		}
	}
	return PrimeNum{counter, val}
}

// find the first n primes
func Primes() chan PrimeNum {
	c := make(chan PrimeNum)
	counter := 1
	go func() {
		for i := 2; ; i++ {
			if Prime(i) {
				c <- PrimeNum{counter, i}
				counter++
			}
		}
		//		close(c)
	}()
	return c
}

func Prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
