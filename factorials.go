package main

import "fmt"
func main() {
	fmt.Println("Enter a few numbers")
	var a int
	fact:=1
	fmt.Scan(&a)
	for i:=a;i>0;i--{
	    fact=fact*i
	}
	fmt.Print(fact)
}
