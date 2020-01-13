/*
 * @Author: dcison
 * @Date: 2020-01-09 19:30:17
 * @LastEditTime : 2020-01-09 20:45:17
 * @Description: 用于随机生成数组 传入的为数组长度
 * @FilePath: /golang-learning-program/Demo/src/genRamArr/main.go
 */
package genarr

import (
	"math/rand"
	"time"
)

func Genarr(length int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}
