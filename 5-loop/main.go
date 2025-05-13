package main

import "fmt"

func main() {

	// while loop
	// i := 1
	// for i <=5 {
	// 	fmt.Println(i)
	// 	i++;
	// }

	// for {// infinite loop
	// 	println("1")
	// }

	// classic for loop
	// for i := 0; i<3; i++{
	// 	fmt.Println(i)
	// }	
	

	sum := 20
	// for i := 16; i<25; i++{
	// 	if(i > sum ){
	// 		fmt.Println("greater then 20 are " , i)
	// 		}else if(i == sum && sum%2 == 0){
	// 			fmt.Println("sum is divible of 2 number is" , i)
	// 		}else{
	// 		fmt.Println("lesser then 20 number are" , i)
	// 	}
	// }

	// shortent
	for i:=range 21 {
		if(i > sum){
			if(i == 3){ 
				continue
			}
			if(i == 20){
				break;
			}
			fmt.Println("greater then 20 number are" , i)
			}else{
			fmt.Println("lesser then 20 number are" , i)
		}
	}
}