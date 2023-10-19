package queuex_test

import (
	"github.com/msales/gox/queuex"
	"testing"
)

func TestFIFO(t *testing.T) {
	queue := queuex.NewFIFO([]int{
		1, 2,
	}...)

	queue.Push(3)

	got, ok := queue.Pop()
	wantVal := 1
	wantOK := true
	if wantVal != got {
		t.Errorf("Got %+v, want %+v", got, wantVal)
	}

	if ok != wantOK {
		t.Errorf("Got %+v, want %+v", ok, wantOK)
	}

	got, ok = queue.Pop()
	wantVal = 2
	wantOK = true
	if wantVal != got {
		t.Errorf("Got %+v, want %+v", got, wantVal)
	}

	if ok != wantOK {
		t.Errorf("Got %+v, want %+v", ok, wantOK)
	}

	got, ok = queue.Pop()
	wantVal = 3
	wantOK = true
	if wantVal != got {
		t.Errorf("Got %+v, want %+v", got, wantVal)
	}

	if ok != wantOK {
		t.Errorf("Got %+v, want %+v", ok, wantOK)
	}

	got, ok = queue.Pop()
	wantVal = 0
	wantOK = false
	if wantVal != got {
		t.Errorf("Got %+v, want %+v", got, wantVal)
	}

	if ok != wantOK {
		t.Errorf("Got %+v, want %+v", ok, wantOK)
	}
}

func TestFIFO_EmptyQueue(t *testing.T) {
	queue := queuex.NewFIFO[int]()

	got, ok := queue.Pop()
	wantVal := 0
	wantOK := false
	if wantVal != got {
		t.Errorf("Got %+v, want %+v", got, wantVal)
	}

	if ok != wantOK {
		t.Errorf("Got %+v, want %+v", ok, wantOK)
	}

	got, ok = queue.Pop()
	if wantVal != got {
		t.Errorf("Got %+v, want %+v", got, wantVal)
	}

	if ok != wantOK {
		t.Errorf("Got %+v, want %+v", ok, wantOK)
	}
}

func BenchmarkFIFO_Pop(b *testing.B) {
	vals := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = i
	}

	queue := queuex.NewFIFO(vals...)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		queue.Pop()
	}
}

func BenchmarkFIFO_Push(b *testing.B) {
	vals := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = i
	}

	queue := queuex.NewFIFO(vals...)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}

func BenchmarkFIFO_EmptyPush(b *testing.B) {
	queue := queuex.NewFIFO[int]()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}

func BenchmarkFIFO_RandomPushPop(b *testing.B) {
	vals := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = i
	}

	queue := queuex.NewFIFO(vals...)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if i%2 == 0 {
			queue.Pop()
			continue
		} else {
			queue.Push(i)
		}
	}
}
