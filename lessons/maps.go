package lessons

import (
	"log"
)

func maps() {
	// variables
	// data types: int, string, bool, list (array)
	// if statement

	dogName := "Tuzik"
	birdName := "Mike"
	catName := "Garfild"

	zoo := []string{dogName, birdName, catName}

	var animal string

	animal = zoo[1]

	if animal == dogName {
		log.Printf("%v say Whoof-Whoof", animal)
	} else if animal == birdName {
		log.Printf("%v say Kar-Kar", animal)
	} else {
		log.Printf("%v say Meow-Meow", animal)
	}

	// new data type: map/dictionary/hashTable
	var fruits map[string]string

	// fruits contains fruit size
	fruits = map[string]string{
		"Apple":     "Red",
		"Banana":    "Yellow",
		"Blueberry": "Blue",
	}
	log.Println(fruits["Banana"])

	// users contains user age
	allowedUsers := map[string]bool{
		"Valentin": true,
		"Alina":    true,
		"Max":      false,
	}
	log.Println(allowedUsers["Valentin"])
}
