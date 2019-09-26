package radixsort_test

import (
	"testing"

	"github.com/loov/radixsort"
	"github.com/zeebo/pcg"
)

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

func BenchmarkOverhead32_1e4(b *testing.B) { bench32(b, 1e4, func(x, y []uint32) {}) }
func BenchmarkOverhead32_1e6(b *testing.B) { bench32(b, 1e6, func(x, y []uint32) {}) }
func BenchmarkOverhead64_1e4(b *testing.B) { bench64(b, 1e4, func(x, y []uint64) {}) }
func BenchmarkOverhead64_1e6(b *testing.B) { bench64(b, 1e6, func(x, y []uint64) {}) }

func BenchmarkUint32_1e4(b *testing.B) { bench32(b, 1e4, radixsort.Uint32) }
func BenchmarkUint32_1e6(b *testing.B) { bench32(b, 1e6, radixsort.Uint32) }
func BenchmarkUint64_1e4(b *testing.B) { bench64(b, 1e4, radixsort.Uint64) }
func BenchmarkUint64_1e6(b *testing.B) { bench64(b, 1e6, radixsort.Uint64) }
