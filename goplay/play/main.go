package main

import "fmt"
import "github.com/lawrencehuang/goplay/prime"

type Person struct {
	Name string
	Age  int8
}

func (p Person) String() string {
	return fmt.Sprintf("**%v (%v years old)", p.Name, p.Age)
}

func main() {
	person1 := Person{"John", 42}
	person2 := Person{"Mary", 18}
	fmt.Println(person1, person2)
	fmt.Println(prime.NthPrime(10))
	c := prime.Primes()
	for i := 0; i < 1000; i++ {
		fmt.Println(<-c)
	}
	//	for p := range prime.Primes(20) {
	//		fmt.Println(p)
	//	}
}
