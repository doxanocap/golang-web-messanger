package main

import (
	"fmt"
)

func main() {
	fmt.Println(reachableNodes(7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}))
}

func validPartition(nums []int) bool {
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] == nums[i+1] {
            return true
        }
        if i < len(nums)-2 {
            if (nums[i] == nums[i+1] && nums[i] ==nums[i+2]) || ( nums[i] + 1 == nums[i+1] && nums[i] + 2 == nums[i+2]) {
                return true
            }
        } 
    }
    return false
}