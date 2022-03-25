package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	for div := 2; div*div <= number; div++ {
		for number%div == 0 {
			number = number / div
			fmt.Println(div)
		}
	}
	if number != 1 {
		fmt.Println(number)
	}
}
