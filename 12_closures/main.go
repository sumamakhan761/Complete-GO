package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int { // returnig func
		count++
		return count
	}
}

func main() {
	res := counter()
	fmt.Println(res())
	fmt.Println(res())
}