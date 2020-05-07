package slicex_test

import (
	"testing"

	. "github.com/msales/gox/slicex"
)

func TestInterfaceToString(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    []string
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []interface{}{"a", "b", "c"},
			want:    []string{"a", "b", "c"},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []interface{}{"a", 1, 3.33},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []interface{}{},
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "nil input",
			input:   nil,
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToString(tt.input)

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

func TestInterfaceToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    []int
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []interface{}{1, 2, 3},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []interface{}{"a", 1, 3.33},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []interface{}{},
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
			got, err := InterfaceToInt(tt.input)

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

func TestInterfaceToInt32(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    []int32
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []interface{}{int32(1), int32(2), int32(3)},
			want:    []int32{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []interface{}{"a", 1, 3.33},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []interface{}{},
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
			got, err := InterfaceToInt32(tt.input)

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

func TestInterfaceToInt64(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    []int64
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []interface{}{int64(1), int64(2), int64(3)},
			want:    []int64{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []interface{}{"a", 1, 3.33},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []interface{}{},
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
			got, err := InterfaceToInt64(tt.input)

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

func TestInterfaceToFloat32(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    []float32
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []interface{}{float32(1), float32(2), float32(3)},
			want:    []float32{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []interface{}{"a", 1, 3.33},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []interface{}{},
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
			got, err := InterfaceToFloat32(tt.input)

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

func TestInterfaceToFloat64(t *testing.T) {
	tests := []struct {
		name    string
		input   []interface{}
		want    []float64
		wantErr bool
	}{
		{
			name:    "correct input",
			input:   []interface{}{float64(1), float64(2), float64(3)},
			want:    []float64{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "incorrect input",
			input:   []interface{}{"a", 1, 3.33},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []interface{}{},
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
			got, err := InterfaceToFloat64(tt.input)

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
