// +build ignore

package radixsort_test

import (
	"testing"

	"github.com/shawnsmithdev/zermelo/zuint32"
	"github.com/shawnsmithdev/zermelo/zuint64"
)

func BenchmarkZermelo32_1e4(b *testing.B) { bench32(b, 1e4, zuint32.SortBYOB) }
func BenchmarkZermelo32_1e6(b *testing.B) { bench32(b, 1e6, zuint32.SortBYOB) }
func BenchmarkZermelo64_1e4(b *testing.B) { bench64(b, 1e4, zuint64.SortBYOB) }
func BenchmarkZermelo64_1e6(b *testing.B) { bench64(b, 1e6, zuint64.SortBYOB) }
