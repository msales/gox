package gofp

// IndexedSet works like a map but it keeps the order.
type IndexedSet[Key comparable, Value any] struct {
	m map[Key]int
	k []Key
	v []Value
}

// Add adds a new element to the set.
func (i *IndexedSet[Key, Value]) Add(k Key, v Value) {
	if i.m == nil {
		i.m = make(map[Key]int)
	}

	if !i.Has(k) {
		i.k = append(i.k, k)
		i.v = append(i.v, v)
	}

	i.m[k] = len(i.k) - 1
}

// Get returns a the value assigned to the provided key.
func (i *IndexedSet[Key, Value]) Get(k Key) (Value, bool) {
	idx, ok := i.m[k]
	return i.v[idx], ok
}

// Has verifies if an element exists in the set.
func (i *IndexedSet[Key, Value]) Has(k Key) bool {
	_, ok := i.m[k]
	return ok
}

// Keys returns the keys in order.
func (i *IndexedSet[Key, Value]) Keys() []Key {
	return i.k
}

// Values returns the values in order.
func (i *IndexedSet[Key, Value]) Values() []Value {
	return i.v
}

// Map returns the map.
func (i *IndexedSet[Key, Value]) Map() map[Key]Value {
	m := make(map[Key]Value, len(i.m))
	for k, idx := range i.m {
		m[k] = i.v[idx]
	}

	return m
}
