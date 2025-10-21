package students

func calculateGroupAverage(marks map[string][]float64) float64 {
	var sum float64
	var count float64

	for _, studentMarks := range marks {

		for _, studentMark := range studentMarks {
			sum += studentMark
			count += 1
		}

	}

	return sum / count
}
