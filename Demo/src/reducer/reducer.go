/*
 * @Author: your name
 * @Date: 2020-01-09 18:46:23
 * @LastEditTime : 2020-01-09 19:04:11
 * @LastEditors  : Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /golang-learning-program/Demo/src/reducer/reducer.go
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genArr(length int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func reducer(cb func(int, int) int, arr []int) int {
	if len(arr) < 1 {
		return arr[0]
	}
	result := cb(arr[0], arr[1])
	for i := 2; i <= len(arr)-1; i++ {
		result = cb(result, arr[i])
	}
	return result
}

func add(total int, num int) int {
	return num + total
}
func mul(total int, num int) int {
	return num * total
}

func main() {
	arr := genArr(5)
	fmt.Println(arr)
	fmt.Println(reducer(add, arr))
	fmt.Println(reducer(mul, arr))
}
