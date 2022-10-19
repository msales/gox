package gofp

// Union merges arrays removing the duplicates.
func Union[T any](first []T, second []T, predicate func(T, T) bool) []T {
	lists := make([]T, len(first)+len(second))
	copy(lists, first)
	copy(lists[len(first):], second)
	cleanList := make([]T, 0)

	for _, v := range lists {
		skip := false
		for _, u := range cleanList {
			if predicate(v, u) {
				skip = true
				break
			}
		}
		if !skip {
			cleanList = append(cleanList, v)
		}
	}

	return cleanList
}
