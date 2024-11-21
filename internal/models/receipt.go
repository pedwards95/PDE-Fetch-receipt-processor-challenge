package models

import "github.com/google/uuid"

// Receipt ...
type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []*Item `json:"items"`
	Total        string  `json:"total"`

	id uuid.UUID
}

// GetID getter
func (r *Receipt) GetID() uuid.UUID {
	return r.id
}

// SetID setter
func (r *Receipt) SetID(id uuid.UUID) bool {
	if r.id == uuid.Nil {
		r.id = id
		return true
	}
	return false
}
