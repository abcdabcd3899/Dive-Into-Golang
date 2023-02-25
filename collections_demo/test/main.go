package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	arr_init()
	arr_traverse()
	arr_truncate()
	slice_test()
	map_test()
	set_simulation() // use map
	fmt.Println("--------")
	slice_set_simulation() // use slice simulate set
	fmt.Println("--------")
	stack_simulation() // use slice
	queue_simulation() // use slice
	os.Exit(0)
}

func arr_init() {
	var arr1 [4]int
	arr2 := [4]int{1, 2, 3, 4}
	arr1[1] = 1
	fmt.Println(arr1, arr2)
	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr3)
}

func arr_traverse() {
	arr := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("--------")

	for index, ele := range arr {
		fmt.Println(index, ele)
	}

	fmt.Println("--------")

	for _, ele := range arr {
		fmt.Println(ele)
	}
}

func arr_truncate() {
	arr := [...]int{1, 2, 3, 4}
	fmt.Println(arr[1:3]) // index = 1, 2
	fmt.Println(arr[:len(arr)-1])
}

// https://go.dev/ref/spec#Slice_types
func slice_test() {
	var s0 []int
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0)) // 1 1
	s0 = append(s0, 2)
	fmt.Println(len(s0), cap(s0)) // 2 2
	s0 = append(s0, 3)
	fmt.Println(len(s0), cap(s0)) // 3 4
	s0 = append(s0, 4)
	fmt.Println(len(s0), cap(s0)) // 4 4
	s0 = append(s0, 5)
	fmt.Println(len(s0), cap(s0)) // 5 8

	for i := 0; i < 10; i++ {
		s0 = append(s0, i)
		fmt.Println(s0, len(s0), cap(s0))
	}
	fmt.Println("--------")
	s1 := s0[1:3] // 切片上的切片，capcity 计算方法要注意，就是从起始坐标一致到原始切片结束
	for _, ele := range s1 {
		fmt.Println(ele)
	}
	fmt.Println("--------")
	fmt.Println(s1, len(s1), cap(s1))
	s2 := s0[1:5]
	fmt.Println(s2, len(s2), cap(s2))
	s1[0] = 100
	s1[1] = 200
	fmt.Println(s2, len(s2), cap(s2))
}

func map_test() {
	var m1 map[string]int = make(map[string]int) // 使用 make 分配空间
	m1["hello"] = 1
	fmt.Println(m1)
	if v, ok := m1["hello"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key world do not exist")
	}
	m1["world"] = 2
	fmt.Println(m1)

	var m2 map[int]func(op int) int = make(map[int]func(int) int)
	m2[1] = func(op int) int {
		return op
	}
	m2[2] = func(op int) int {
		return op * op
	}
	m2[3] = func(op int) int {
		return op * op * op
	}
	fmt.Println(m2[1](2), m2[2](2), m2[3](2)) // 2 4 8
}

// use map simulate set
// 满足：相同值仅插入一次，set 的插入与删除其实都是排序的，按照 Map 使用 hash table
// 不好支持这个功能
func set_simulation() {
	var s map[int]bool = make(map[int]bool)
	s[1] = true
	s[2] = true

	if _, ok := s[1]; ok {
		fmt.Println("key 1 exists")
	} else {
		s[1] = true
	}
	delete(s, 1) // 删除 key 1
	// 很多功能不好模拟，使用时还是去搜现成的 package
}

// 只能牺牲插入速度来换可以排序
func slice_set_simulation() {
	var s []int
	s = append(s, 5)

	// 插入 10
	var insert_value = 10
	s = insert(insert_value, s)
	fmt.Println(s)

	// 插入 3
	insert_value = 3
	s = insert(insert_value, s)
	fmt.Println(s)

}

func insert(new_value int, s []int) []int {
	var insert_value = new_value
	var count = 0
	for _, ele := range s {
		if ele == insert_value {
			break
		} else {
			count++
		}
	}
	if count >= len(s) {
		s = append(s, insert_value)
	}
	sort.Ints(s)
	for _, ele := range s {
		fmt.Println(ele)
	}
	return s
}

// use slice simulate stack
func stack_simulation() {
	var s []int
	for i := 0; i < 10; i++ {
		s = pushs(i, s)
	}
	fmt.Println(s)
	var pop_value int
	pop_value, s = pops(s)
	fmt.Println("pop value is", pop_value, s)
	pop_value, s = pops(s)
	fmt.Println("pop value is", pop_value, s)
}

func pushs(push_value int, s []int) []int {
	s = append(s, push_value)
	return s
}

func pops(s []int) (int, []int) {
	if len(s) > 0 {
		var v = s[len(s)-1]
		s = s[:len(s)-1]
		return v, s
	}
	return -1, s // -1 表示没有值可 pop
}

// use slice simulate queue
// 先进先出
func queue_simulation() {
	var q []int
	for i := 0; i < 10; i++ {
		q = pushs(i, q)
	}
	fmt.Println(q)
	for i := 0; i < 10; i++ {
		q = pushq(i, q)
	}
	fmt.Println(q)
	rvalue, q := frontq(q)
	fmt.Println(rvalue, q)
	rvalue, q = frontq(q)
	fmt.Println(rvalue, q)
}

func pushq(v int, s []int) []int {
	s = append(s, v)
	return s
}

func frontq(s []int) (int, []int) {
	rvalue := s[0]
	s = s[1:]
	return rvalue, s
}
