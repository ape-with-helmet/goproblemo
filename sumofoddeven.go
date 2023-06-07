package main
import "fmt"
func main(){
	fmt.Print("Enter the max size")
	var MS int
	var odd =0
	var even = 0
	fmt.Scan(&MS)
	for i := 0; i < MS; i++ {
		if i%2==0 {
			even=even+i
		}
		if i%2!=0 {
			odd=odd+i
		}
	} 
	fmt.Print("even are = ")
	fmt.Println(even)
	fmt.Print("odd are = ")
	fmt.Println(odd)
}