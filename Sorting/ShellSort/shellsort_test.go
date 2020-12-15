package ShellSort

import (
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	got := ShellSort([]int{2, 1, 3, 5, 4, 7, 6, 9, 8})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
