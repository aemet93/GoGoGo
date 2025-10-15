package main

import (
	"log"
)

func vars() {
	// повторяем переменные (declaration, initialization)
	var age uint64
	var ageCategory string

	// повторяем типы (int, string, bool)
	categoryKid := "Kid"
	categoryTeen := "Teenager"
	categoryAdult := "Adult"

	age = 25
	if age < 10 {
		ageCategory = categoryKid
	} else if age >= 10 && age < 19 {
		ageCategory = categoryTeen
	} else {
		ageCategory = categoryAdult
	}

	// повторяем условие if (if statement)
	var isAdult bool
	if ageCategory == categoryAdult {
		isAdult = true
	}

	if !isAdult {
		log.Println("Ты слишком мелкий, подрастешь – приходи")
	} else {
		log.Println("Ты взросылй, все ок")
	}

	isVerified := false
	if !isVerified {
		log.Println("Имейл не верифицирован")
	}
}
