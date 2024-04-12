package gofp

// IndexedSet works like a map but it keeps the order.
type IndexedSet[Key comparable, Value any] struct {
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

// Map returns the map.
func (i *IndexedSet[Key, Value]) Map() map[Key]Value {
	return i.m
}
