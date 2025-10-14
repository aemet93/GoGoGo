package main

import "log"

// LOOPS (циклы)
func main() {
	fruits := []string{"apple", "banana", "orange"}

	log.Println(fruits)
	log.Println("===================================")

	// for loop
	for index, value := range fruits {
		log.Printf("Friut %v has index %v", value, index)
	}
	log.Println("===================================")

	prices := []float64{5.15, 10.8, 3}
	var total float64
	for _, price := range prices {
		total = total + price
	}
	log.Print(total)

	fruitPrice := map[string]float64{
		"apple":  5.15,
		"banana": 10.8,
		"orange": 3,
	}

	log.Println("===================================")
	log.Print("Menu:")
	for key, price := range fruitPrice {
		log.Printf("- %v %v uah", key, price)
	}
}
