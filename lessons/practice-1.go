package lessons

import "log"

type Car struct {
	Model string
	Color string
}

func (c Car) run() {
	// start engine
}
func main() {
	var name string
	var age int
	var people map[string]int

	name = "Ivan"
	age = 15
	people = map[string]int{
		name: age,
	}

	people["Vasya"] = 22

	log.Println(people[name])

	sumAge(age, people["Vasya"])

	totalAge := sumAge(age, people["Vasya"])
	log.Println(totalAge)

	list := []int{1, 2, 3, 4, 5}
	for _, value := range list {
		log.Println(value)

	}

	for n, a := range people {
		log.Print(n, a)
	}
	if people["Vasya"] >= 18 {
		log.Println("Adult")
	}

	car := Car{
		Model: "Audi",
		Color: "White",
	}
	car.run()
}
func sumAge(a, b int) int {
	return a + b
}
