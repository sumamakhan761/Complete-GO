package main

import "fmt"

// func printSlice[T int | string | bool](items []T){ // here we give generic type
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

func printSlice[T comparable, V string](items []T , name V){ // here we give generic type
	for _, item := range items{
		fmt.Println(item, name)
	}
}



type stack[T any]struct {
	elements []T
}

func main() {
		mystack := stack[int]{
			elements: []int{1,2,3,4},
		}
		fmt.Println(mystack)

		mystringstack := stack[string]{
			elements: []string{"golang" , "typescript" , "java"},
		}
		fmt.Println(mystack)
		fmt.Println(mystringstack)

	// nums := []int{1,2,3}
	// names := []string{"golang" , "java" , "typescript"}
	// boolean := []bool{true , false, true}
	// printSlice(nums)
	// printSlice(names)
	// printSlice(boolean)
}