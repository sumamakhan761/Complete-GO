package main

import (
	"fmt"
	"sync"
)

func task(id int , w *sync.WaitGroup) {//using a pointer
	defer w.Done() // cleaning function
	fmt.Println("doing task" , id)
}

func main() {
	// creting wait group
	var wg sync.WaitGroup
	for i:=0; i<=10; i++{
		wg.Add(1)
		go task(i , &wg) //multiple threads working same time

		// go func(i int){//closures
		// 	fmt.Println(i)
		// }(i)
	}
	wg.Wait()
}