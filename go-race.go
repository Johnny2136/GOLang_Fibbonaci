package main

import ("fmt"
		"time")

// fibonacci in GOlang
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
	start1 := time.Now()
	n := fiboIt()
	//45 is the largest int for GOLan but we will use 25
	for i := 0; i < 40; i++ {
		fmt.Println(n())
	}
	fmt.Println("Iterative func time", time.Since(start1))
	start2 := time.Now()

	for i := 0; i < 40; i++ {
		fmt.Println(fiboRec(i))
	}
	fmt.Println("Recursive func time", time.Since(start2))
	if


}