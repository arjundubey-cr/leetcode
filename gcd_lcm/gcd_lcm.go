package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	on1 := num1
	on2 := num2
	var lcm, gcd int
	for true {
		rem := num1 % num2
		if rem == 0 {
			break
		}
		num1 = num2
		num2 = rem

	}
	gcd = num2
	lcm = (on1 * on2) / gcd

	fmt.Printf("GCD=%d\nLCM=%d\n", gcd, lcm)
}
