package main

import "fmt"

func main() {
	var nums = make([]int, 0, 3)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Println(cap(nums))

}
