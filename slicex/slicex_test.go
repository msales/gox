package slicex_test

import (
	"testing"

	"github.com/msales/gox/slicex"
	"gotest.tools/assert"
)

func TestStringToInterface(t *testing.T) {
	arg := []string{"a", "b", "c"}
	want := []interface{}{"a", "b", "c"}
	got := slicex.StringToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestIntToInterface(t *testing.T) {
	arg := []int{1, 2, 3}
	want := []interface{}{1, 2, 3}
	got := slicex.IntToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestInt32ToInterface(t *testing.T) {
	arg := []int32{1, 2, 3}
	want := []interface{}{int32(1), int32(2), int32(3)}
	got := slicex.Int32ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestInt64ToInterface(t *testing.T) {
	arg := []int64{1, 2, 3}
	want := []interface{}{int64(1), int64(2), int64(3)}
	got := slicex.Int64ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestFloat32ToInterface(t *testing.T) {
	arg := []float32{1, 2, 3}
	want := []interface{}{float32(1), float32(2), float32(3)}
	got := slicex.Float32ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestFloat64ToInterface(t *testing.T) {
	arg := []float64{1, 2, 3}
	want := []interface{}{float64(1), float64(2), float64(3)}
	got := slicex.Float64ToInterface(arg)

	assert.DeepEqual(t, want, got)
}

func TestInterfaceToString(t *testing.T) {
	argSuccess := []interface{}{"a", "b", "c"}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []string{"a", "b", "c"}

	gotSuccess, err := slicex.InterfaceToString(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := slicex.InterfaceToString(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToInt(t *testing.T) {
	argSuccess := []interface{}{1, 2, 3}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []int{1, 2, 3}

	gotSuccess, err := slicex.InterfaceToInt(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := slicex.InterfaceToInt(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToInt32(t *testing.T) {
	argSuccess := []interface{}{int32(1), int32(2), int32(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []int32{1, 2, 3}

	gotSuccess, err := slicex.InterfaceToInt32(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := slicex.InterfaceToInt32(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToInt64(t *testing.T) {
	argSuccess := []interface{}{int64(1), int64(2), int64(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []int64{1, 2, 3}

	gotSuccess, err := slicex.InterfaceToInt64(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := slicex.InterfaceToInt64(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToFloat32(t *testing.T) {
	argSuccess := []interface{}{float32(1), float32(2), float32(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []float32{1, 2, 3}

	gotSuccess, err := slicex.InterfaceToFloat32(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := slicex.InterfaceToFloat32(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}

func TestInterfaceToFloat64(t *testing.T) {
	argSuccess := []interface{}{float64(1), float64(2), float64(3)}
	argFailure := []interface{}{"a", 1, 3.33}
	want := []float64{1, 2, 3}

	gotSuccess, err := slicex.InterfaceToFloat64(argSuccess)
	assert.NilError(t, err)
	assert.DeepEqual(t, want, gotSuccess)

	gotFailure, err := slicex.InterfaceToFloat64(argFailure)
	assert.Assert(t, gotFailure == nil)
	assert.Assert(t, err != nil)
}