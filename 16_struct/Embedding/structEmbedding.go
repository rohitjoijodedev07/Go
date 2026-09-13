package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
	customer
}

// func newOrder(id string, amount float32, status string) *order {

// 	//intial setup goes here...
// 	myOrderNew := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrderNew
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// // * (star) when used when changes status otherwise when get value doesn't use
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {

	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "9876543210",
	// }

	// newOrder := order{
	// 	id:     "1",
	// 	amount: 30,
	// 	status: "received",
	// }

	// fmt.Println(newOrder.customer)

	// newOrder := order{
	// 	id:       "1",
	// 	amount:   30,
	// 	status:   "received",
	// 	customer: newCustomer,
	// }

	// fmt.Println(newOrder.customer)

	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890"},
	}

	fmt.Println(newOrder.customer)

}
