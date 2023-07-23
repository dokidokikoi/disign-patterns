package main

import (
	"disign_patterns/factory/factory"
	"disign_patterns/factory/product"
)

func main() {
	var fruit product.Fruit
	var fruitFactory factory.Factory

	fruitFactory = &factory.AppleFactory{}
	fruit = fruitFactory.CreateFruit()
	fruit.Show()

	fruitFactory = &factory.PearFactory{}
	fruit = fruitFactory.CreateFruit()
	fruit.Show()

	fruitFactory = &factory.BananaFactory{}
	fruit = fruitFactory.CreateFruit()
	fruit.Show()
}
