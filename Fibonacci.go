package main

import "fmt"

// fibonacci in GOlang
// fibo function returns Iterative version
func fibo() func() int {
	F1 := 0 //Seed value
	F2 := 1 //Seed value
	return func() int {
		F2, F1 = F1, F1+F2
		return F2
	}
}

func main() {
	n := fibo()
	//45 is the largest int for GOLang
	for i := 0; i < 45; i++ {
		fmt.Println(n())
	}
}
