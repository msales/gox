package gofp

func Intersection[T comparable](lists ...[]T) []T {
	occurrences := make(map[T]int, 0)
	result := make([]T, 0)

	for _, list := range lists {
		for _, element := range list {
			if occurrences[element] == 1 {
				result = append(result, element)
			}

			occurrences[element] += 1
		}
	}

	return result
}
