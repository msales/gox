package gofp

// Intersection finds the elements that appear in both arrays and returns them (without duplicates)
func Intersection[T comparable](first, second []T, rest ...[]T) []T {
	lists := make([][]T, 0)
	lists = append(lists, first)
	lists = append(lists, second)
	lists = append(lists, rest...)

	occurrences := make(map[T]int, 0)
	result := make([]T, 0)

	for _, list := range lists {
		for _, element := range list {
			if occurrences[element] == 1 {
				result = append(result, element)
			}

			occurrences[element]++
		}
	}

	return result
}
