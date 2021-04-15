package slice

func IntegerContain(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	// Not contain
	return false
}
