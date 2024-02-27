package gofp_test

import (
	"testing"

	"github.com/msales/gox/gofp"
	"github.com/stretchr/testify/assert"
)

func TestInitializeIndexedSet(t *testing.T) {
	testCases := []struct {
		name     string
		initFunc func() gofp.IndexedSet[int, int]
	}{
		{
			name: "Initialize with empty value",
			initFunc: func() gofp.IndexedSet[int, int] {
				return gofp.IndexedSet[int, int]{}
			},
		},
		{
			name: "Initialize with var",
			initFunc: func() gofp.IndexedSet[int, int] {
				var val gofp.IndexedSet[int, int]
				return val
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.initFunc()

			assert.NotPanics(t, func() { set.Add(1, 1) })
		})
	}
}

func Test_IndexedSet_Has(t *testing.T) {
	var intSet gofp.IndexedSet[int, int]
	intSet.Add(1, 2)
	intSet.Add(2, 1)

	assert.True(t, intSet.Has(1))

	type simpleStruct struct {
		Foo int
		Bar string
	}

	var simpleStructSet gofp.IndexedSet[simpleStruct, int]
	simpleStructSet.Add(simpleStruct{Foo: 1, Bar: "foo"}, 1)
	simpleStructSet.Add(simpleStruct{Foo: 2, Bar: "baz"}, 1)

	assert.True(t, simpleStructSet.Has(simpleStruct{Foo: 1, Bar: "foo"}))
}

func Test_IndexedSet_Add(t *testing.T) {
	var intSet gofp.IndexedSet[int, int]
	// add unique elements
	intSet.Add(1, 2)

	val, ok := intSet.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 2, val)

	intSet.Add(2, 3)

	val, ok = intSet.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 3, val)
}

func Test_IndexedSet_ToSlice(t *testing.T) {
	var intSet gofp.IndexedSet[string, int]
	intSet.Add("foo", 1)
	intSet.Add("bar", 2)
	intSet.Add("baz", 3)

	intSlice := intSet.ToSlice()
	assert.Equal(t, []int{1, 2, 3}, intSlice)
}
