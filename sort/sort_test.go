package sort

import (
	"fmt"
	"testing"
)

func Test_merge_sort(t *testing.T) {
	arr := []int{5, 4, 9, 3, 2, 0, 7, 5, 1}
	res := MergeSort(arr)
	fmt.Println(res)
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	t.Log(binarySearch(arr, 13))
}

func TestFirstEqual(t *testing.T) {
	arr := []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 11, 13, 13}
	t.Log(binaryFindFirstEqual(arr, 13))
}

func TestLastEqual(t *testing.T) {
	arr := []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 11, 13, 13}
	t.Log(binaryFindLastEqual(arr, 3))
}

func TestFirstGte(t *testing.T) {
	arr := []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 11, 13, 13}
	t.Log(binaryFirstGte(arr, 1))
}

func TestLastLte(t *testing.T) {
	arr := []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 11, 13, 13}
	t.Log(binaryLastLte(arr, 1))
}
