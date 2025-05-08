package main

import "fmt"

// range for iteration
func main(){
	// nums := []int{6,7,8}
	// for i:=0 ; i < len(nums); i++{
	// 	fmt.Print(nums[i])
	// 	fmt.Print(" ")
	// }
	// fmt.Println()

	// for i := range(len(nums)){
	// 	fmt.Print(nums[i])
	// 	fmt.Print(" ")
	// }
	// fmt.Println()

	// sum := 0
	// for _, num:= range nums {
	// 	fmt.Print(num)
	// 	sum += num
	// 	fmt.Print(" ")
	// }
	// fmt.Println()
	// fmt.Println(sum)

	// for i, num:= range nums { // here i is index
	// 	fmt.Println(i ,num)
	// }

	// m := map[string]string{"hello" : "assala wlikum" , "how are you":"aap kaise hain"}
	
	// for k , v :=range m {
	// 	fmt.Println(k , v)
	// }
	// for k :=range m {// for giving single argument, only come key
	// 	fmt.Println(k)
	// }

	// unicode code for all the alpahbet -> point rune
	// first arg -> starting byte of rune
	// 255 -> come under (1 byte) , for more order charc may be more then 255 so , now come under ( 2 bytes) 
	for i , c := range "golang"{
		// fmt.Println(i , c) // rune come
		fmt.Println(i , string(c))
	}
}