package factory

import (
	"strings"

	"github.com/alemelomeza/poc-domain/domain/entity"
)

var (
	units = map[string]struct{}{
		"KG": {},
		"LT": {},
		"EA": {},
		"UN": {},
		"GR": {},
		"CC": {},
		"ML": {},
	}
)

func NewItem(itemNumber int, productNumber int, sellQuantity float64, sellUomCode string) *entity.Item {
	if itemNumber <= 0 {
		return &entity.Item{}
	}
	if productNumber <= 0 {
		return &entity.Item{}
	}
	if sellQuantity <= 0.000 {
		return &entity.Item{}
	}
	sellUomCode = strings.ToUpper(sellUomCode)
	if _, ok := units[sellUomCode]; !ok {
		return &entity.Item{}
	}
	return &entity.Item{
		ItemNumber:    itemNumber,
		ProductNumber: productNumber,
		SellQuantity:  sellQuantity,
		SellUomCode:   sellUomCode,
	}
}
