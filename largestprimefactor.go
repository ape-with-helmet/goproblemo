package main
import(
    "fmt"
    //"math"
) 
func LPF(n int) int {
    largest :=1
    for n%2==0{
        largest =2
        n/=2
    }
    for i:=3;i<=n;i+=2{
        for n%i==0{
            largest =i
            n/=i
        }
    }
    return largest
}
func main() {
    var cap int
    fmt.Printf("Enter the number: ")
    fmt.Scan(&cap)
    fmt.Println("The largest prime factor is: ",LPF(cap))
}
