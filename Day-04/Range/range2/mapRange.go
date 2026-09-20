package main

import "fmt"

func main() {
	//m := map[string]string{"fname": "John", "lname": "doe"}
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	//Unicode code
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
