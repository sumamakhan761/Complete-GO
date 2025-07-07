package main

import (
	"fmt"
	"sync"
)

func printSomething(s string , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := [] string {
		"alpha",
		"beta",
		"gamma",
		"theta",
	}

	wg.Add(len(words));

	for i , val := range words {
		go printSomething(fmt.Sprintf("%d: %s", i , val), &wg)
	}
	wg.Wait()

}