package twoSum

/*
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	for i, iv := range nums {
		for j, jv := range nums {
			if i == j {
				continue
			}
			if iv+jv == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
