package main

import "fmt"

func factorial(number int) int {
	if number == 1 || number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}

func main() {
	var number, result int

	fmt.Println("Enter number...")
	fmt.Scan(&number)

	result = factorial(number)

	fmt.Println(result)

}
