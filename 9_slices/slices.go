package main

import "fmt"

// slice
// dynamic, most used construct in Go

func main() {
	// uninitialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 2, 5)
	// capacity -> maximum number of elements the slice can hold before it needs to be resized
	fmt.Println(nums == nil)
	fmt.Println(cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))
	nums[0] = 3
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	var nums2 = make([]int, len(nums))
	nums = append(nums, 5)
	copy(nums2, nums)
	fmt.Println(nums2)
}
