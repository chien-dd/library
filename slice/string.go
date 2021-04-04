package slice

func StringContain(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	// Not contain
	return false
}

func StringCompare(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	// Equal
	return true
}

func StringUnique(slice []string) []string {
	keys := make(map[string]bool)
	list := make([]string, 0)
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	// Success
	return list
}
