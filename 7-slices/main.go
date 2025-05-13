package main

import (
	"fmt"
	"slices"
)

func main(){
	// - not fixed size, use slice like what 
	// - size dynamical
	// most user construct in go
	//  useful methods on the arrays
	// we do in the java -> in java we use arraylist 
	// var nums []int //uninitialized slice is nil (nil means null)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int , 0, 2) // here we added a size in second argument third argument is the intial capacity  
	// fmt.Println(nums)
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums)) // max numbers of elements can fit
	// nums = append(nums , 1)
	// nums = append(nums , 2)
	// nums = append(nums , 3)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))// now you see when you run now capacity become a double of len(nums)

	// another way
	// nums1 := []int{} 
	// fmt.Println(nums1)
	// fmt.Println(nums1 == nil)
	// fmt.Println(len(nums1))
	// fmt.Println(cap(nums1))
	// nums1 = append(nums1 , 1)
	// nums1 = append(nums1 , 2)
	// nums1 = append(nums1 , 3)
	// nums1[0] = 100;
	// fmt.Println(nums1)
	// fmt.Println(cap(nums1))

	// copy fucntion

	// var nums2 = make([]int , 0, 5)
	// nums2 = append(nums2 , 1)
	// nums2 = append(nums2 , 2)
	// var nums3 = make([]int ,len(nums2))
	// fmt.Println(nums2)
	// fmt.Println(nums3)
	// copy(nums3 , nums2)
	// fmt.Println(nums3)

	// slice operator
	var nums4 = []int{1,2,3}
	fmt.Println(nums4[0:2])// print element from 0 to 2 but not included 2
	fmt.Println(nums4[0:3])
	fmt.Println(nums4[:1])
	fmt.Println(nums4[1:])


	// slice package
	var nums1 = []int{1,2}
	var nums2 = []int{1,2}
	fmt.Println(slices.Equal(nums1, nums2))// means num1 and num2 equal -> true

}