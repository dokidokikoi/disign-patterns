package product

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("i'm apple!!")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("i'm pear!!")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("i'm banana!!")
}
