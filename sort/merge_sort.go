package sort

// MergeSort x_x
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	index := len(arr) / 2

	a1 := MergeSort(arr[:index])
	a2 := MergeSort(arr[index:])

	return merge(a1, a2)
}

func merge(a1 []int, a2 []int) []int {
	i, j, k := 0, 0, 0
	res := make([]int, len(a1)+len(a2))
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			res[k] = a1[i]
			i++
		} else {
			res[k] = a2[j]
			j++
		}
		k++
	}
	for j < len(a2) {
		res[k] = a2[j]
		j++
		k++
	}
	for i < len(a1) {
		res[k] = a1[i]
		i++
		k++
	}
	return res
}
