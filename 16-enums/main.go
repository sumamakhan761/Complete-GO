package main

import "fmt"

// type OrderStatus int
type OrderStatus string

// enumared type 
// const (
// 	Recieved OrderStatus = iota
// 	Process
// 	Ongoing
// 	Confirmed
// )

const (
	Recieved OrderStatus 	= "Recieved"
	Process  							= "Process"
	Ongoing  							= "Ongoing"
	Confirmed 						= "Confirmed"
)

func changeOrderStatus (status OrderStatus){
	fmt.Println("changing order status" , status)
}
func main()  {
	changeOrderStatus(Process)
}