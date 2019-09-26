package radixsort_test

import (
	"testing"

	"github.com/loov/radixsort"
	"github.com/zeebo/pcg"
)

type SizeBenchmark struct {
	Name string
	Size int
}

var BenchmarkSizes = []SizeBenchmark{
	{"1e2", 1e2},
	{"1e3", 1e3},
	{"1e4", 1e4},
	{"1e5", 1e5},
	{"1e6", 1e6},
	{"1e7", 1e7},
}

func bench32(b *testing.B, size int, algo func(src, dst []uint32)) {
	rng := pcg.New(0)
	data := make([]uint32, size)
	buf := make([]uint32, len(data))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = rng.Uint32()
		}
		algo(data, buf)
	}
}

func bench64(b *testing.B, size int, algo func(src, dst []uint64)) {
	rng := pcg.New(uint64(0))
	data := make([]uint64, size)
	buf := make([]uint64, len(data))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(data); i++ {
			data[i] = rng.Uint64()
		}
		algo(data, buf)
	}
}

func BenchmarkOverhead32(b *testing.B) {
	for _, bench := range BenchmarkSizes {
		b.Run(bench.Name, func(b *testing.B) {
			bench32(b, bench.Size, func(x, y []uint32) {})
		})
	}
}

func BenchmarkOverhead64(b *testing.B) {
	for _, bench := range BenchmarkSizes {
		b.Run(bench.Name, func(b *testing.B) {
			bench64(b, bench.Size, func(x, y []uint64) {})
		})
	}
}

func BenchmarkUint32(b *testing.B) {
	for _, bench := range BenchmarkSizes {
		b.Run(bench.Name, func(b *testing.B) {
			bench32(b, bench.Size, radixsort.Uint32)
		})
	}
}
func BenchmarkUint64(b *testing.B) {
	for _, bench := range BenchmarkSizes {
		b.Run(bench.Name, func(b *testing.B) {
			bench64(b, bench.Size, radixsort.Uint64)
		})
	}
}
