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