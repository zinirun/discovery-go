package datastructure

import (
	"sort"
	"testing"
)

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSort(t *testing.T) {
	n1 := []int{7, 6, 5, 4, 3, 2, 1, 0}
	n1Test := make([]int, len(n1))
	copy(n1Test, n1)
	n1Test = bubbleSort(n1Test)
	sort.Ints(n1)
	if !equal(n1, n1Test) {
		t.Error("Not sorted, ", bubbleSort(n1))
	}
}
