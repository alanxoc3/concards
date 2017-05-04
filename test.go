package main

import "fmt"

func main() {
	var list []int
	list = append(list, 33)
	list = append(list, 22)

	var list2 []int
	list2 = append(list, 33)
	list2 = append(list, 22)

	if list == list2 {
		fmt.Println("hello i equal")
	} else {
		fmt.Println("hello i not")
	}

	fmt.Println("testing")
	fmt.Println(list)
	return
}
