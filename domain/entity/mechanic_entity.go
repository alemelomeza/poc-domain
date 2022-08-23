package entity

const (
	beforAndNow     int = 1
	nXP             int = 2
	combo           int = 3
	whitoutMechanic int = 4
)

type Mechanic struct {
	CampaginID      string // DestacadoID
	MechanicType    int    // MechanicID
	ProductNumbers  []int
	OfferPrice      int
	OportunityPrice int
	ProductQuantity int
	UnitOfferPrice  float64
	Status          int // LoadStatusID
}

func (e *Mechanic) Validate() bool {
	switch e.MechanicType {
	case beforAndNow:
		if !validMechanicTypeBeforeAndNow(e.MechanicType, e.OfferPrice, e.OportunityPrice) {
			return false
		}
		if e.ProductQuantity > 0 {
			return false
		}
		if e.UnitOfferPrice > 0 {
			return false
		}
	case nXP:
		if !validMechanicTypeNXP(e.MechanicType, e.ProductQuantity, e.UnitOfferPrice) {
			return false
		}
		if e.OfferPrice > 0 {
			return false
		}
		if e.OportunityPrice > 0 {
			return false
		}
	default:
		return false
	}
	return true
}

func validMechanicTypeBeforeAndNow(mechanicType int, offerPrice int, oportunityPrice int) bool {
	if mechanicType == beforAndNow {
		return false
	}
	if offerPrice <= 0 {
		return false
	}
	if oportunityPrice < 0 {
		return false
	}
	if oportunityPrice > offerPrice {
		return false
	}
	return true
}

func validMechanicTypeNXP(mechanicType int, productQuanity int, unitOfferPrice float64) bool {
	if mechanicType == nXP {
		return false
	}
	if productQuanity <= 0 {
		return false
	}
	if unitOfferPrice < 1 {
		return false
	}
	return true
}
