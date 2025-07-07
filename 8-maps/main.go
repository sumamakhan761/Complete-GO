package main

import (
	"fmt"
	"maps"
)

func main(){
		// creating map

		// m:=make(map[string]string)
		// adding an element
		// m["name"] = "golang"
		// m["area"] = "backend"

		// get element
		// IMP : if key does not exists in the map then it returns zero
		// fmt.Println(m["name"], m["area"], m["phone"]) // phone not added in the maps

		// m:=make(map[string]int)

		// m["age"] = 18
		// m["price_prod"] = 2000
		// fmt.Println(m)
		// fmt.Println(len(m))
		// delete(m , "price_prod")
		// fmt.Println(len(m))
		// fmt.Println(m)
		// clear(m)
		// fmt.Println(m)

		// m :=map[string]int{"price" : 1200 , "warranty" : 12}

		// fmt.Println(m)


	// m :=map[string]int{"price" : 1200 , "warranty" : 12}

	// v, ok := m["price"]
	// // v, ok := m["apple"] // ok checking element there or not and v give a value of the element if exist if not give zero

	// fmt.Println(v)
	// if ok { 
	// 	fmt.Println("all okh")
	// }else{
	// 	fmt.Println("not okh")
	// }
	// // etiher
	// if !ok { 
	// 	fmt.Println("not okh")
	// 	}else{
	// 	fmt.Println("all okh")
	// }

	m1 :=map[string]int{"price" : 1200 , "warranty" : 12}
	// m2 :=map[string]int{"price" : 1200 , "warranty" : 12}
	m2 :=map[string]int{"price" : 1200 , "warranty" : 11}
	
	fmt.Println(maps.Equal(m1 , m2))
		
}