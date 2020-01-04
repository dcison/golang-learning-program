/*
 * @Author: dcison
 * @Date: 2020-01-04 16:36:07
 * @LastEditTime : 2020-01-04 17:09:35
 * @Description: 斐波那契数列demo 主要学习递归的用法
 * @FilePath: /golang-learning-program/Demo/src/fibon.go
 */
package main

import (
	"fmt"
)

func fibon(num int) int {
	switch num {
	case 1:
		return 0
	case 2:
		return 1
	default:
		return fibon(num-1) + fibon(num-2)
	}
}

func fibonBylast(num float64, res1, res2 int) float64 {
	if num == 1 {
		return float64(res1)
	}
	return fibonBylast(num-1, res2, res1+res2)
}

func main() {
	resByNormal, resByLast := make(chan int), make(chan float64)
	go func() { resByNormal <- fibon(20) }()
	go func() { resByLast <- fibonBylast(20.0, 0, 1) }()
	fmt.Printf("斐波那契数列正常测试： %d \n", <-resByNormal)
	fmt.Printf("斐波那契数列尾递归测试： %f \n", <-resByLast)
}
