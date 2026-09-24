package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision

}

func (o *order) changeStatus(status string) {
	o.status = status

}
func main() {
	Myorder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	Myorder.changeStatus("confirmed")
	fmt.Println(Myorder)
}
