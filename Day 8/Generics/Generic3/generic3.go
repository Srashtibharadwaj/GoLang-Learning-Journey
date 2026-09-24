package main

import "fmt"

func printSlice[T comparable](items []T, name string) {
	fmt.Println("Name:", name)

	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	vals := []bool{true, false, true}

	printSlice(vals, "John")
}
