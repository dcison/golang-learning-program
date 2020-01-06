/*
 * @Author: dcison
 * @Date: 2020-01-04 18:03:43
 * @LastEditTime : 2020-01-04 18:46:33
 * @Description: 两数之和问题
 * @FilePath: /golang-learning-program/Demo/src/twosum.go
 */

package main

import (
	"fmt"
)

// func twoSum(nums []int, target int) []int {
// 	res := make([]int, 2)
// 	for index, number := range nums {
// 		_target := target - number
// 		for _index, _number := range nums[index+1:] {
// 			if _number == _target {
// 				res[0] = index
// 				res[1] = _index + index + 1
// 				return res
// 			}
// 		}
// 	}
// 	return nil
// }

func twoSum(nums []int, target int) []int {
	// 使用map进行优化
	resMap := make(map[int][]int, 0)
	for index, number := range nums {
		// 取出每一项的值 做key用
		value, ok := resMap[number]
		if ok {
			value = append(value, index)
		} else {
			resMap[number] = []int{index}
		}
	}
	for index, number := range nums {
		// 取出每一项的值 做key用
		value, ok := resMap[target-number]
		if ok {
			for _, _index := range value {
				if _index != index {
					return []int{index, _index}
				}
			}
		}
	}
	return nil
}
func main() {
	test := []int{11, 2, 2, 15}
	demo := twoSum(test, 13)
	fmt.Printf("2数测试： %d \n", demo)
}
