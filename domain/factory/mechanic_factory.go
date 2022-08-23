package factory

import "github.com/alemelomeza/poc-domain/domain/entity"

func NewMechanic(campaignId string, mechanicType int, productNumbers []int, offerPrice int, oportunityPrice int, productQuanity int, unitOfferPrice float64, status int) *entity.Mechanic {
	if len(campaignId) < 24 {
		return &entity.Mechanic{}
	}
	if len(productNumbers) < 1 {
		return &entity.Mechanic{}
	}
	if status <= 0 {
		return &entity.Mechanic{}
	}
	if mechanicType <= 0 {
		return &entity.Mechanic{}
	}
	return &entity.Mechanic{
		CampaginID:      campaignId,
		MechanicType:    mechanicType,
		ProductNumbers:  productNumbers,
		OfferPrice:      offerPrice,
		OportunityPrice: oportunityPrice,
		ProductQuantity: productQuanity,
		UnitOfferPrice:  unitOfferPrice,
		Status:          status,
	}
}
