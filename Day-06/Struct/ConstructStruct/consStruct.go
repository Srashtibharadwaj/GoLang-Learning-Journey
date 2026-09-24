package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision

}

func newOrder(id string, amount float32, status string) *order {
	Myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &Myorder

}
func (o *order) changeStatus(status string) {
	o.status = status
}
func (o order) getAmount() float32 {
	return o.amount
}
func main() {
	Myorder := newOrder("1", 30.50, "received")
	fmt.Println(Myorder)

}
