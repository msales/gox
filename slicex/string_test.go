package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestStringToInterface(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []interface{}
	}{
		{
			name:  "correct input",
			input: []string{"a", "b", "c"},
			want:  []interface{}{"a", "b", "c"},
		},
		{
			name:  "empty input",
			input: []string{},
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
			got := StringToInterface(tt.input)

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

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []int
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []string{"1", "2", "3", "4", "5", "6"},
			want:    []int{1, 2, 3, 4, 5, 6},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []string{"1", "2", "3", "4", "5", "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []string{},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    []int{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("Expected error got nil")
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

func TestStringToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []int32
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []string{"1", "2", "3", "4", "5", "6"},
			want:    []int32{1, 2, 3, 4, 5, 6},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []string{"1", "2", "3", "4", "5", "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []string{},
			want:    []int32{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    []int32{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt32(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("Expected error got nil")
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

func TestStringToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []int64
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []string{"1", "2", "3", "4", "5", "6"},
			want:    []int64{1, 2, 3, 4, 5, 6},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []string{"1", "2", "3", "4", "5", "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []string{},
			want:    []int64{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    []int64{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt64(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("Expected error got nil")
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

func TestStringToFloat32(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []float32
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []string{"1.33", "2", "3", "4", "5", "6"},
			want:    []float32{1.33, 2, 3, 4, 5, 6},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []string{"1.33", "2", "3", "4", "5", "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []string{},
			want:    []float32{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    []float32{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToFloat32(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("Expected error got nil")
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

func TestStringToFloat64(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []float64
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []string{"1.33", "2", "3", "4", "5", "6"},
			want:    []float64{1.33, 2, 3, 4, 5, 6},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []string{"1.33", "2", "3", "4", "5", "test"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []string{},
			want:    []float64{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    []float64{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToFloat64(tt.input)

			if tt.want == nil && got != nil {
				t.Errorf("Expected nil result got %+v instead", got)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("Expected error got nil")
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

func TestContainsString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		contains []string
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a"},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a", "b", "c"},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a", "B", "C"},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A", "B"},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A"},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsString(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestMatchesString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		contains []string
		want     bool
	}{
		{
			name:     "found 1 element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a"},
			want:     true,
		},
		{
			name:     "found 3 elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a", "b", "c"},
			want:     true,
		},
		{
			name:     "found 1 element insensitive",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A"},
			want:     true,
		},
		{
			name:     "found 3 elements insensitive",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"A", "B", "C"},
			want:     true,
		},
		{
			name:     "found some but not all elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"a", "x", "y"},
			want:     false,
		},
		{
			name:     "found no elements",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"x", "y"},
			want:     false,
		},
		{
			name:     "found no element",
			slice:    []string{"a", "b", "c", "d", "e", "f"},
			contains: []string{"x"},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchesString(tt.slice, tt.contains...)
			if tt.want != got {
				t.Errorf("Got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkContainsString_Single(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsString(haystack, "c")
	}
}

func BenchmarkContainsString_Multiple(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ContainsString(haystack, "b", "c")
	}
}

func BenchmarkMatchesString_Single(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MatchesString(haystack, "C")
	}
}

func BenchmarkMatchesString_Multiple(b *testing.B) {
	haystack := []string{"a", "b", "c", "d", "e", "f"}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		MatchesString(haystack, "B", "C")
	}
}

func TestTrimString(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []string
	}{
		{
			name:    "no trimming of a full slice ",
			input:   []string{"a", "b", "c", "d", "e", "f"},
			want:    []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:    "trimming of a full slice ",
			input:   []string{"a", "", "b", " ", "c", "\t", "d", "\r", "e", "\n", "f"},
			want:    []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:    "complete trimming of a slice ",
			input:   []string{"", " ", "\t", "\r", "\n"},
			want:    nil,
		},
		{
			name:    "trimming of an empty slice ",
			input:   []string{},
			want:    nil,
		},
		{
			name:    "trimming of a nil slice ",
			input:   nil,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimString(tt.input)

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