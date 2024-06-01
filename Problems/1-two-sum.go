func twoSum(nums []int, target int) []int {
	// initialize map to store values of array and their indecies {value: index}
	itteratorMap := make(map[int]int)

	for i := 0; i < len(nums); i++{
		isInMap := target - nums[i]
		

		index, exsist := itteratorMap[isInMap]
		if exsist {
			return []int{index, i}
		}else {
			itteratorMap[nums[i]] = i
		}
	}
	return nil
}

