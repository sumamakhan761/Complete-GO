package main

import "fmt"

func sum(nums ...int)int {
	total := 0
	for _,num:= range nums {
		total += num
	}
	return total
}

func printAnything(nums ...interface{}) {
	fmt.Println(nums...)
}

func main() {
	// fmt.Println(1,2,3,4,5, "hello") // variadic funtion are the func where we give a N number of argument

//  res :=sum(1,2,3,4,5,6,7,8,9) // this is called varadic function
//  fmt.Println(res)
// num := []int{1,2,3,4,5} 
//  res :=sum(num...) // this is called varadic function
//  fmt.Println(res)

  printAnything(1,2,3,4,6,"hii", true, 2.35)
}