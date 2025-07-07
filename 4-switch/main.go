package main

import (
	"fmt"
	// "time"
)

func main() {
	
		// switch
		// i:= 20
		// switch i {}
		// case 1:
		// 	fmt.Println(i)
		// case 5:
		// 	fmt.Println(i)

		// case 10:
		// 	fmt.Println(i)

		// case 15:
		// 	fmt.Println(i)
		// case 20:
		// 	fmt.Println(i)

		// default:
		// 	fmt.Println(i)
		// 

	// multiple condition switch
	
	// switch time.Now().Weekday(){
	// case time.Saturday, time.Sunday :
	// 		fmt.Println("it's weekend")
	// 		default :
	// 		fmt.Println("it's workday")
	// }

	// type switch
		whoAmI := func(i interface{}){// here means of interface type any or you right any
			switch t := i.(type){
				// return the type
			case int:
				fmt.Println("its an integer")
			case string:
				fmt.Println("its a string")
			case bool:
				fmt.Println("its a boolean")
			default:
				fmt.Println("other" , t)
		}
	}

	var whatYourName = func (string)  {
		fmt.Println("sumama")
	}
	whatYourName("sumama")
	
	whoAmI(5)
	whoAmI("golang")
	whoAmI("go")
	whoAmI(true)
	whoAmI(1.345)
}