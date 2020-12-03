package combsort

/*
	CombSort 梳排序
*/

func CombSort(arr []int) []int {
	tmp := 0
	arrLen := len(arr)
	gap := arrLen

	for gap > 1 {
		gap = gap * 10 / 13 //递减率固定值1.3

		for i := 0; i+gap < arrLen; i++ {
			if arr[i] > arr[i+gap] {
				tmp = arr[i]
				arr[i] = arr[i+gap]
				arr[i+gap] = tmp
			}
		}
	}
	return arr
}
