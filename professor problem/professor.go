package main

import(
	"fmt"
	"sync"
	"time"
)
const(
	numphilosophers=3
	maxDiningCycles=2
)
type Philosopher struct{
	id int 
	leftfork,rightfork *sync.Mutex
	diningCycles int
}
type DiningTable struct{
	philosophers []*Philosopher
	waiter *sync.Mutex
}
func (p*Philosopher)think(){
	fmt.Printf("Philosopher %d is thinking\n",p.id);
	time.Sleep(time.Second)
}
func (p*Philosopher)eat(){
	p.leftfork.Lock()
	p.rightfork.Lock()
	fmt.Printf("Philosopher %d is eating\n",p.id);
	time.Sleep(time.Second)
	p.leftfork.Unlock()
	p.rightfork.Unlock()
	p.diningCycles++
}
func (p*Philosopher) dine (table*DiningTable) {
	for p.diningCycles<maxDiningCycles{
		p.think()
		table.waiter.Lock()
		p.eat()
		table.waiter.Unlock()
	}
}
func main(){
	table:=&DiningTable{
		philosophers: make([]*Philosopher, numphilosophers),
		waiter: &sync.Mutex{},
	}
	forks:=make([]*sync.Mutex,numphilosophers)
	for i := 0; i < numphilosophers; i++ {
		forks[i]=&sync.Mutex{}
	}
	for i := 0; i < numphilosophers; i++ {
		table.philosophers[i]=&Philosopher{
			id: i,
			leftfork: forks[i],
			rightfork: forks[(i+1)%numphilosophers],
		}
	}
	var wg sync.WaitGroup
	wg.Add(numphilosophers)
	for i := 0; i < numphilosophers; i++ {
		go func(p*Philosopher) {
			defer wg.Done()
			p.dine(table)
		}(table.philosophers[i])
	}
}