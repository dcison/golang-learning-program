/*
 * @Author: dcison
 * @Date: 2020-01-06 20:21:09
 * @LastEditTime : 2020-01-07 10:19:04
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /golang-learning-program/Demo/src/bubble.go
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubble(arr []int, order string) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		for j := length - 1; j > i; j-- {
			if order == "asc" && arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else if order == "desc" && arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func genArr(length int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func main() {
	testDemo := genArr(30)
	fmt.Printf("排序前 %v \n", testDemo)
	fmt.Printf("升序排序后 %v \n", bubble(testDemo, "asc"))
	fmt.Printf("降序排序后 %v \n", bubble(testDemo, "desc"))
}
