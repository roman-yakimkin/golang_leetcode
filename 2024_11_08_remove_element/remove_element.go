package main

func RemoveElement(nums []int, val int) int {
	left, right := 0, len(nums)-1

	var changed bool
	for left <= right {
		if nums[left] == val {
			changed = true
			for nums[right] == val && right > left {
				right--
			}
			if left < right {
				nums[left], nums[right] = nums[right], nums[left]
			}
		}
		left++
	}
	if changed {
		left--
	}

	return left
}
