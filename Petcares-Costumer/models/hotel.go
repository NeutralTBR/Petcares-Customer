package models

type Hotel struct {
	HotelID  string `json:"hotel_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Capacity int    `json:"capacity"`
}
