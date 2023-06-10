package main
import "fmt"
func main() {
  var prev, current, sum, MS int
  fmt.Print("Enter the maximum fibonacci number:")
  fmt.Scanln(&MS)
  prev =0
  current =1
  sum=0
  for prev+current<MS{
      next := prev+current
      if next%2==0{
          sum+=next
      }
      prev=current
      current=next
  }
  fmt.Println("THe sum of even valued fibonacci numbers upto 40 are = ",sum)
}
