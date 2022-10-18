package gofp_test

import (
	"sort"
	"testing"

	"github.com/msales/gox/gofp"
)

func TestUnion_Comparable_Comparable_Int_TwoArrays(t *testing.T) {
	slices := [][]int{
		{1, 3, 5},
		{2, 5, 7},
	}

	predicate := func(elementOne, elementTwo int) bool {
		return elementOne == elementTwo
	}

	got := gofp.Union(slices[0], slices[1], predicate)

	want := []int{1, 2, 3, 5, 7}

	if len(want) != len(got) {
		t.Errorf("Got %+v, want %+v", got, want)
		return
	}

	sort.Ints(got)

	for idx, value := range want {
		if value != got[idx] {
			t.Errorf("Got %+v, want %+v", got, want)
			break
		}
	}
}

func TestUnion_Comparable_Comparable_String_TwoArrays(t *testing.T) {
	slices := [][]string{
		{"1", "3", "5"},
		{"2", "5", "7"},
	}

	predicate := func(elementOne, elementTwo string) bool {
		return elementOne == elementTwo
	}

	got := gofp.Union(slices[0], slices[1], predicate)

	want := []string{"1", "2", "3", "5", "7"}

	if len(want) != len(got) {
		t.Errorf("Got %+v, want %+v", got, want)
		return
	}

	sort.Strings(got)

	for idx, value := range want {
		if value != got[idx] {
			t.Errorf("Got %+v, want %+v", got, want)
			break
		}
	}
}

func TestUnion_SomeType_TwoArrays(t *testing.T) {
	type person struct {
		name    string
		age     int
		country string
	}

	slices := [][]person{
		{
			{
				name:    "aaa",
				age:     33,
				country: "PL",
			},
			{
				name:    "bbb",
				age:     32,
				country: "PL",
			},
			{
				name:    "cc",
				age:     31,
				country: "DE",
			},
		},
		{
			{
				name:    "ddd",
				age:     45,
				country: "GB",
			},
			{
				name:    "eee",
				age:     31,
				country: "FIN",
			},
			{
				name:    "fff",
				age:     31,
				country: "DE",
			},
		},
	}

	predicate := func(elementOne, elementTwo person) bool {
		return elementOne.age == elementTwo.age && elementOne.country == elementTwo.country
	}

	got := gofp.Union(slices[0], slices[1], predicate)

	sort.Slice(got, func(i, j int) bool {
		return got[i].name < got[j].name
	})

	want := []person{
		{name: "aaa", age: 33, country: "PL"},
		{name: "bbb", age: 32, country: "PL"},
		{name: "cc", age: 31, country: "DE"},
		{name: "ddd", age: 45, country: "GB"},
		{name: "eee", age: 31, country: "FIN"},
	}

	if len(want) != len(got) {
		t.Errorf("Got %+v, want %+v", got, want)
		return
	}

	for k, v := range want {
		if v != got[k] {
			t.Errorf("Got %+v, want %+v", got, want)
		}
	}
}

func BenchmarkUnion_TwoSlices(b *testing.B) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 6, 7, 8, 9}
	predicate := func(elementOne, elementTwo int) bool {
		return elementOne == elementTwo
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gofp.Union(slice1, slice2, predicate)
	}

	b.ReportAllocs()
}
