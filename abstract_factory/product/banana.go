package product

import "fmt"

type Banana interface {
	Show()
}

type ChinaBanana struct{}

func (cb *ChinaBanana) Show() {
	fmt.Println("i'm china banana!!")
}

type JapanBanana struct{}

func (jb *JapanBanana) Show() {
	fmt.Println("i'm japan banana!!")
}
