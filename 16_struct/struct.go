package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

func newOrder(id string, amount float32, status string) *order {

	//intial setup goes here...
	myOrderNew := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrderNew
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

// * (star) when used when changes status otherwise when get value doesn't use
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	//if you don't set any field,default value is zero value
	//int => 0,float => 0,string "", bool => false

	myOrder := order{
		id:     "1",
		amount: 59.423,
		status: "received",
	}

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder.status)

	fmt.Println("Order struct", myOrder)

	myOrder2 := order{
		id:        "2",
		amount:    100,
		status:    "delivered",
		createdAt: time.Now(),
	}

	myOrder2.status = "paid"

	fmt.Println("Order struct", myOrder2)

	myOrder3 := order{
		id:     "1",
		amount: 59.423,
		status: "received",
	}

	myOrder3.changeStatus("confirmed")
	fmt.Println(myOrder3)
	fmt.Println(myOrder3.getAmount())

	myOrderNew := newOrder("1", 45.54, "received")
	fmt.Println(myOrderNew.amount)

	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)
}
