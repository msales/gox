package gofp_test

import (
	"testing"

	"github.com/msales/gox/gofp"
)

func TestIntersection_Comparable_Int_TwoArrays(t *testing.T) {
	slices := [][]int{
		{1, 3, 5},
		{2, 5, 7},
	}

	predicate := func(elementOne, elementTwo int) bool {
		return elementOne == elementTwo
	}

	got := gofp.Intersection(slices[0], slices[1], predicate)

	want := []int{5}

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

func TestIntersection_Comparable_String_TwoArrays(t *testing.T) {
	slices := [][]string{
		{"a", "b", "c"},
		{"c", "d", "e"},
	}

	predicate := func(elementOne, elementTwo string) bool {
		return elementOne == elementTwo
	}

	got := gofp.Intersection(slices[0], slices[1], predicate)

	want := []string{"c"}

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

func TestIntersection_SomeType_TwoArrays(t *testing.T) {
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
				name:    "ggg",
				age:     45,
				country: "GB",
			},
			{
				name:    "ff",
				age:     31,
				country: "FIN",
			},
			{
				name:    "cc",
				age:     31,
				country: "DE",
			},
		},
	}

	predicate := func(elementOne, elementTwo person) bool {
		return elementOne.age == elementTwo.age && elementOne.country == elementTwo.country
	}

	got := gofp.Intersection(slices[0], slices[1], predicate)

	want := []person{
		{
			name:    "cc",
			age:     31,
			country: "DE",
		},
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
