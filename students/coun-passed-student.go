package students

func countPassedStudents(students []string, marks map[string][]float64) uint {
	var count uint

	for _, name := range students {

		isPassed := isAllSubjectsPassed(marks[name])

		if isPassed {
			count += 1
		}

	}

	return count
}
