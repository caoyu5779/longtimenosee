package Leetcode

import "fmt"

func twoSum(nums []int, target int) int{
	fmt.Println(nums)
	for x := range nums{
		fmt.Println(x)
		if ((nums[x] + nums[x+1]) == target){
			return target
		}
	}
	return target
}

func main(){
		nums := []int{1,2,3,4,5}
		target := 6
		fmt.Println(target)
		twoSum(nums, target)
}