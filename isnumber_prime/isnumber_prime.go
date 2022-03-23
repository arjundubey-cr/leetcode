package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	fmt.Println("Please enter the number")
	fmt.Scanln(&number)

	isPrime := true
	if number == 2 {
		fmt.Println("This is a prime number")
		return
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			fmt.Printf("Not a prime number")
			isPrime = false
			break
		}
	}
	if isPrime == true {
		fmt.Println("This is a prime number")
	}
}
