package other

func TwoSum(nums []int, target int) []int {
	length := len(nums)
	for index, value := range nums {
		for i := index+1; i < length; i++ {
			if value + nums[i] == target {
				return []int{index, i}
			}
		}
	}
	return []int{}
}

func TwoSumMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}