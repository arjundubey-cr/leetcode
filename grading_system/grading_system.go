package main

import "fmt"

func main() {
	var marks int
	fmt.Println("Please enter your marks")
	fmt.Scanln(&marks)

	if marks >= 90 {
		fmt.Println("Excellent")
	} else if marks >= 80 {
		fmt.Println("Good")
	} else if marks >= 70 {
		fmt.Println("Fair")
	} else if marks >= 60 {
		fmt.Println("Meets expectation")
	} else if marks <= 60 {
		fmt.Println("Below par")
	}

}
