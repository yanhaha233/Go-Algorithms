package gnomesort

/*
	GnomeSort    地精排序
*/

func GnomeSort(arr []int)[]int {
	tmp := 0
	i := 1

	for i < len(arr) {
		if arr[i] > arr[i-1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp

			if i > 1 {
				i--
			}
		}
	}

	return arr
}
