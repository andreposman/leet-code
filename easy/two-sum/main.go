package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// nums := []int{3, 2, 4}
	// target := 6

	// nums := []int{3, 3}
	// target := 6

	solution := twoSum(nums, target)

	fmt.Println(solution)
}

func twoSum(nums []int, target int) []int {
	fmt.Println("EASY - TWO SUM")
	fmt.Println("-------------------------")

	seen := make(map[int]int)

	for i, number := range nums {
		partnerNum := target - number

		if partnerIndex, exists := seen[partnerNum]; exists {
			return []int{partnerIndex, i}
		}

		seen[number] = i

	}
	return nil
}
