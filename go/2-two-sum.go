// Again, here I added main() func for display purposes.
// You should delete import("fmt") or else Go will throw an error. >:/

package main

import (
    "fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] + nums[j] == target && i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

