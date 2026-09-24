package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", num)
}
func main() {
	num := 1
	changeNum(&num)
	//for Memory address use "&"
	fmt.Println("After changeNum in main", num)
}
