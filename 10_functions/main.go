package main

// import "fmt"

func add(a int, b int) int {
	return a + b
} // or we can use like this for the same type of element
func add2(a, b int) int {
	return a + b
}

func getLanguages() (string, string, bool) {
	return "golang", "javascript", true
}

// func process(fn func(a int)int){
// 	fn(1)
// }

func process() func(a int) int {
	return func(a int) int {
		// fmt.Println(2)
		return 2
	}
}

// range for iteration
func main() {
	// res := add(3, 2)
	// res2 := add2(3, 2)

	// fmt.Println(res)
	// fmt.Println(res2)

	// lang1 , lang2 , lang3:= getLanguages()
	// fmt.Println(lang1, lang2,lang3)

	// lang1 , lang2 , _:= getLanguages()
	// fmt.Println(lang1, lang2)

	// fn := func (a int)int  {
	// 	return 2;
	// }
	// process(fn)

	fn := process()
	fn(6)
}