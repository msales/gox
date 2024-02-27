package gofp

// IndexedSet works like a map but it keeps the order.
type IndexedSet[Key, Value comparable] struct {
	m map[Key]Value
	i []Key
}

// Add adds a new element to the set.
func (i *IndexedSet[Key, Value]) Add(k Key, v Value) {
	if i.m == nil {
		i.m = make(map[Key]Value)
	}

	if !i.Has(k) {
		i.i = append(i.i, k)
	}

	i.m[k] = v
}

// ToSlice returns a list of values in the order they were added in.
func (i *IndexedSet[Key, Value]) ToSlice() []Value {
	ret := make([]Value, 0, len(i.Keys()))
	for _, key := range i.Keys() {
		if val, ok := i.Get(key); ok {
			ret = append(ret, val)
		}
	}

	return ret
}

// Get returns a the value assigned to the provided key.
func (i *IndexedSet[Key, Value]) Get(k Key) (Value, bool) {
	v, ok := i.m[k]
	return v, ok
}

// Has verifies if an element exists in the set.
func (i *IndexedSet[Key, Value]) Has(k Key) bool {
	_, ok := i.m[k]
	return ok
}

// Keys returns the keys in order.
func (i *IndexedSet[Key, Value]) Keys() []Key {
	return i.i
}
