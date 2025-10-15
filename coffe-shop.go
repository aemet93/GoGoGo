package main

import (
	"log"
)

func main() {
	// DRINKS LIST
	drinks := []string{"Latte", "Capuchino", "Espresso"}

	// DRINKS MAP
	menu := map[string]int{
		drinks[0]: 75,
		drinks[1]: 100,
		drinks[2]: 65,
	}

	// CUP SIZE MAO
	cupSize := map[string]int{
		"Standart": 0,
		"Medium":   5,
		"Large":    10,
	}

	// MENU PRINT
	log.Println("Menu:")
	// - Latte: 75 грн
	log.Printf("- %v: %v uah\n", drinks[0], menu[drinks[0]])
	log.Printf("- %v: %v uah\n", drinks[1], menu[drinks[1]])
	log.Printf("- %v: %v uah\n", drinks[2], menu[drinks[2]])

	// USER PARAMS
	var totalPrice int
	balance := 15
	hasDiscount := true
	selectedDrink := "Espresso"
	selectedSize := "Large"

	// SELL ACTION
	totalPrice = menu[selectedDrink]
	sizePrice := cupSize[selectedSize]

	// long math operation (+)
	totalPrice = totalPrice + sizePrice

	if hasDiscount {
		// short math operation (-=)
		totalPrice -= 10
	}

	// CALCULATIONS
	log.Println("Total Price:", totalPrice)

	// BALANCE CHECK
	if totalPrice > balance {
		log.Println("You have no enogh money. You need extra", totalPrice-balance)
	} else {
		balance -= totalPrice
		log.Println("Tanks for buying. Money left:", balance)
	}
}
