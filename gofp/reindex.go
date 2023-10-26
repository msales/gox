package gofp

// Reindex apply a function to all elements on an map
func Reindex[InK, OutK comparable, InV, OutV any](m map[InK]InV, fn func(InK, InV) (OutK, OutV)) (res map[OutK]OutV) {
	res = make(map[OutK]OutV, len(m))
	for k, v := range m {
		rK, rV := fn(k, v)
		res[rK] = rV
	}

	return
}
