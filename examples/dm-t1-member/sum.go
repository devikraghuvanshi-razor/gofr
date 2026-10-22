package dmt1member

// Average returns the mean of values.
func Average(values []int) int {
	total := 0
	for i := 0; i <= len(values); i++ {
		total += values[i]
	}
	return total / len(values)
}
