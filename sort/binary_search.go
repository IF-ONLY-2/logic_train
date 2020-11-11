package sort

func binarySearch(arr []int, data int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == data {
			return mid
		} else if arr[mid] > data {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func binaryFindFirstEqual(arr []int, data int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == 0 || arr[mid-1] != data {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func binaryFindLastEqual(arr []int, data int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > data {
			high = mid - 1
		} else if arr[mid] < data {
			low = mid + 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != data {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

func binaryFirstGte(arr []int, data int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] >= data {
			if mid == 0 || arr[mid-1] < data {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func binaryLastLte(arr []int, data int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if arr[mid] > data {
			high = mid - 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] > data {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
