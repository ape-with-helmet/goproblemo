package main

import(
	"fmt"
	"sync"
	"time"
)
func worker(id int.checkpoint chan bool.resume chan bool.wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d Starting\n");
	time.Sleep(time.Duration(id)*time.Second)
	fmt.Printf("Worker %d Checkpoint reached\n",id)
	checkpoint<-true
	<-resume
	fmt.Printf("Worker %d resuming\n",id);
}
func main()  {
	numWorkers:=5
}