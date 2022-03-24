package main

import "fmt"

func main() {
	var limit int
	fmt.Scanln(&limit)
	first, second := 0, 1
	fmt.Println(first)
	for i := 0; i < limit; i++ {
		temp := second
		second = first + second
		first = temp
		fmt.Println(first)
	}
}
