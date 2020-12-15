package SelectionSort

/*
	选择排序
*/

func SelectionSort(arr []int) []int {
	var min = 0
	var tmp = 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
	}
	return arr
}
