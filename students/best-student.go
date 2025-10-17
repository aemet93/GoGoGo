package students

func getBestStudent(studentsMarks map[string][]float64) (string, float64) {
	var bestStudent string
	var bestMark float64

	for name, marks := range studentsMarks {
		averageMark := calculateAverageMark(marks)

		if averageMark > bestMark {
			bestStudent = name
			bestMark = averageMark
		}
	}

	return bestStudent, bestMark
}
