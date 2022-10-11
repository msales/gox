package gofp

// Union merges both arrays removing the duplicates
func Union[T comparable](first, second []T, other ...[]T) []T {
	lists := make([][]T, 0)
	lists = append(lists, first)
	lists = append(lists, second)
	lists = append(lists, other...)

	occurred := make(map[T]struct{}, 0)
	result := make([]T, 0)

	for _, list := range lists {
		for _, element := range list {
			if _, ok := occurred[element]; !ok {
				occurred[element] = struct{}{}
				result = append(result, element)
			}
		}
	}

	return result
}
