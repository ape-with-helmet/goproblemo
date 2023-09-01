package main
import(
    "fmt"
    //"math"
) 
func LPF(n int) int {
    var largest int
    for i:=2;i<=n;i++{
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
