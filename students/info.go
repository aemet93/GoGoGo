package students

import "log"

func GetStudentsInfo() {
	students := []string{"Bodya", "Alina", "Sasha", "Taran'ka"}

	marks := map[string][]float64{
		"Bodya":    {88, 55, 90, 30},
		"Alina":    {65, 46, 95, 47},
		"Sasha":    {82, 63, 51, 65},
		"Taran'ka": {100, 66, 46, 78},
	}

	printAverageMarks(marks)

	log.Print("================================")

	bestStudent, bestMark := getBestStudent(marks)

	log.Printf("%v is the best student with %v mark", bestStudent, bestMark)

	log.Print("================================")

	passedStudentsNumber := countPassedStudents(students, marks)
	log.Printf("Number of passed students: %v", passedStudentsNumber)

	// if student passed all subjects, print to console "<student name> - красавчик"
	// if not – print to console "<student name> - еблан"
}
