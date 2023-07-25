package product

import "fmt"

type Apple interface {
	Show()
}

type ChinaApple struct{}

func (ca *ChinaApple) Show() {
	fmt.Println("i'm china apple!!")
}

type JapanApple struct{}

func (ja *JapanApple) Show() {
	fmt.Println("i'm japan apple!!")
}
