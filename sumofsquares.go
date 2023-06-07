package main
import "fmt"
func main() {
	fmt.Println("Enter the Maxsize")
	var n int
	fmt.Scan(&n)
	var sum = 0
	for i := 0; i <= n; i++ {
		sum = i*i+sum
	}
	fmt.Print(sum)
}
