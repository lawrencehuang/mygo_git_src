package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w)
}

func main() {
	go ready("Tea", 1)
	go ready("Coffee", 3)
	time.Sleep(5 * time.Second)
}
