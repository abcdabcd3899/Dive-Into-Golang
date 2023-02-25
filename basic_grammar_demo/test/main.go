package main // 必须是 main 包才能被编译执行

import (
	"fmt"
	"os"
)

const (
	MONDAY = iota + 1 // iota = 0
	TUSDAY            // iota = 1, 隐含了 iota + 1 = 2
	THU               // iota = 2, 隐含了 iota + 1 = 3
)

// go run main.go dive-into-golang
func main() {
	a := 10
	fmt.Println(fib(a))
	exchange_variable_1()
	exchange_variable_2()
	fmt.Println(MONDAY, " ", TUSDAY, " ", THU)
	type_system()
	pointer()
	string_test()
	arr()
	bit_clear()
	loop()
	switch_condition()
	os.Exit(0)
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func exchange_variable_1() {
	a := 10
	b := 20
	tmp := a
	a = b
	b = tmp
	fmt.Println(a, " ", b)
}

func exchange_variable_2() {
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, " ", b)
}

func type_system() {
	// 不支持隐式类型转换
	var (
		a int = 10
		b int64
	) // 和下面声明效果一样，但是 golang 推荐要合并变量的声明
	// var a int = 10
	// var b int64
	// b = a  // error
	b = int64(a) + 1
	fmt.Println(a, b)
}

func pointer() {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1
	fmt.Println(a, aPtr)
	fmt.Printf("%T, %T\n", a, aPtr)

}

func string_test() {
	var s string
	if s == "" {
		fmt.Println("string is nil")
	}
}

func arr() {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 3, 2}
	c := [...]int{4, 5, 6}
	d := [...]int{1, 2, 3}
	e := [3]int{}
	fmt.Println(e)
	e = d
	f := 0
	f = 1
	fmt.Println(a == b, a == c, a == d, a == e, f) // false  false true
}

// &^ 运算符的使用
func bit_clear() {
	a := 6            //0110
	b := 11           //1011
	c := a &^ b       // a &(^b) = 0110 &  0100 = 0100 = 4
	d := b &^ a       // b &(^a) = 1011 & 1001 = 1001 = 9
	fmt.Println(c, d) // 4 9
}

// go 只有 25 个关键字，只支持 for 循环
func loop() {
	a := 0
	for a < 5 {
		fmt.Println(a)
		a++
	}

	a = 0
	for {
		fmt.Println(a)
		a++
		if a == 5 {
			break
		}
	}
}

func switch_condition() {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2, 4:
			fmt.Println("Even")
		case 1, 3:
			fmt.Println("Odd")
		default:
			fmt.Println("It is not 0-5")
		}
	}
	fmt.Println("--------")
	for i := 0; i < 5; i++ {
		switch {
		case i == 0 || i == 2 || i == 4:
			fmt.Println("Even")
		case i == 1 || i == 3:
			fmt.Println("Odd")
		default:
			fmt.Println("It is not 0-5")
		}
	}
	fmt.Println("--------")
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			fmt.Println("Even")
		case i%2 != 0:
			fmt.Println("Odd")
		default:
			fmt.Println("It is not 0-5")
		}
	}
}
