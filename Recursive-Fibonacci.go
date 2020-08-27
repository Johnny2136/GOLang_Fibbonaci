package main

import (
	"fmt"
	"time"
)

// fibonacci in GOlang
// fibo function returns Recursive Solution
func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	start := time.Now()

	for i := 0; i < 45; i++ {
		fmt.Println(fibo(i))
	}
	fmt.Println(time.Since(start))
}
