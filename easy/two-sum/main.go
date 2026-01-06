package main

import "fmt"

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9

	// nums := []int{3, 2, 4}
	// target := 6

	nums := []int{3, 3}
	target := 6

	solution := twoSum(nums, target)

	fmt.Println(solution)
}

func twoSum(nums []int, target int) []int {
	var twoSum []int

	fmt.Println("EASY - TWO SUM")
	fmt.Println("-------------------")

	for i := range len(nums) {
		for j := range len(nums) {
			if nums[i]+nums[j] == target && i != j {
				twoSum = append(twoSum, i, j)
				return twoSum
			}
		}
	}

	return twoSum
}
