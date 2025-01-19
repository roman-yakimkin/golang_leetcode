package main

func SingleNumber(nums []int) int {
	numSet := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		_, ok := numSet[nums[i]]
		if !ok {
			numSet[nums[i]] = struct{}{}
		} else {
			delete(numSet, nums[i])
		}
	}

	var result int
	for i := range numSet {
		result = i
	}

	return result
}
