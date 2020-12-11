package oddevensort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := OddEvenSort([]int{3, 2, 7, 6, 8, 1, 4, 5})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
