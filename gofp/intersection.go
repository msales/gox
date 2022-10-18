package gofp

// Intersection finds the elements that appear in arrays and returns them (without duplicates). You need to put at least two slices.
func Intersection[T any](first []T, second []T, predicate func(T, T) bool) []T {
	repeatedOccurences := make([]T, 0)
	cleanList := make([]T, 0)

	for _, elementFromFirst := range first {
		for _, elementFromSecond := range second {
			if predicate(elementFromFirst, elementFromSecond) {
				{
					repeatedOccurences = append(repeatedOccurences, elementFromFirst)
				}
			}
		}
	}

	for _, v := range repeatedOccurences {
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
