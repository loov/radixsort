# radixsort

`radixsort` package implements radix sorting for `int32` and `int64`.

The usage looks like:

```go
import "github.com/loov/radixsort"

func Example() {
	data := []int64{1,2,3,4}
	tmpbuf := make([]int64, len(data))

	radixsort.Uint64(data, tmpbuf)
}
```