package main
import "fmt"
func main() {
	fmt.Println("Enter a few numbers")
	var n int
	fmt.Scan(&n)
	var sum = 0
	var a = 0
	fmt.Print("now enter n numbers")
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		sum = a + sum
	}
	fmt.Print(sum)
}
