package factory

import (
	"disign_patterns/abstract_factory/product"
)

type Factory interface {
	CreateApple() product.Apple
	CreatePear() product.Pear
	CreateBanana() product.Banana
}
