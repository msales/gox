package gofp_test

import (
	"fmt"
	"testing"

	"github.com/msales/gox/gofp"
	"github.com/stretchr/testify/assert"
)

func TestInitializeSet(t *testing.T) {
	testCases := []struct {
		name     string
		initFunc func() gofp.Set[int]
		withErr  bool
	}{
		{
			name: "Initialize with 'make'",
			initFunc: func() gofp.Set[int] {
				return make(gofp.Set[int])
			},
		},
		{
			name: "Initialize with empty value",
			initFunc: func() gofp.Set[int] {
				return gofp.Set[int]{}
			},
		},
		{
			name: "Initialize with var",
			initFunc: func() gofp.Set[int] {
				var val gofp.Set[int]
				return val
			},
			withErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := tc.initFunc()
			if tc.withErr {
				assert.Panics(t, func() { set.Add(1) })
				return
			}

			assert.NotPanics(t, func() { set.Add(1) })
		})
	}
}

func TestToString(t *testing.T) {
	intSet := gofp.Set[int]{1: struct{}{}, 2: struct{}{}}
	strVal := fmt.Sprint(intSet)
	// Using Contains because the iteration over golang maps is not sortered.
	assert.Contains(t, []string{"[2 1]", "[1 2]"}, strVal)

	type simpleStruct struct {
		Foo int
		Bar string
	}
	simpleStructSet := gofp.Set[simpleStruct]{{Foo: 1, Bar: "abc"}: struct{}{}, {Foo: 35, Bar: "xyz"}: struct{}{}}
	// Using Contains because the iteration over golang maps is not sortered.
	strVal = fmt.Sprint(simpleStructSet)
	assert.Contains(t, []string{"[{1 abc} {35 xyz}]", "[{35 xyz} {1 abc}]"}, strVal)
}

func TestHas(t *testing.T) {
	intSet := gofp.Set[int]{1: struct{}{}, 2: struct{}{}}
	assert.True(t, intSet.Has(1))

	type simpleStruct struct {
		Foo int
		Bar string
	}
	simpleStructSet := gofp.Set[simpleStruct]{{Foo: 1, Bar: "abc"}: struct{}{}, {Foo: 35, Bar: "xyz"}: struct{}{}}
	// Using Contains because the iteration over golang maps is not sortered.
	assert.True(t, simpleStructSet.Has(simpleStruct{Foo: 1, Bar: "abc"}))
}

func TestAdd(t *testing.T) {
	intSet := make(gofp.Set[int])
	// add unique elements
	intSet.Add(1, 2, 3, 4)
	assert.Len(t, intSet, 4)
	// add already existing elements plus new elements
	intSet.Add(2, 3, 24)
	assert.Len(t, intSet, 5)
}

func TestRemove(t *testing.T) {
	intSet := gofp.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}}
	// remove existing element
	intSet.Remove(3)
	assert.Len(t, intSet, 3)

	// remove no-existing element
	intSet.Remove(3423)
	assert.Len(t, intSet, 3)
}

func TestToSlice(t *testing.T) {
	intSet := gofp.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}}
	intSlice := intSet.ToSlice()
	assert.Len(t, intSlice, len(intSet)) // verify that both length are the same.

	for _, v := range intSlice {
		assert.True(t, intSet.Has(v))
	}
}

func TestToSliceWithSort(t *testing.T) {
	intSet := gofp.Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}}
	expectedSlice := []int{1, 2, 3, 4}
	intSlice := intSet.ToSliceWithSort(func(a, b int) bool { return a < b })
	assert.Len(t, intSlice, len(intSet)) // verify that both length are the same.

	for i := 0; i < len(expectedSlice); i++ {
		assert.Equal(t, expectedSlice[i], intSlice[i])
	}
}
