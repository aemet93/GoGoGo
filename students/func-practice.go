package students

import "log"

func Greetings(name string) {
	log.Printf("Hello, my name is %v", name)
}

func SumIntsList(list ...int) int {
	var sum int
	for _, value := range list {
		sum += value
	}

	return sum
}

func PrintOnlyEven(list ...int) {
	for _, value := range list {
		if value%2 == 0 {
			log.Print(value)
		}
	}
}
