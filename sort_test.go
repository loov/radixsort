package radixsort_test

import (
	"reflect"
	"sort"
	"testing"
	"testing/quick"

	"github.com/loov/radixsort"
)

func TestUint32_Quick(t *testing.T) {
	if err := quick.Check(CheckUint32, nil); err != nil {
		t.Error(err)
	}
}

func TestUint64_Quick(t *testing.T) {
	if err := quick.Check(CheckUint64, nil); err != nil {
		t.Error(err)
	}
}

func TestUint32(t *testing.T) {
	tests := [][]uint32{
		{},
		{0},
		{0xFF},
		{0xFFFFFFFF},
		{0, 1},
		{1, 0},
		{0, 0xFF},
		{0xFF, 0},
		{0, 0xFFFFFFFF},
		{0xFFFFFFFF, 0},

		{74, 59, 238, 784, 9845, 959, 905, 0, 0, 42, 7586, 5467984, 7586},
	}

	for _, test := range tests {
		if !CheckUint32(test) {
			t.Errorf("failed on input %v", test)
		}
	}
}

func TestUint64(t *testing.T) {
	tests := [][]uint64{
		{},
		{0},
		{0xFF},
		{0xFFFFFFFFFFFFFFFF},
		{0, 1},
		{1, 0},
		{0, 0xFF},
		{0xFF, 0},
		{0, 0xFFFFFFFFFFFFFFFF},
		{0xFFFFFFFFFFFFFFFF, 0},
		{0xFFFFFFFFFFFFFFF0, 0xFFFFFFFFFFFFFFFF},
		{0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFF0},

		{74, 59, 238, 784, 9845, 959, 905, 0, 0, 42, 7586, 5467984, 7586},
	}

	for _, test := range tests {
		if !CheckUint64(test) {
			t.Errorf("failed on input %v", test)
		}
	}
}

func CheckUint32(data []uint32) bool {
	expected := make([]uint32, len(data))
	copy(expected, data)
	sort.Slice(expected, func(i, k int) bool {
		return expected[i] < expected[k]
	})

	sorting := make([]uint32, len(data))
	copy(sorting, data)
	buffer := make([]uint32, len(data))
	radixsort.Uint32(sorting, buffer)
	return reflect.DeepEqual(expected, sorting)
}

func CheckUint64(data []uint64) bool {
	expected := make([]uint64, len(data))
	copy(expected, data)
	sort.Slice(expected, func(i, k int) bool {
		return expected[i] < expected[k]
	})

	sorting := make([]uint64, len(data))
	copy(sorting, data)
	buffer := make([]uint64, len(data))
	radixsort.Uint64(sorting, buffer)
	return reflect.DeepEqual(expected, sorting)
}
