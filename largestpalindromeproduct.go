package main
import (
	"fmt"
	"strconv"
)
var n int
func ispalindrome(n int) bool {
	str:=strconv.Itoa(n)
	length:= len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1]{
			return false
		}
	}
	return true
}
func largestpalindromeproduct(n int) (int,int,int){
	largestPalindrome:=0
	var multiplicant1, multiplicant2 int
	for i := n; i >= 10; i-- {
		for j := i; j >= 10; j-- {
			product:=i*j
			if ispalindrome(product)&&product>largestPalindrome{
				largestPalindrome=product
				multiplicant1=i
				multiplicant2=j
			}
		}
	}
	return largestPalindrome,multiplicant1,multiplicant2
}
func main(){
	fmt.Println("Enter the limits: ")
	fmt.Scan(&n)
	result,multiplicant1,multiplicant2:=largestpalindromeproduct(n)
	fmt.Println("The largest palindrome product is: ",result)
	fmt.Printf("The multiplicands are: %d and %d",multiplicant1,multiplicant2)
}