package students

import "log"

func printAverageMarks(studentsMarks map[string][]float64) {
	for name, marks := range studentsMarks {
		averageMark := calculateAverageMark(marks)
		log.Printf("%v has %v average mark", name, averageMark)
	}
}

func calculateAverageMark(marks []float64) float64 {
	var sum float64
	var count int

	for _, mark := range marks {
		sum = sum + mark // long adding
		count += 1       // short adding
	}

	return sum / float64(count)
}
