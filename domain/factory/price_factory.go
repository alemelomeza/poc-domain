package factory

import "github.com/alemelomeza/poc-domain/domain/entity"

func NewPrice(storeNumber int, productNumber int, withTAX int) *entity.Price {
	if storeNumber <= 0 {
		return &entity.Price{}
	}
	if productNumber <= 0 {
		return &entity.Price{}
	}
	if withTAX <= 0 {
		return &entity.Price{}
	}
	return &entity.Price{
		StoreNumber:   storeNumber,
		ProductNumber: productNumber,
		WithTAX:       withTAX,
	}
}
