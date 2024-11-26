package models

type Dokter struct {
	DokterID   string `json:"dokter_id"`
	DokterName string `json:"dokter_name"`
	Specialize string `json:"specialize"`
	Email      string `json:"email"`
	NoPhone    string `json:"nophone"`
	Password   string `json:"password"`
}
