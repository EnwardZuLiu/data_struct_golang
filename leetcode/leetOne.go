package leetcode

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	var hash = map[int]int{}

	for i := 0; i < len(nums); i++ {
		cal := target - nums[i]
		if _, ok := hash[cal]; ok {
			return []int{hash[cal], i}
		}
		hash[nums[i]] = i
	}
	return []int{}
}
