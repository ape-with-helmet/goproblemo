package main

import "fmt"

func fact(a int) {
	if a == 0 {
		return 1
	}
	return a * fact(a-1)
}
func main() {
	fmt.Println("Enter a few numbers")
	var n int
	fmt.Scan(&n)
	fmt.Print(fact(n))
}
