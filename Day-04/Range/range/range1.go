package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}
	//for i := 0; i < len(nums); i++ {
	//fmt.Println(nums[i])

	//}
	//sum := 0
	//for _, num := range nums {
	//sum = sum + num
	//}
	//fmt.Println(sum)
	for i, num := range nums {
		fmt.Println(num, i)
	}
}
