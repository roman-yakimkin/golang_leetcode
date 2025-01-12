package main

func SearchInsert(nums []int, target int) int {
	minIdx, maxIdx := 0, len(nums)-1
	var position int

	if target <= nums[minIdx] {
		return 0
	}
	if target > nums[maxIdx] {
		return len(nums)
	}
	for minIdx < maxIdx {
		position = (minIdx + maxIdx) / 2
		if position == minIdx {
			return position + 1
		}
		if position == maxIdx {
			return position
		}
		if target == nums[position] {
			return position
		}
		if target > nums[position] {
			minIdx = position
		}
		if target < nums[position] {
			maxIdx = position
		}
	}

	return position
}
