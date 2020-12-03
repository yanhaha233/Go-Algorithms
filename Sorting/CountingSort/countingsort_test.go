package countingsort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	arr := []int{3, 1, 5, 4, 2}
	got := CountingSort(arr)
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
