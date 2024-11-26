package models

type Reservasi struct {
	ReservasiID     string `json:"reservasi_id"`
	AnimalID        string `json:"animal_id"`
	DokterID        string `json:"dokter_id"`
	RoomID          string `json:"room_id"`
	JenisReservasi  string `json:"jenis_reservasi"`
	HotelCheckin    string `json:"hotelcheckin"`
	HotelCheckout   string `json:"hotelcheckout"`
	WaktuKunjungan  string `json:"waktukunjungan"`
	Nomorantri      int    `json:"nomorantri"`
	FinishReservasi string `json:"finish_reservasi"`
}
