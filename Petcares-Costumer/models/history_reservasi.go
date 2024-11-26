package models

type HistoryReservasi struct {
	HistoriID   string `json:"histori_id"`
	ReservasiID string `json:"reservasi_id"`
	Description string `json:"description"`
}
