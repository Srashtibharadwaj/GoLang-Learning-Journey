package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["name"] = "Srashti"
	m["location"] = "Hamirpur"
	fmt.Println(m["name"])
	fmt.Println(m["location"])
	fmt.Println(m["area"])
}
