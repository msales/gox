package gofp_test

import (
	"fmt"
	"testing"

	"github.com/msales/gox/gofp"
	"github.com/stretchr/testify/assert"
)

func TestIntersection_Int(t *testing.T) {
	slice1 := []int{1, 3, 5}
	slice2 := []int{2, 5, 7}
	predicate := func(elementOne, elementTwo int) bool {
		return elementOne == elementTwo
	}

	got := gofp.Intersection(slice1, slice2, predicate)
	want := []int{5}
	assert.Len(t, got, len(want), fmt.Sprintf("Got %+v, want %+v", got, want))
	assert.Equal(t, want, got, fmt.Sprintf("Got %+v, want %+v", got, want))
}

func TestIntersection_String(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{"c", "d", "e"}
	predicate := func(elementOne, elementTwo string) bool {
		return elementOne == elementTwo
	}

	got := gofp.Intersection(slice1, slice2, predicate)
	want := []string{"c"}
	assert.Len(t, got, len(want), fmt.Sprintf("Got %+v, want %+v", got, want))
	assert.Equal(t, want, got, fmt.Sprintf("Got %+v, want %+v", got, want))
}

func TestIntersection_ComplexStruct(t *testing.T) {
	type person struct {
		name    string
		age     int
		country []string
	}

	slice1 := []person{
		{name: "aaa", age: 33, country: []string{"PL", "CU"}},
		{name: "bbb", age: 32, country: []string{"PL"}},
		{name: "cc", age: 31, country: []string{"DE"}},
	}

	slice2 := []person{
		{name: "ggg", age: 45, country: []string{"GB"}},
		{name: "ff", age: 31, country: []string{"FIN"}},
		{name: "cc", age: 31, country: []string{"DE"}},
	}

	predicate := func(elementOne, elementTwo person) bool {
		if len(elementOne.country) != len(elementTwo.country) {
			return false
		}

		areCountryEquals := true
		for _, val := range elementOne.country {
			if ok := gofp.Any(elementTwo.country, func(item string) bool { return item == val }); !ok {
				areCountryEquals = false
				break
			}
		}
		return elementOne.age == elementTwo.age &&
			areCountryEquals
	}

	got := gofp.Intersection(slice1, slice2, predicate)
	want := []person{{name: "cc", age: 31, country: []string{"DE"}}}

	assert.Len(t, got, len(want), fmt.Sprintf("Got %+v, want %+v", got, want))
	assert.Equal(t, want, got, fmt.Sprintf("Got %+v, want %+v", got, want))
}

func BenchmarkIntersection_TwoSlices(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 6, 7, 8, 9}
	predicate := func(elementOne, elementTwo int) bool {
		return elementOne == elementTwo
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.Intersection(slice1, slice2, predicate)
	}

	b.ReportAllocs()
}
