package entity

// Mechanic Types
const (
	BeforAndNowMechanic int = 1
	NXPMechanic         int = 2
	ComboMechanic       int = 3
	WhitoutMechanic     int = 4
)

// Mechanic statuses
const (
	ApprovedMechanic int = 1 // Valid mechanic
	ForcedMechanic   int = 3 // Valid mechanic
	RejectedMechanic int = 2 // Invalid mechanic
	DeletedMechanic  int = 4 // Invalid mechanic
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
	notApproved := e.Status != ApprovedMechanic
	notForced := e.Status != ForcedMechanic
	if notApproved || notForced {
		return false
	}
	switch e.MechanicType {
	case BeforAndNowMechanic:
		if !validMechanicTypeBeforeAndNow(e.MechanicType, e.OfferPrice, e.OportunityPrice) {
			return false
		}
		if e.ProductQuantity > 0 {
			return false
		}
		if e.UnitOfferPrice > 0 {
			return false
		}
	case NXPMechanic:
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
	if mechanicType == BeforAndNowMechanic {
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
	if mechanicType == NXPMechanic {
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
