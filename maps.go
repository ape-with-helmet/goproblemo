package main
import "fmt"
func main(){
	myMap := make(map[string]int)
	myMap["apple"]=1
	myMap["banana"]=2
	myMap["orange"]=3

	appleValue:=myMap["apple"]
	bananaValue:=myMap["banana"]
	orangeValue:=myMap["orange"]
	fmt.Println("Value of apple: ",appleValue)
	fmt.Println("Value of banana: ",bananaValue)
	fmt.Println("Value of orange: ",orangeValue)

	myMap["apple"]=5
	fmt.Println("Updated value of apple: ",myMap["apple"])
	delete(myMap,"orange")
	fmt.Println("After deleting orange: ",myMap)

	value,exists:=myMap["banana"]
	if exists{
		fmt.Println("Value of banana: ",value)
	}else{
		fmt.Println("Banana not found in map")
	}

	for key,value:=range myMap{
		fmt.Println("Key:",key,"Value:",value)
	}
	fmt.Println("Length of the map: ",len(myMap))
}