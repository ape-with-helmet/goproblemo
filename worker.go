package main

import(
	"fmt"
	"sync"
	"time"
)
func worker(id int,checkpoint chan bool,resume chan bool,wg *sync.WaitGroup){
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
	checkpoint:=make(chan bool)
	resume:=make(chan bool)
	var wg sync.WaitGroup
	for i := 1; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i,checkpoint,resume,&wg)
	}
	for i := 1; i < numWorkers; i++ {
		<-checkpoint
	}
	fmt.Println("All workers reached the checkpoint");
	fmt.Println("Resuming all workers");
	for i := 1; i < numWorkers; i++ {
		resume<-true
	}
	wg.Wait()
	fmt.Printf("All workers completed work");
}