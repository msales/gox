package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
	"gotest.tools/assert"
)

func TestStringToInterface(t *testing.T) {
	arg := []string{"a", "b", "c"}
	want := []interface{}{"a", "b", "c"}
	got := StringToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestIntToInterface(t *testing.T) {
	arg := []int{1, 2, 3}
	want := []interface{}{1, 2, 3}
	got := IntToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestInt32ToInterface(t *testing.T) {
	arg := []int32{1, 2, 3}
	want := []interface{}{int32(1), int32(2), int32(3)}
	got := Int32ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestInt64ToInterface(t *testing.T) {
	arg := []int64{1, 2, 3}
	want := []interface{}{int64(1), int64(2), int64(3)}
	got := Int64ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestFloat32ToInterface(t *testing.T) {
	arg := []float32{1, 2, 3}
	want := []interface{}{float32(1), float32(2), float32(3)}
	got := Float32ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestFloat64ToInterface(t *testing.T) {
	arg := []float64{1, 2, 3}
	want := []interface{}{float64(1), float64(2), float64(3)}
	got := Float64ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestInterfaceToString(t *testing.T) {
	argSuccess := []interface{}{"a", "b", "c"}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []string{"a", "b", "c"}

	gotSuccess, err := InterfaceToString(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := InterfaceToString(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToInt(t *testing.T) {
	argSuccess := []interface{}{1, 2, 3}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []int{1, 2, 3}

	gotSuccess, err := InterfaceToInt(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := InterfaceToInt(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToInt32(t *testing.T) {
	argSuccess := []interface{}{int32(1), int32(2), int32(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []int32{1, 2, 3}

	gotSuccess, err := InterfaceToInt32(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := InterfaceToInt32(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToInt64(t *testing.T) {
	argSuccess := []interface{}{int64(1), int64(2), int64(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []int64{1, 2, 3}

	gotSuccess, err := InterfaceToInt64(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := InterfaceToInt64(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToFloat32(t *testing.T) {
	argSuccess := []interface{}{float32(1), float32(2), float32(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []float32{1, 2, 3}

	gotSuccess, err := InterfaceToFloat32(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := InterfaceToFloat32(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToFloat64(t *testing.T) {
	argSuccess := []interface{}{float64(1), float64(2), float64(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []float64{1, 2, 3}

	gotSuccess, err := InterfaceToFloat64(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := InterfaceToFloat64(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestContainsInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		contains []int
		want     bool
	}{
		{
			name:     "a slice containing exacly 1 element",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{3},
			want:     true,
		},
		{
			name:     "a slice containing exacly 3 elements",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{3, 5, 6},
			want:     true,
		},
		{
			name:     "a slice containing only part of needed",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{3, 5, 6, 7},
			want:     false,
		},
		{
			name:     "a slice not contains element with 1 element to check",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{7},
			want:     false,
		},
		{
			name:     "a slice not contains any element with 2 elements to check",
			slice:    []int{1, 2, 3, 4, 5, 6},
			contains: []int{7, 9},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ContainsInt(tt.slice, tt.contains...))
		})
	}
}

func BenchmarkContainsIntOne(b *testing.B) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	checkOne := 4
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ContainsInt(slice, checkOne)
	}
}

func BenchmarkContainsIntMany(b *testing.B) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	checkMany := []int{2, 4, 9, 10}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = ContainsInt(slice, checkMany...)
	}
}
