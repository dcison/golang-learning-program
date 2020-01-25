package main

import "fmt"

func testAnd() bool {
	fmt.Println("and")
	return true
}

func testOr() bool {
	fmt.Println("or")
	return false
}

func main() {
	if testOr() || testAnd() && true {
		fmt.Println("测试顺序")
	}
}