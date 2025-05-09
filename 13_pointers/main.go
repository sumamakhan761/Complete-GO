package main

import "fmt"

// there num copy the num not the actual num we need to give a extact memory location so that num change
// by vlaue func
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In change Num" , num)
// }

// by refrence
func changeNum(num *int) {
	*num = 5//derefrane 
	fmt.Println("In change Num" , num)
}

func main() {
	num := 1
	// changeNum(num)
	// fmt.Println("Memory address" , num)
	// fmt.Println("after changeNum in main" , num)

	fmt.Println("Memory address" , &num) // getting refrence

	changeNum(&num) // memory index

	fmt.Println("after changeNum in main" , num)
}