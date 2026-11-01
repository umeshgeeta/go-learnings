package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	i := 0
	for i < len(nums)-2 {
		frozen := nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := frozen + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(closestSum-target)) {
				closestSum = sum
			}
		}
		i++
	}
	return closestSum
}

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
	nums = []int{0, 0, 0}
	target = 1
	fmt.Println(threeSumClosest(nums, target))
}
