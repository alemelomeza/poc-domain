package factory

import "github.com/alemelomeza/poc-domain/domain/entity"

func NewPosition(storeNumber int, productNuber int, listPrice int, salePrice int, tlmcPrice int, packPrice int, packSize int, defaultQuantity float64, soldByWeight bool) *entity.Position {
	if storeNumber < 1 {
		return &entity.Position{}
	}
	if productNuber < 1 {
		return &entity.Position{}
	}
	if listPrice < 0 {
		return &entity.Position{}
	}
	if salePrice < 0 {
		return &entity.Position{}
	}
	if tlmcPrice < 0 {
		return &entity.Position{}
	}
	if packPrice < 0 {
		return &entity.Position{}
	}
	if packSize < 0 {
		return &entity.Position{}
	}
	if defaultQuantity < 0.000 {
		return &entity.Position{}
	}
	return &entity.Position{
		StoreNumber:     storeNumber,
		ProductNumber:   productNuber,
		ListPrice:       listPrice,
		SalePrice:       salePrice,
		TlmcPrice:       tlmcPrice,
		PackPrice:       packPrice,
		PackSize:        packSize,
		DefaultQuantity: defaultQuantity,
		SoldByWeight:    soldByWeight,
	}
}
