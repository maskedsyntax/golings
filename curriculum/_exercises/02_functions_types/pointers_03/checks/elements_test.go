package elements

import (
	"slices"
	"testing"
)

func TestElementPointersReadValues(t *testing.T) {
	nums := []int{7, 8, 9}
	ptrs := ElementPointers(nums)
	if len(ptrs) != len(nums) {
		t.Fatalf("got %d pointers, want %d", len(ptrs), len(nums))
	}
	for i, p := range ptrs {
		if p == nil || *p != nums[i] {
			t.Errorf("*ptrs[%d] = %v, want %d", i, p, nums[i])
		}
	}
}

func TestElementPointersWriteThrough(t *testing.T) {
	nums := []int{1, 2, 3}
	ptrs := ElementPointers(nums)
	*ptrs[1] = 20
	*ptrs[2]++
	if !slices.Equal(nums, []int{1, 20, 4}) {
		t.Errorf("after writing through the pointers, nums = %v, want [1 20 4]", nums)
	}
	nums[0] = 100
	if *ptrs[0] != 100 {
		t.Errorf("after nums[0] = 100, *ptrs[0] = %d, want 100", *ptrs[0])
	}
}
