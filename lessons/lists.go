package lessons

import (
	"log"
)

func lists() {
	// разбираем array/list || масив/список
	var a = 1
	var b = 5
	var c = 15

	d := "Алина"
	e := "Меня"
	f := "зовут"

	// декларация масива с инициализацией
	var allInts []int = []int{a, b, c}

	// декларация масива с короткой инициализацией
	allStrings := []string{d, e, f}

	log.Println(allInts[2])

	// обращение по индексу
	log.Println(allStrings[1], allStrings[2], allStrings[0])
}
