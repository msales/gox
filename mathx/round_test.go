package mathx_test

import (
	"testing"

	. "github.com/msales/gox/mathx"
	"gotest.tools/assert"
)

func TestRound(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		decimals float64
		want     float64
	}{
		{
			name:     "Round number with more decimals to 2 decimals last number down",
			in:       91.23234543,
			decimals: 2,
			want:     91.23,
		},
		{
			name:     "Round number with more decimals to 2 decimals last number up",
			in:       91.23634543,
			decimals: 2,
			want:     91.24,
		},
		{
			name:     "Round number with less decimals to 4 decimals",
			in:       91.23,
			decimals: 4,
			want:     91.23,
		},
		{
			name:     "Small number",
			in:       0.00000000017,
			decimals: 10,
			want:     0.0000000002,
		},
		{
			name:     "No decimals",
			in:       123.321,
			decimals: 0,
			want:     123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Round(tt.in, tt.decimals))
		})
	}
}
