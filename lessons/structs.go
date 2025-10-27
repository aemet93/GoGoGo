package lessons

type Student struct {
	Name         string
	Surname      string
	Grade        int
	Group        int
	IsOnContract bool
}

func GetAdultPeople(people map[string]int) map[string]int {
	AdultPeople := make(map[string]int)

	for name, age := range people {
		if age >= 18 {
			AdultPeople[name] = age
		}
	}
	return AdultPeople
}
