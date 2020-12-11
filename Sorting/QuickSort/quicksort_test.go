package quicksort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := QuickSort([]int{3, 2, 6, 9, 7, 8, 1, 5, 4})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
