package tinytools

// RemoveFromSlice returns a new slice with all occurrences of toDelete removed.
func RemoveFromSlice(target []string, toDelete string) []string {
	result := target[:0]
	for _, item := range target {
		if item != toDelete {
			result = append(result, item)
		}
	}
	return result
}

// RemoveAllFromSlice returns a new slice with all occurrences of all elements in toDelete removed.
func RemoveAllFromSlice(target, toDelete []string) []string {
	deleteMap := make(map[string]struct{})
	for _, item := range toDelete {
		deleteMap[item] = struct{}{}
	}
	result := target[:0]
	for _, item := range target {
		if _, found := deleteMap[item]; !found {
			result = append(result, item)
		}
	}
	return result
}

// UniqueSlice returns a new slice with all duplicates removed.
func UniqueSlice(target []string) []string {
	seen := make(map[string]struct{})
	result := target[:0]
	for _, item := range target {
		if _, found := seen[item]; !found {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
