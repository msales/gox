package gofp

// Union merges arrays removing the duplicates
func Union[T any](first []T, second []T, predicate func(T, T) bool) []T {
	lists := make([]T, 0)
	cleanList := make([]T, 0)

	lists = append(lists, first...)
	lists = append(lists, second...)

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
