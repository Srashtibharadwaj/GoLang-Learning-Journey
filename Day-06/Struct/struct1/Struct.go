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

func main() {
	Myorder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}
	Myorder.createdAt = time.Now()
	fmt.Println("Order struct", Myorder)
}
