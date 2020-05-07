package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestIntToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []interface{}{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []interface{}{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToInterface(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []string
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []string{"1", "2", "3"},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []string{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToString(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestIntToInt32(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int32
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []int32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []int32{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []int32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToInt32(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestIntToInt64(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int64
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []int64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []int64{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToInt64(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestIntToFloat32(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []float32
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []float32{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []float32{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []float32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToFloat32(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestIntToFloat64(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []float64
	}{
		{
			name:  "correct input",
			input: []int{1, 2, 3},
			want:  []float64{1, 2, 3},
		},
		{
			name:  "empty input",
			input: []int{},
			want:  []float64{},
		},
		{
			name:  "nil input",
			input: nil,
			want:  []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntToFloat64(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if len(tt.want) != len(got) {
				t.Errorf("Got %+v, want %+v", got, tt.want)
				return
			}

			for k, v := range tt.want {
				if v != got[k] {
					t.Errorf("Got %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}
