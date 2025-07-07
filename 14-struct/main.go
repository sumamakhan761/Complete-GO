package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer // here added a customer struct
}

// receiver type
func(o *order) changeStatus(status string){
	o.status = status // struct do derefrencing not needed again give a star * 
}

// when want to modify somthing so need to use a pointer( * ) on struct but want to return or get somthing not needed a pointer

func main(){
	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
// newcustomer:= customer{
// 			name: "sumama",
// 			phone: "789979",
// 		}
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 		customer:newcustomer,
// 	}


	// another way
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
		customer:customer{
			name: "sumama",
			phone: "789979",
		},
	}

	// if we not added a any feild require if not give it will not give any error give a default wvalue is zero value

	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)
	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)

	// myOrder2 := order {
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// }

	// myOrder.status = "paid"
	// fmt.Println("Order struct", myOrder)
	// fmt.Println("Order struct", myOrder2)
	
	// when want only use one time
	// language := struct {
	// 	name string
	// 	isGood bool
	// }{"golang" , true}

	// fmt.Println(language)


	// struct embedding
	fmt.Println(myOrder)
	fmt.Println(myOrder.customer.name)
	fmt.Println(myOrder.customer.phone)
	myOrder.customer.name = "usama"
	fmt.Println(myOrder.customer.name)
}