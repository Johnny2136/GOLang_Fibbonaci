package main

import (
	"fmt"
	"time"
)

// Fibonacci in GOlang analysis of Recursive vs Iterative Algorithms:
// fibo function returns Iterative version
func fiboIt() func() int {
	F1 := 0 //Seed value
	F2 := 1 //Seed value
	return func() int {
		F2, F1 = F1, F1+F2
		return F2
	}
}

// fibo function returns Recursive Solution
func fiboRec(n int) int {
	if n < 2 {
		return n
	}
	return fiboRec(n-1) + fiboRec(n-2)
}

func main() {
	t1 := time.Now()
	n := fiboIt()
	//45 is the largest int for GOLan type int but we will use 25
	for i := 0; i < 25; i++ {
		fmt.Println(n())
	}
	itTime := time.Since(t1)
	fmt.Println("Iterative func time", itTime)

	t2 := time.Now()

	for i := 0; i < 25; i++ {
		fmt.Println(fiboRec(i))
	}
	recTime := time.Since(t2)
	fmt.Println("Recursive func time", recTime)
	// Prints output
	fmt.Println("Iterative func time", itTime)
	subtract := (recTime) - (itTime)
	fmt.Printf("Iterative func faster than Recursive func by = %v\n", subtract)

}
