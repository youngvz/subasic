package models

type Location struct {
	LocationId int `json:"locationId"`
	Address *string `json:"address,omitempty"`
	Coordinate Point `json:"coordinate"`
	GoogleId *string `json:"googleId,omitempty"`
	Name *string `json:"name,omitempty"`
}
