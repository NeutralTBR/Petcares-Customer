package models

type Room struct {
	RoomID    string `json:"room_id"`
	TypeID    string `json:"type_id"`
	RoomName  string `json:"room_name"`
	Deskripsi string `json:"deskripsi"`
	Available string `json:"available"`
}
