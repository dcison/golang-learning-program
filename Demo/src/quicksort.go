/*
 * @Author: dcison
 * @Date: 2020-01-04 19:23:13
 * @LastEditTime : 2020-01-04 19:58:05
 * @Description: 快排算法
 * @FilePath: /golang-learning-program/Demo/src/quicksort.go
 */
package main

import "fmt"

// func quickSort(array []int) []int { // 标杆取自 数组第一个元素
// 	if len(array) < 1 {
// 		return array
// 	}
// 	mid := array[0]
// 	left, right := 0, len(array)-1
// 	for i := 1; i <= right; {
// 		if array[i] > mid {
// 			array[i], array[right] = array[right], array[i]
// 			right--
// 		} else {
// 			array[i], array[left] = array[left], array[i]
// 			left++
// 			i++
// 		}
// 	}
// 	array[left] = mid
// 	quickSort(array[:left])
// 	quickSort(array[left+1:])
// 	return array
// }

func quickSort(array []int) []int { // 标杆取自 数组中间元素
	if len(array) < 1 {
		return array
	}
	mid := array[len(array)/2]
	lArray, rArray, mArray := []int{}, []int{}, []int{}
	for _, number := range array {
		switch {
		case number > mid:
			rArray = append(rArray, number)
		case number < mid:
			lArray = append(lArray, number)
		default:
			mArray = append(mArray, number)
		}
	}
	lArray = quickSort(lArray)
	rArray = quickSort(rArray)

	return append(append(lArray, mArray...), rArray...)
}

func main() {
	arr := []int{53, 18, 68, 0, 98, 90, 26, 60, 33}
	fmt.Println(quickSort(arr))
}
