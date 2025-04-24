package main

import "fmt"

func main() {
	sum := 24
	if sum > 20 {
		if sum >= 22 {
			fmt.Println("you are greater then 24")
		}
		fmt.Println("you are greater then 20")
	}else{
		fmt.Println("you are less then 20")
	}

	// drive_car := true

	// we can declare variable inside if construnct a if-else fairly normal
	// if age := 45 ; age > 18 && drive_car {
	// 	fmt.Println("you can drive a card becauses you age is" , age)
	// }else{
	// 	fmt.Println("you can't drive a card becauses you age is" , age)
	// }


	// go dost not have ternary operater , you will have to use normal if else



}