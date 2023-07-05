package main
import "fmt"
func findtriplet(n int) (int,int,int){
	for a := 1; a <= n/3; a++ {
		for b := a; b <= (n-a)/2; b++ {
			c:=n-a-b
			if a*a+b*b==c*c{
				return a,b,c
			}
		}
	}
	return -1,-1,-1
}
func main(){
	var n int
	fmt.Printf("ENTER THE SUM OF THE TRIPLET\n")
	fmt.Scan(&n)
	a,b,c:=findtriplet(n)
	if a!=-1 && b!=-1 && c!=-1 {
		fmt.Printf("The special pythagorean triplet for sum %d is (%d,%d,%d)\n",n,a,b,c)
	}else{
		fmt.Printf("No special pythagorean tripet found for sum %d\n",n)
	}
}