package factory

import "disign_patterns/abstract_factory/product"

type JapanFactory struct{}

func (jf *JapanFactory) CreateApple() product.Apple {
	return new(product.JapanApple)
}

func (jf *JapanFactory) CreatePear() product.Pear {
	return new(product.JapanPear)
}

func (jf *JapanFactory) CreateBanana() product.Banana {
	return new(product.JapanBanana)
}
