package main

import "fmt"

func main() {
	var nums []int
	//uninitialised slice is nil
	fmt.Println(nums)
	fmt.Println(nums == nil)

}
