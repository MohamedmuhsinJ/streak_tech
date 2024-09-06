package main

func findPairs(numbers []int, target int) [][]int {
	left, right := 0, len(numbers)-1
	ans := [][]int{}
	sum := 0
	//using two pointer approach to find the sub array sum
	for left < right {
		sum = numbers[left] + numbers[right]
		if sum > target {
			//if sum  is greater  then target then we can assume that that's not the subarray pair
			//check for others by reducing the sum
			right--
		} else if sum < target {
			//if sum  is lesser  then target then we can assume that that's not the subarray pair
			//check for others by increasing the sum
			left++
		} else {
			//after find the pairs go to the next one by increasing left and decreasing right then we can avoid same index duplications
			ans = append(ans, []int{left, right})
			left++
			right--
		}
	}

	return ans
}
