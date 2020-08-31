package main

import (
	"fmt"
	"time"
)

// fibonacci in GOlang
// fibo function returns Recursive Solution
func fiboRec(n int) int {
	if n < 2 {
		return n
	}
	return fiboRec(n-1) + fiboRec(n-2)
}

func main() {
	start := time.Now()

	for i := 0; i < 45; i++ {
		fmt.Println(fiboRec(i))
	}
	fmt.Println(time.Since(start))
}
