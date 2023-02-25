package main

import (
	"fmt"
)

type Pet struct {
	speakHandle   func()
	speakToHandle func(string)
}

func (p *Pet) Speak() {
	if p.speakHandle == nil {
		p.speakHandle = func() {
			fmt.Print("...")
		}
	}
	p.speakHandle()
}

func (p *Pet) SpeakTo(host string) {
	if p.speakToHandle == nil {
		p.speakToHandle = func(_host string) {
			p.Speak()
			fmt.Println(_host)
		}
	}
	p.speakToHandle(host)
}

type Dog struct {
	Pet
}

func (d *Dog) Ctor() *Dog {
	overrideSpeakHandle(d)
	overrideSpeakToHandle(d)
	return d
}

func overrideSpeakHandle(d *Dog) {
	if d.speakHandle == nil {
		d.speakHandle = func() {
			fmt.Print("Wang!")
		}
	}
}

func overrideSpeakToHandle(d *Dog) {
	if d.speakToHandle == nil {
		d.speakToHandle = func(_host string) {
			d.Speak()
			fmt.Println(_host + "...")
		}
	}
}

func main() {
	pet := Pet{}
	dog := new(Dog).Ctor()
	pet.SpeakTo("Pet")
	dog.SpeakTo("Dog")
	pet.speakToHandle("Pet")
	dog.speakToHandle("Dog")
}
