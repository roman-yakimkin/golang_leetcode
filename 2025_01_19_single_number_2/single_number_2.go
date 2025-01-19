package main

func SingleNumber(nums []int) int {
	numSet := make(map[int]int, len(nums)/3)
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		numSet[key] = numSet[key] + 1
	}

	var result int
	for i, cnt := range numSet {
		if cnt == 1 {
			result = i
			break
		}
	}

	return result
}
