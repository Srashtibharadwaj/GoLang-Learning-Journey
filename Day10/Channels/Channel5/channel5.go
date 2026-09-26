package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true
	}()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}
}
