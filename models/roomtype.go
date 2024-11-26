package models

type RoomType struct {
	TypeID        string  `json:"type_id"`
	Typename      string  `json:"typename"`
	Pricepernight float64 `json:"pricepernight"`
}
