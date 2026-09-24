package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type payment struct {
	gateway paymenter
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using Razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using Stripe", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

func main() {

	// Fake payment gateway for testing
	fakeGw := fakepayment{}

	newPayment := payment{
		gateway: fakeGw,
	}

	newPayment.makePayment(100)
}
