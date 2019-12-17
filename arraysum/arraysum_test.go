package arraysum

import (
	"testing"
	"time"
)

func TestSumParallel(t *testing.T) {
	tests := [][]int{
		getArrayWithLength(0),
		getArrayWithLength(1),
		getArrayWithLength(2),
		getArrayWithLength(6),
		getArrayWithLength(24),
		getArrayWithLength(120),
		getArrayWithLength(720),
		getArrayWithLength(5040),
		getArrayWithLength(40320),
		getArrayWithLength(362880),
		getArrayWithLength(3628800),
	}

	for _, tt := range tests {
		sequential, parallel := Sum(tt), SumParallel(tt)

		if sequential != parallel {
			t.Fatalf("Unexpected result, expect %d got %d", sequential, parallel)
		}
	}
}

func BenchmarkSumParallel120(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(120)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = SumParallel(arr)
	}
}

func BenchmarkSumParallel720(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(720)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = SumParallel(arr)
	}
}

func BenchmarkSumParallel5040(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(5040)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = SumParallel(arr)
	}
}

func BenchmarkSumParallel3628800(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(3628800)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = SumParallel(arr)
	}
}

func BenchmarkSum120(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(120)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Sum(arr)
	}
}

func BenchmarkSum720(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(720)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Sum(arr)
	}
}

func BenchmarkSum5040(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(5040)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Sum(arr)
	}
}

func BenchmarkSum3628800(b *testing.B) {
	b.StopTimer()
	arr := getArrayWithLength(3628800)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = Sum(arr)
	}
}

func getArrayWithLength(l int) []int {
	arr := make([]int, 0, l)
	for i := 0; i < l; i++ {
		arr = append(arr, int(time.Now().Unix())%20)
	}

	return arr
}

