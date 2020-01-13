package randx_test

import (
	"math/rand"
	"testing"

	"github.com/msales/gox/randx"
)

func TestHappens(t *testing.T) {
	tests := []struct {
		name string
		p    float32
		seed int64
		want bool
	}{
		{
			name: "happens because of seeded chance",
			p:    0.2,
			seed: 9,
			want: true,
		},
		{
			name: "does not happen because of seeded chance",
			p:    0.2,
			seed: 0,
			want: false,
		},
		{
			name: "happens because of border probability",
			p:    1,
			seed: 0,
			want: true,
		},
		{
			name: "does not happen because of border probability",
			p:    0,
			seed: 0,
			want: false,
		},
		{
			name: "happens because of exceeding probability",
			p:    1.5,
			seed: 9,
			want: true,
		},
		{
			name: "does not happen because of negative exceeding probability",
			p:    -1.5,
			seed: 0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(tt.seed)
			if got := randx.Happens(tt.p); got != tt.want {
				t.Errorf("Happens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhichHappens(t *testing.T) {
	tests := []struct {
		name string
		p []float32
		seed int64
		want int
	}{
		{
			name: "no probability happens due to seeded chance",
			p:    []float32{0.15,0.10},
			seed: 0,
			want: -1,
		},
		{
			name: "first probability happens due to seeded chance",
			p:    []float32{0.15,0.10},
			seed: 9,
			want: 0,
		},
		{
			name: "second probability happens due to seeded chance",
			p:    []float32{0.15,0.10},
			seed: 4,
			want: 1,
		},
		{
			name: "first probability due to its exceeding probability",
			p:    []float32{1.5,0.10},
			seed: 0,
			want: 0,
		},
		{
			name: "no probability due to negatively exceeding first probability",
			p:    []float32{-1.5,0.10},
			seed: 9,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rand.Seed(tt.seed)
			if got := randx.WhichHappens(tt.p...); got != tt.want {
				t.Errorf("WhichHappens() = %v, want %v", got, tt.want)
			}
		})
	}
}