package main

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result, counter := "", 0

	for {
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == counter || strs[i][counter] != strs[0][counter] {
				return result
			}
		}
		result = result + string(strs[0][counter])
		counter++
	}
}
