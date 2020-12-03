package gnomesort

import (
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	got := GnomeSort([]int{3, 4, 5, 6, 1, 2})
	want := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
