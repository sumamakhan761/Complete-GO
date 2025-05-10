package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task" , id)
}

func main() {
	for i:=0; i<=10; i++{
		// go task(i) //multiple threads working same time

		go func(i int){//closures
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}