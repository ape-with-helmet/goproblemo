package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Post struct{
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}
func main(){
	url:="https://jsonplaceholder.typicode.com/posts/1"
	response.err:=http.Get(url)
	if err!=nil{
		fmt.printf("HTTTP GET request failed: %s\n",err);
		
	}
}