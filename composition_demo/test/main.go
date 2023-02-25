package main

import "fmt"

func main() {
	// d := new(Dog)
	d := Dog{}
	d.Speak()
	d.SpeakTo("wangwang")
}

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 组合优于继承，在其他语言中也认为继承是万恶之源，因此
// 在 golang 中直接不支持继承
type Dog struct {
	Pet // 组合了匿名字段
}

// 我们说，如果上面这是 golang 中的继承，那么继承就能 override 父类的方法
// 全部 override 父类所有的方法才叫继承
func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("Wang!", host)
}
