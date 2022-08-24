package factory

import "github.com/alemelomeza/poc-domain/domain/entity"

func NewCampaign(id string, name string, localesIds []int, status string, icons []string) *entity.Campagin {
	if len(id) < 24 {
		return &entity.Campagin{}
	}
	if len(localesIds) < 1 {
		return &entity.Campagin{}
	}
	if len(status) == 0 {
		return &entity.Campagin{}
	}
	return &entity.Campagin{
		ID:        id,
		Name:      name,
		LocalesID: localesIds,
		Status:    status,
		Icons:     icons,
	}
}
