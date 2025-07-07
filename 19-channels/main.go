package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// send

func processNum(numChan chan int){
	for num:=range numChan{
	fmt.Println("processing number",num)
	time.Sleep(time.Second)
	}
}

// recieve
func sum(res chan int, num1, num2 int){
		sumres := num1 +num2
		res <- sumres
}

// goroutine synchronizer
func task (done chan bool){
	defer func () {done <- true}()
	fmt.Println("processing...")
}

func emailSender(emailChan <-chan string, done chan<- bool){//emailChan <-chan for security we are use <- so that emailChan only use for recienve and here done chan<- we are only can send
	defer func () {done <- true}()
	
	for email := range emailChan{
		fmt.Println("sending email to" , email)
		time.Sleep(time.Second)
	}
}

func main() {
	// messageChan := make(chan string)
	// messageChan <- "ping"// blocking 
	// msg := <- messageChan
	// fmt.Println(msg)
	// numChan := make(chan int)
	// go processNum(numChan)


	// for {
	// 	numChan <- rand.Intn(100)
	// }// for infinite loop no need a time sleep
	// time.Sleep(time.Second * 2)


	// result := make(chan int)
	// go sum(result , 3, 5)
	// ans := <-result // blocking channel so that no need to added a time sleep
	// fmt.Println(ans)



	// done := make(chan bool)
	// go task(done)
	// <- done // block 


	//before we are using unbuffer channel

	// now starting a buffer channel
	// emailChan := make(chan string, 100)//we can add data limited 100
	// done := make(chan bool)
	// go emailSender(emailChan , done)
	// for i:=0 ; i < 5 ; i++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com" , i)
	// }

	// fmt.Println("done sending")
	// close(emailChan)// this is very imp to close a channel if not do emailchan going to block
	// <- done


	chan1 := make(chan int)
	chan2 := make(chan string)

	go func () {
		chan1 <- 10
	}()
	go func ()  {
		chan2 <- "ping"
	}()

		for i:= 0; i < 2 ; i++{
			select {
			case chan1Val := <- chan1:
				fmt.Println("received data from chan1", chan1Val)
			case chan2Val := <- chan2:
				fmt.Println("received data from chan2", chan2Val)
			}
		}	
} 