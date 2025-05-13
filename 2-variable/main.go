package main

import "fmt"

func main() {
	// fmt.Println(1);
	// fmt.Println(1+1);
	// fmt.Println("1" + "1");
	// fmt.Println(1*1);// mul
	// fmt.Println(10.5);// floats
	// fmt.Println(true);// boolean

	// variable 
	// var name = "sumama"
	// var secondname string = "sumama"
	// fmt.Println(name)
	// fmt.Println(secondname)

	// var isAdult = true
	// fmt.Println(isAdult)

	// shorthand syntax
	// name:="sumama"
	// fmt.Println(name)

	// when you not declate so you need to declare

	// var name string
	// name = "golang"
	// name2:= "Go"

	// fmt.Println(name2)
	// fmt.Println(name)


	// constant
	// const c = 2;
	// // c = 4 // we can't assign to value bcz contant value is contant 
	// fmt.Println(c)

	const(
		host = "localhost"
		port = 5000
	)

	fmt.Println(host,port)

	const(
		name = "sumama"
		age = "18"
		language = "Go"
	)

	fmt.Println("name -> " + name, "age -> " + age ,"language -> " + language)
}