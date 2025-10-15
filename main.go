package main

import "log"

func main() {
	animals := map[string]string{
		"Cat":  "Meow",
		"Dog":  "Whoof",
		"Bird": "Kar",
	}

	sayHello(animals)
	log.Print()

	numbers := []int{2, 5, 10, 9, 7, 6, 11}

	res := onlyEven(numbers)
	log.Print(res)
}
