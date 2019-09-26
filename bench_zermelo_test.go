// +build ignore

package radixsort_test

import (
	"testing"

	"github.com/shawnsmithdev/zermelo/zuint32"
	"github.com/shawnsmithdev/zermelo/zuint64"
)

func BenchmarkZermeloUint32(b *testing.B) {
	for _, bench := range BenchmarkSizes {
		b.Run(bench.Name, func(b *testing.B) {
			bench32(b, bench.Size, zuint32.SortBYOB)
		})
	}
}
func BenchmarkZermeloUint64(b *testing.B) {
	for _, bench := range BenchmarkSizes {
		b.Run(bench.Name, func(b *testing.B) {
			bench64(b, bench.Size, zuint64.SortBYOB)
		})
	}
}
