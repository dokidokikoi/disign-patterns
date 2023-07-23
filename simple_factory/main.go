package main

import "fmt"

// -------------------- product
type fruit interface {
	Show()
}

type apple struct{}

func (a *apple) Show() {
	fmt.Println("i'm apple!!")
}

type pear struct{}

func (p *pear) Show() {
	fmt.Println("i'm pear!!")
}

type banana struct{}

func (b *banana) Show() {
	fmt.Println("i'm banana!!")
}

// ---------------------- factory
type factory struct{}

func (f *factory) CreateFruit(kind string) fruit {
	if kind == "apple" {
		return new(apple)
	} else if kind == "pear" {
		return new(pear)
	} else if kind == "banana" {
		return new(banana)
	}

	return nil
}

func main() {
	f := factory{}

	apple := f.CreateFruit("apple")
	apple.Show()
	pear := f.CreateFruit("pear")
	pear.Show()
	banana := f.CreateFruit("banana")
	banana.Show()
}
