package main

import (
	"fmt"
	"math"
)

func main() {
	var low, high int
	fmt.Scanf("%d %d", &low, &high)

	i := low
	for i <= high {
		if checkPrime(i) {
			fmt.Println(i)
		}
		i++
	}
}

func checkPrime(number int) bool {
	if number == 1 {
		return false
	} else if number == 2 {
		return true
	} else {
		for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
			if number%i == 0 {
				return false
			}
		}
	}
	return true

}
