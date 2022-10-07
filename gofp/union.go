package gofp

func Union[T comparable](lists ...[]T) []T {
	occurred := make(map[T]bool, 0)
	result := make([]T, 0)

	for _, list := range lists {
		for _, element := range list {
			if !occurred[element] {
				occurred[element] = true
				result = append(result, element)
			}
		}
	}

	return result
}
