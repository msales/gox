package gofp

// Reduce applies passed function to each value of the list allowing to accumulate the T values inside to V.
//
// Example:
// intSlice := int {1, 2, 3, 4, 5}
// resp := gofp.Reduce(intSlice, func(accum, item int) int { return accum + item }, 0)
// fmt.Println(resp) // ---> 15
func Reduce[T, V any](list []T, fn func(V, T) V, initValue V) V {
	acc := initValue
	for _, v := range list {
		acc = fn(acc, v)
	}
	return acc
}
