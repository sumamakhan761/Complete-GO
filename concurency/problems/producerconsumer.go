package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	buffersize = 5  // ma num of order comes under the queue(consumer)
	numCustomers = 10 // total number of customers(producer)
)

type Order struct {
	id    int
	items int
}

func customers(id int, orderQueue chan Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ { // Each customer places 3 orders
		order := Order{
			id:    id*100 + i,
			items: rand.Intn(5) + 1,
		}
		fmt.Printf("🧍 Customer %d placing order %d with %d items\n", id, order.id, order.items)
		orderQueue <- order // Will block if queue is full
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func warehouse(orderQueue chan Order, done chan bool) {
	for order := range orderQueue {
		fmt.Printf("processing order %d (%d items)\n", order.id, order.items)
		time.Sleep(time.Millisecond * time.Duration(order.items*200))
	}
	done <- true
}

func main() {
	rand.Seed(time.Now().UnixNano()); // ensure diffrent random number on each run

	orderQueue := make(chan Order , buffersize)

	var wg sync.WaitGroup

	done := make(chan bool)
	go warehouse(orderQueue, done)

	for i:= 1; i <= numCustomers; i++{
		wg.Add(1)
		go customers(i, orderQueue, &wg)
	}

	wg.Wait()
	close(orderQueue)

	<-done
	fmt.Println("all orders process , system shutting down")

}