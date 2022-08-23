package factory

import "github.com/alemelomeza/poc-domain/domain/entity"

func NewWeighable(productNumber int, defaultQuantity float64) *entity.Weighable {
	if productNumber <= 0 {
		return &entity.Weighable{}
	}
	if defaultQuantity <= 0.000 {
		return &entity.Weighable{}
	}
	return &entity.Weighable{
		ProductNumber:   productNumber,
		DefaultQuantity: defaultQuantity,
	}
}
