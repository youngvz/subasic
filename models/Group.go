package models


type Group struct {
	GroupID int `json:"id"`
	Name string `json:"name"`
	UserID int `json:"userId"`
	PhotoID *int `json:"photoId"`
}
