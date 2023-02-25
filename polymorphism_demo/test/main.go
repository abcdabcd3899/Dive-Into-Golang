package main

import (
	"fmt"
)

// func main() {
// 	// g := GoProgrammer{}
// 	g := new(GoProgrammer)
// 	ret := PI(g)
// 	fmt.Println("--------", ret)
// 	// r := RustProgrammer{}
// 	r := new(RustProgrammer)
// 	ret = PI(r)
// 	fmt.Println("--------", ret)
// }

// type Programmer interface {
// 	WriteHelloWorld() string
// }

// func PI(p Programmer) string {
// 	return p.WriteHelloWorld()
// }

// type GoProgrammer struct {
// }

// func (g GoProgrammer) WriteHelloWorld() string {
// 	fmt.Println("Golang")
// 	return "hello, golang"
// }

// type RustProgrammer struct {
// }

// func (r RustProgrammer) WriteHelloWorld() string {
// 	fmt.Println("Rust")
// 	return "hello, rust"
// }

// func main() {

// 	// 下面四行，我们可以使用 GoProgrammer 的值或者指针调用 GoProgrammer 的方法
// 	// 这和普通函数的调用原则一样
// 	g := new(GoProgrammer)
// 	g.WriteHelloWorld()
// 	// g := GoProgrammer{}
// 	// fmt.Println(g.WriteHelloWorld())

// 	// PI 是一个函数，且它是 GoProgrammer 和 RustProgrammer 的接口
// 	// 因此 PI 实现了多态性，但是由于 p Programmer 无法知道到底其实现了接口方法的类到底是使用
// 	// 值接收者还是指针接收者，因此这里传递给 PI 的参数遵循规则如下：
// 	// 1. 如果方法接收者是值，那么既可以传递指针对象，也可以传递值对象
// 	// 2. 如果方法接收者是指针，那么只能传递指针对象
// 	// 这是 golang 实现多态性最大的不同
// 	ret := PI(g)
// 	fmt.Println("--------", ret)
// 	r := new(RustProgrammer)
// 	ret = PI(r)
// 	fmt.Println("--------", ret)
// }

// type Programmer interface {
// 	WriteHelloWorld() string
// }

// // 这里 p 不能定义指针 p *Programmer
// // PI 是一个函数
// func PI(p Programmer) string {
// 	return p.WriteHelloWorld()
// }

// type GoProgrammer struct {
// }

// // 这些方法是针对指针实现的
// // (1)一个函数增加了接受者之后，就变成了类的方法，接收者分为：值接收者和指针接收者
// func (g *GoProgrammer) WriteHelloWorld() string {
// 	fmt.Println("Golang")
// 	return "hello, golang"
// }

// type RustProgrammer struct {
// }

// // (1)一个函数增加了接受者之后，就变成了类的方法，接收者分为：值接收者和指针接收者
// func (r *RustProgrammer) WriteHelloWorld() string {
// 	fmt.Println("Rust")
// 	return "hello, rust"
// }

// 接口继承，golang 灵活但是其仍然图灵完备
// 1. 接口实现
type FileAttrs interface {
	getPath() string
}

type File struct {
	Name string
}

// File 实现了 Duck Type Interface FileAttrs
// 这个方法在编译后实际上是 "File + getPath"
func (file File) getPath() string {
	return fmt.Sprintf("%s/%T", file.Name, file)
}

func printFileAttrs(file FileAttrs) {
	fmt.Printf("file name: %s ? %T\n", file.getPath(), file)
}

// 2. 组合优于继承
// 并且因为 File 实现了 getPath 方法，由于组合存在，实际上 BlobFile 和 TextFile 都实现了 getPath方法
type BlobFile struct {
	File // 组合，组合后这些 struct 都拥有了 File 的方法
}

// 使用 BlobFile 的对象调用
// 这个方法在编译后实际上是 "BlobFile + getPath"
// func (file *BlobFile) getPath() string {
// 	return fmt.Sprintf("%s/%T BolbFile", file.Name, file)
// }

type TextFile struct {
	File // 组合
}

// 使用 TextFile 的对象调用
// 这个方法在编译后实际上是 "TextFile + getPath"
// 这三个 getPath 在 golang 内部看来是不同的
// func (file *TextFile) getPath() string {
// 	return fmt.Sprintf("%s/%T TextFile", file.Name, file)
// }

func main() {

	file := &File{Name: "0"}
	bf := &BlobFile{File{Name: "1"}}
	tf := &TextFile{File{Name: "2"}}
	printFileAttrs(file)
	printFileAttrs(bf)
	printFileAttrs(tf)
	fmt.Println(bf.getPath())
	fmt.Println(tf.getPath())
}
