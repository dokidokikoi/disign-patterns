package main

import (
	"disign_patterns/abstract_factory/factory"
	"disign_patterns/abstract_factory/product"
)

func main() {
	var chinaFactory factory.Factory
	var japanFactory factory.Factory

	chinaFactory = &factory.ChinaFactory{}
	japanFactory = &factory.JapanFactory{}

	showAllFruits(chinaFactory)
	showAllFruits(japanFactory)
}

func showAllFruits(f factory.Factory) {
	var apple product.Apple
	var pear product.Pear
	var banana product.Banana

	apple = f.CreateApple()
	apple.Show()
	pear = f.CreatePear()
	pear.Show()
	banana = f.CreateBanana()
	banana.Show()
}
