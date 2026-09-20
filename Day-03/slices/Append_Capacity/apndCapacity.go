package main

import "fmt"

func main() {
	var nums = make([]int, 2, 5)
	//fmt.Println(cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 2)
	fmt.Println(nums)
	fmt.Println(cap(nums))
}
