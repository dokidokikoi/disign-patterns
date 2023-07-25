package factory

import "disign_patterns/abstract_factory/product"

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() product.Apple {
	return new(product.ChinaApple)
}

func (cf *ChinaFactory) CreatePear() product.Pear {
	return new(product.ChinaPear)
}

func (cf *ChinaFactory) CreateBanana() product.Banana {
	return new(product.ChinaBanana)
}
