package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updatemsg(s string, mutex *sync.Mutex){
	defer wg.Done()
	mutex.Lock()
	msg = s
	mutex.Unlock()
}
func main (){
	msg = " Helloe, world!"
	var mutex sync.Mutex

	wg.Add(2);
	updatemsg("hello", &mutex)
	updatemsg("world", &mutex)

	wg.Wait();
	fmt.Println(msg)
}