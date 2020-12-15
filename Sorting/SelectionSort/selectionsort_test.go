package SelectionSort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	got := SelectionSort([]int{2, 4, 5, 1, 3, 7, 6, 9, 8})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
