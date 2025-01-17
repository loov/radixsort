package radixsort

import "math"

// Uint64 implements radix sort using secondary buffer.
//
// buf must be larger than arr.
func Uint64(arr, buf []uint64) {
	if len(arr) > len(buf) {
		panic("len(arr) > len(buf)")
	}
	buf = buf[:len(arr)]
	if len(arr) <= 1 {
		return
	}
	if int64(len(arr)) >= math.MaxUint32 {
		uint64_large(arr, buf)
		return
	}

	prev := arr[0]
	sorted := true

	var count [8][256]uint32
	for _, v := range arr {
		count[0][byte(v>>(0*8))]++
		count[1][byte(v>>(1*8))]++
		count[2][byte(v>>(2*8))]++
		count[3][byte(v>>(3*8))]++
		count[4][byte(v>>(4*8))]++
		count[5][byte(v>>(5*8))]++
		count[6][byte(v>>(6*8))]++
		count[7][byte(v>>(7*8))]++

		sorted = sorted && prev <= v
		prev = v
	}
	if sorted {
		return
	}

	var offset [8][256]uint32
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
		offset[4][k] = offset[4][k-1] + count[4][k-1]
		offset[5][k] = offset[5][k-1] + count[5][k-1]
		offset[6][k] = offset[6][k-1] + count[6][k-1]
		offset[7][k] = offset[7][k-1] + count[7][k-1]
	}

	swaps := 0
	src, dst := arr, buf
	for k := uint8(0); k < 8; k++ {
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

func uint64_large(arr, buf []uint64) {
	prev := arr[0]
	sorted := true

	var count [8][256]uint64
	for _, v := range arr {
		count[0][byte(v>>(0*8))]++
		count[1][byte(v>>(1*8))]++
		count[2][byte(v>>(2*8))]++
		count[3][byte(v>>(3*8))]++
		count[4][byte(v>>(4*8))]++
		count[5][byte(v>>(5*8))]++
		count[6][byte(v>>(6*8))]++
		count[7][byte(v>>(7*8))]++

		sorted = sorted && prev <= v
		prev = v
	}
	if sorted {
		return
	}

	var offset [8][256]uint64
	for k := 1; k < 256; k++ {
		offset[0][k] = offset[0][k-1] + count[0][k-1]
		offset[1][k] = offset[1][k-1] + count[1][k-1]
		offset[2][k] = offset[2][k-1] + count[2][k-1]
		offset[3][k] = offset[3][k-1] + count[3][k-1]
		offset[4][k] = offset[4][k-1] + count[4][k-1]
		offset[5][k] = offset[5][k-1] + count[5][k-1]
		offset[6][k] = offset[6][k-1] + count[6][k-1]
		offset[7][k] = offset[7][k-1] + count[7][k-1]
	}

	swaps := 0
	src, dst := arr, buf
	for k := uint8(0); k < 8; k++ {
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
