package service

import (
	"math"

	"github.com/alemelomeza/poc-domain/domain/entity"
)

func LowestPriceBeforeAndNow(campaign entity.Campagin) (int, int) {
	// TODO: think more in that
	if len(campaign.Mechanics) == 0 {
		return 0, 0
	}
	lowestSalePrice := math.MaxInt
	lowestTlmcPrice := math.MaxInt
	for _, mechanic := range campaign.Mechanics {
		if mechanic.MechanicType != entity.BeforAndNowMechanic {
			continue
		}
		if lowestSalePrice > mechanic.OfferPrice {
			lowestSalePrice = mechanic.MechanicType
		}
		if lowestTlmcPrice > mechanic.OportunityPrice {
			lowestTlmcPrice = mechanic.OportunityPrice
		}
	}
	if lowestSalePrice == math.MaxInt {
		lowestSalePrice = 0
	}
	if lowestTlmcPrice == math.MaxInt {
		lowestTlmcPrice = 0
	}
	return lowestSalePrice, lowestTlmcPrice
}

func LowestPriceNXP(campagin entity.Campagin) (int, int) {
	if len(campagin.Mechanics) == 0 {
		return 0, 0
	}
	lowestPackPriceXSize := math.MaxInt
	packSize := math.MaxInt
	for _, mechanic := range campagin.Mechanics {
		if mechanic.MechanicType != entity.NXPMechanic {
			continue
		}
		packPriceXSize := mechanic.UnitOfferPrice * float64(mechanic.ProductQuantity)
		if lowestPackPriceXSize > int(packPriceXSize) {
			lowestPackPriceXSize = int(packPriceXSize)
			packSize = mechanic.ProductQuantity
		}
	}
	if lowestPackPriceXSize == math.MaxInt {
		lowestPackPriceXSize = 0
		packSize = 0
	}
	return lowestPackPriceXSize, packSize
}
