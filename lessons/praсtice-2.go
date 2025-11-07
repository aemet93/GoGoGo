package lessons

import (
	"log"
)

func Practice() {
	var a int
	var b uint64

	c := a + int(b)
	log.Println(c)

	arr := []int{0, 1, 2, 3}
	arr2 := []string{"a", "b", "c"}

	log.Println(arr[2])

	if len(arr) > 0 {
		log.Println("OK")
	}

	arr3 := append(arr2, "d")
	log.Println(arr3)

	for _, value := range arr {
		if value%2 == 0 {
			log.Println(value)
		}
	}

	log.Println(add(arr))

	printElement(arr2)

	catWeight := make(map[string]float64)
	catWeight["Barsik"] = 7
	catWeight["Murzik"] = 5
	catWeight["Neko"] = 5.5

	printWeight(catWeight)

	log.Println(catWeight["Neko"])
}
func add(a []int) int {
	var sum1 int
	for _, value := range a {
		sum1 += value
	}
	return sum1
}
func printElement(arr2 []string) {
	for _, value := range arr2 {
		log.Println(value)
	}
}
func printWeight(catWeight map[string]float64) {
	for name, weight := range catWeight {
		log.Printf("%v is %v", name, weight)
	}
}
