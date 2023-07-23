package factory

import (
	"disign_patterns/factory/product"
)

type Factory interface {
	CreateFruit() product.Fruit
}

type AppleFactory struct{}

func (af *AppleFactory) CreateFruit() product.Fruit {
	return new(product.Apple)
}

type PearFactory struct{}

func (pf *PearFactory) CreateFruit() product.Fruit {
	return new(product.Pear)
}

type BananaFactory struct{}

func (bf *BananaFactory) CreateFruit() product.Fruit {
	return new(product.Banana)
}
