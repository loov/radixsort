package radixsort

import "math"

// Uint32 implements radix sort using secondary buffer.
//
// buf must be larger than arr.
func Uint32(arr, buf []uint32) {
	if len(arr) > len(buf) {
		panic("len(arr) > len(buf)")
	}
	buf = buf[:len(arr)]
	if len(arr) <= 1 {
		return
	}
	if int64(len(arr)) >= math.MaxUint32 {
		uint32_large(arr, buf)
		return
	}

	prev := arr[0]
	sorted := true

	var count [4][256]uint32
	for _, v := range arr {
		count[0][byte(v>>(0*8))]++
		count[1][byte(v>>(1*8))]++
		count[2][byte(v>>(2*8))]++
		count[3][byte(v>>(3*8))]++

		sorted = sorted && prev <= v
		prev = v
	}
	if sorted {
		return
	}

	var offset [4][256]uint32
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
	}

	swaps := 0
	src, dst := arr, buf
	for k := uint8(0); k < 4; k++ {
		nz := 0
		cnt := &count[k]
		for i := range cnt {
			if cnt[i] != 0 {
				nz++
			}
		}
		if nz == 1 {
			continue
		}
		swaps++

		off := &offset[k]
		p := k * 8

		for _, v := range src {
			key := uint8(v >> p)
			dst[off[key]] = v
			off[key]++
		}

		dst, src = src, dst
	}

	if swaps&1 == 1 {
		copy(arr, src)
	}
}

func uint32_large(arr, buf []uint32) {
	prev := arr[0]
	sorted := true

	var count [4][256]uint64
	for _, v := range arr {
		count[0][byte(v>>(0*8))]++
		count[1][byte(v>>(1*8))]++
		count[2][byte(v>>(2*8))]++
		count[3][byte(v>>(3*8))]++

		sorted = sorted && prev <= v
		prev = v
	}
	if sorted {
		return
	}

	var offset [4][256]uint64
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
	}

	swaps := 0
	src, dst := arr, buf
	for k := uint8(0); k < 4; k++ {
		nz := 0
		cnt := &count[k]
		for i := range cnt {
			if cnt[i] != 0 {
				nz++
			}
		}
		if nz == 1 {
			continue
		}
		swaps++

		off := &offset[k]
		p := k * 8

		for _, v := range src {
			key := uint8(v >> p)
			dst[off[key]] = v
			off[key]++
		}

		dst, src = src, dst
	}

	if swaps&1 == 1 {
		copy(arr, src)
	}
}
