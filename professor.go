package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosophers = 3
	maxDiningCycles = 2
)

type Philosopher struct {
	id                  int
	leftFork, rightFork *sync.Mutex
	diningCycles        int
}
type DiningTable struct {
	philososphers []*Philosopher
	waiter        *sync.Mutex
}

func (p *Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking\n", p.id)
	time.Sleep(time.Second)
}
func (p *Philosopher) eat() {
	p.leftFork.Lock()
	p.rightFork.Lock()
	fmt.Printf("Philosopher %d is eating\n", p.id)
	time.Sleep(time.Second)
	p.rightFork.Unlock()
	p.leftFork.Unlock()
	p.diningCycles++
}
func (p *Philosopher) dine(table *DiningTable) {
	for p.diningCycles < maxDiningCycles {
		p.think()
		table.waiter.Lock()
		p.eat()
		table.waiter.Unlock()
	}
}
func main() {
	table := &DiningTable{
		philososphers: make([]*Philosopher, numPhilosophers),
		waiter:        &sync.Mutex{},
	}
	forks := make([]*sync.Mutex, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}
	for i := 0; i < numPhilosophers; i++ {
		table.philososphers[i] = &Philosopher{
			id:        i,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%numPhilosophers],
		}
	}
	var wg sync.WaitGroup
	wg.Add(numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		go func(p *Philosopher) {
			defer wg.Done()
			p.dine(table)
		}(table.philososphers[i])
	}
	wg.Wait()
}