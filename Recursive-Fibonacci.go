package main

import "fmt"

// fibonacci in GOlang
// fibo function returns Recursive Solution
func fibo(n int) int {
	if n < 2 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {

	for i := 0; i < 25; i++ {
		fmt.Println(fibo(i))
	}
}
