// https://stackoverflow.com/questions/66643946/how-to-remove-duplicates-strings-or-int-from-slice-in-go
// I've modified it a bit from the original solution, so that I could see the output in terminal.
// For it to pass on leetcode, you just need to return len(nums) (;

package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 1, 2, 3}

	fmt.Println(removeDuplicates(nums...))
}

func removeDuplicates(nums ...int) []int {
    slices.Sort(nums)
    nums = slices.Compact(nums)

    return nums
}
