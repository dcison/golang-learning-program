package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stus {
		// for range 创建了每个元素的副本，而不是直接返回每个元素的引用，如果使用该值变量的地址作为指向每个元素的指针，就会导致错误
		// https://studygolang.com/articles/9701
		stucopy := stu
		m[stu.name] = &stucopy
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
