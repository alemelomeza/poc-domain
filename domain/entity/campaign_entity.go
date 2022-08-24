package entity

import (
	"fmt"
	"strconv"
)

const (
	inProcessCampaign string = "En Proceso"
	draftCampaign     string = "Borrador"
	bloquedCampaign   string = "Bloqueado"
	inStoreCampaign   string = "En Tienda"
	filledCampaign    string = "Completado"
	deletedCampaign   string = "Eliminado"
	stopedCampaign    string = "Detenido"
)

type Campagin struct {
	ID        string
	Name      string
	LocalesID []int
	Status    string
	Icons     []string
	Mechanics map[string]Mechanic
}

func (e *Campagin) AddMechanic(mechanic Mechanic) error {
	if mechanic.CampaginID != e.ID {
		return fmt.Errorf("error can not add a mechanic with different id campaign")
	}
	hash := hashMechanic(mechanic.CampaginID, mechanic.MechanicType)
	e.Mechanics[hash] = mechanic
	return nil
}

func (e Campagin) Active() bool {
	return e.Status == inStoreCampaign
}

func (e Campagin) Inactive() bool {
	switch e.Status {
	case filledCampaign, deletedCampaign, stopedCampaign:
		return true
	}
	return false
}

func (e *Campagin) UpdateStatus(status string) error {
	e.Status = status
	return nil
}

func hashMechanic(campaignId string, mechanicType int) string {
	return fmt.Sprintf("%s-%s", campaignId, strconv.Itoa(mechanicType))
}
