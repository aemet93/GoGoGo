package students

func isAllSubjectsPassed(studentsMarks []float64) bool {
	for _, mark := range studentsMarks {
		if mark < float64(60) {
			return false
		}
	}

	return true
}
