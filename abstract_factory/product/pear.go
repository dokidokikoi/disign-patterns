package product

import "fmt"

type Pear interface {
	Show()
}

type ChinaPear struct{}

func (cp *ChinaPear) Show() {
	fmt.Println("i'm china pear!!")
}

type JapanPear struct{}

func (jp *JapanPear) Show() {
	fmt.Println("i'm japan pear!!")
}
