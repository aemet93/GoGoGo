package lessons

import "log"

func sum(a []int) int {
	// key word 'func'
	// func name
	// arguments - ?
	// return - ?

	var result int
	for _, value := range a {
		result += value
	}

	return result
}

func sayHello(animals map[string]string) {
	for name, value := range animals {
		log.Printf("%v says %v", name, value)
	}
}

func onlyEven(numbers []int) []int {
	var result []int
	for _, value := range numbers {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}
