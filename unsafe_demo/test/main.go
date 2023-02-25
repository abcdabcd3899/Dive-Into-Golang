package main

import (
	"fmt"
	"unsafe"
)

/*
#include <stdio.h>
#include <math.h>  // call the sqrt function
#cgo LDFLAGS: -lm
int sum(int a, int b) {
	return a + b;
}

int add(int a, int b) {
	return a + b;
}

*/
import "C" // 这一行和上面代码不能有空行

/*
Unsafe programming in Go is used in specific cases where it is necessary to bypass the type safety and memory safety guarantees provided by the language. Some common cases where unsafe programming might be necessary include:

Interoperability with C code: In some cases, you may need to use C libraries in Go, and unsafe programming can be used to interface with C code.

Low-level system programming: Unsafe programming can be used for direct memory manipulation or accessing hardware resources that are not accessible through the standard Go libraries.

Performance optimization: Unsafe programming can be used to write faster and more efficient code, but it's important to note that it should only be used as a last resort after profiling and optimizing your code using other techniques.

In general, it is recommended to avoid using unsafe programming unless it is absolutely necessary, as it can introduce bugs and security vulnerabilities into your code.
*/

func main() {
	// 1. go 语言没有强制类型转换，当使用 unsafe 之后就能随意转换了
	var a int = 10
	b := (unsafe.Pointer(&a))
	c := *(*float64)(b)
	fmt.Println(a, b, c)

	// 调用 C 函数
	a1, b1 := 3, 4
	c1 := int(C.add(C.int(a1), C.int(b1)))
	fmt.Println("The sum of", a1, "and", b1, "is", c1)

	fmt.Println(C.sum(5, 6))

	a2 := 64
	fmt.Println(C.sqrt(C.double(a2)))
}
