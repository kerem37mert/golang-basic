package main

import "fmt"

func isPrime(number int) bool {

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var number int
	var result bool

	fmt.Println("Enter number...")
	fmt.Scan(&number)

	result = isPrime(number)

	if result {
		fmt.Println(number, "is a prime number")
	} else {
		fmt.Println(number, "is not a prime number")
	}
}
