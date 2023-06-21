package main
import (
	"fmt"
	"strconv"
)
func ispalindrome(n int) bool {
	str:=strconv.Itoa(n)
	length:= len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length+1]{
			return false
		}
	}
}