package main

import "github.com/aemet93/GoGoGo/lessons"

func main() {
	//students.GetStudentsInfo()
	//students.PrintOnlyEven(1, 55, 79, 35, 45, 65, 56, 75)
	//lessons.Practice()

	var Robert lessons.Dog
	Robert.Name = "Robert"
	Robert.Color = "brown"

	Richard := lessons.Dog{
		Name:  "Richard",
		Color: "grey",
	}

	Richard.Run()
	Robert.Run()

	Richard.SayName()
	Robert.SayName()

	Dogs := []lessons.Dog{Robert, Richard}

	SayDogsName(Dogs)
}

func SayDogsName(dogs []lessons.Dog) {
	for _, dog := range dogs {
		dog.SayName()
	}
}
