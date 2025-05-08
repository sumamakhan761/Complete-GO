package main

import "fmt"
func main()  {
// array is go -> start with zero index
	var nums[4]int //making a array

	var n = len(nums)
	fmt.Println(n) // arrary length  

	for i:=range(len(nums)) {
		fmt.Println(i)
	}

	for i:=range(len(nums)) {
		nums[i] = i
	}
	fmt.Println(nums)


	var vis[4]bool
	fmt.Println(vis)
	vis[2] = true
	fmt.Println(vis)

	var str[4]string
	fmt.Println(str)
	str[2] = "golang"
	fmt.Println(str)

	// manully assigning value or say to decalring it in single line
	number:=[3]int{1,2,3}
	fmt.Println(number)

	// 2d array
	// nums2d :=[2][2]int{{1,2}, {2,3}}
	// fmt.Println(nums2d)

	// - fixed size, use array
	// - memory optimazation
	//  - constant time access

}