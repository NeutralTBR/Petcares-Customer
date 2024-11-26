// models/payment.go

package models

type Payment struct {
	PaymentID   string `json:"payment_id"`
	ReservasiID string `json:"reservasi_id"`
	StaffID     string `json:"staff_id"` // Assuming staff_id is a string, adjust as per your database schema
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
