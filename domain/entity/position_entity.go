package entity

import (
	"fmt"
	"strconv"
)

type Position struct {
	StoreNumber           int
	ProductNumber         int
	ListPrice             int
	SalePrice             int
	TlmcPrice             int
	PackPrice             int
	PackSize              int
	DefaultQuantity       float64
	SoldByWeight          bool
	PricePerUnitOfMessure string
	Campagins             map[string]Campagin
}

func (e *Position) AddCampagin(campagin Campagin) error {
	if !campagin.Active() {
		return fmt.Errorf("error can not add a inactived campagin")
	}
	if len(campagin.Mechanics) < 1 {
		return fmt.Errorf("error can not add a campagin without mechanics")
	}
	e.Campagins[campagin.ID] = campagin
	return nil
}

func (e *Position) RemoveCampaign(campagin Campagin) error {
	if !campagin.Inactive() {
		return fmt.Errorf("error can not remove a inactive campagin")
	}
	delete(e.Campagins, campagin.ID)
	return nil
}

func (e *Position) UpdateListPrice(price Price) error {
	e.ListPrice = price.WithTAX
	e.updateSalePrice()
	return nil
}

func (e *Position) updateSalePrice() {
	if e.SalePrice <= 0 {
		e.SalePrice = e.ListPrice
		e.updatePricePerUnitOfMessure()
	}
}

func (e *Position) updatePricePerUnitOfMessure() {
	e.PricePerUnitOfMessure = fmt.Sprintf("PPUM $%s", strconv.Itoa(e.SalePrice))
}
