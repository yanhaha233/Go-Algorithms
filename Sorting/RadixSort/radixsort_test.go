package radixsort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := RadixSort([]int32{2, 3, 5, 1, 4, 6, 7, 9, 8})
	want := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
