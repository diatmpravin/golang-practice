package main

import "fmt"

type Foo struct {
	name string
}

func (f *Foo) SetName(name string) {
	f.name = name
}

func (f Foo) GetName() string {
	return f.name
}

func main() {
	p := new(Foo)
	p.SetName("Abc")
	name := p.GetName()
	fmt.Println(name)
}
