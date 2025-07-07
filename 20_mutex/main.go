package main

import (
	"fmt"
	"sync"
)

// mutex uses to handle a rase problem
type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.views += 1
	fmt.Println(p.views)
}

func main() {
	var wg sync.WaitGroup
	mypost := post{views: 0}
	for i:=0; i < 100; i++{
		wg.Add(1)
	go mypost.inc(&wg)
}	
// lots of time request come concrently 
	wg.Wait()
	fmt.Println(mypost.views)
}