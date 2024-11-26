package controllers

import (
	"database/sql"
	"net/http"
	"petcares/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ReservationHandler renders the reservation selection page
func ReservationHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		customerID := session.Get("customerID").(string)

		// Check if customer ID is provided
		if customerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
			return
		}

		// You might want to validate if the customer ID exists in the database
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Customer WHERE customer_id = ?)", customerID).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
			return
		}
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}

		// Fetch reservations
		query := `
		SELECT reservasi_id, jenis_reservasi, animal_id, room_id, dokter_id, hotelcheckin, hotelcheckout, waktukunjungan, finishreservasi
		FROM Reservasi
		WHERE customer_id = ?
	`
		rows, err := db.Query(query, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
			return
		}
		defer rows.Close()

		var reservations []models.Reservasi
		for rows.Next() {
			var reservation models.Reservasi

			// Define sql.NullString variables for nullable string columns
			var reservasiID, jenisReservasi, animalID, roomID, dokterID, hotelCheckin, hotelCheckout, waktuKunjungan, finishReservasi sql.NullString

			// Scan into sql.NullString variables
			err := rows.Scan(
				&reservasiID,
				&jenisReservasi,
				&animalID,
				&roomID,
				&dokterID,
				&hotelCheckin,
				&hotelCheckout,
				&waktuKunjungan,
				&finishReservasi,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan reservation rows"})
				return
			}

			// Assign sql.NullString values to reservation struct fields
			if reservasiID.Valid {
				reservation.ReservasiID = reservasiID.String
			} else {
				reservation.ReservasiID = "" // or any default value as per your logic
			}

			if jenisReservasi.Valid {
				reservation.JenisReservasi = jenisReservasi.String
			} else {
				reservation.JenisReservasi = "" // or any default value
			}

			if animalID.Valid {
				reservation.AnimalID = animalID.String
			} else {
				reservation.AnimalID = "" // or any default value
			}

			if roomID.Valid {
				reservation.RoomID = roomID.String
			} else {
				reservation.RoomID = "" // or any default value
			}

			if dokterID.Valid {
				reservation.DokterID = dokterID.String
			} else {
				reservation.DokterID = "" // or any default value
			}

			if hotelCheckin.Valid {
				reservation.HotelCheckin = hotelCheckin.String
			} else {
				reservation.HotelCheckin = "" // or any default value
			}

			if hotelCheckout.Valid {
				reservation.HotelCheckout = hotelCheckout.String
			} else {
				reservation.HotelCheckout = "" // or any default value
			}

			if waktuKunjungan.Valid {
				reservation.WaktuKunjungan = waktuKunjungan.String
			} else {
				reservation.WaktuKunjungan = "" // or any default value
			}

			if finishReservasi.Valid {
				reservation.FinishReservasi = finishReservasi.String
			} else {
				reservation.FinishReservasi = "" // or any default value
			}

			// Append to reservations slice
			reservations = append(reservations, reservation)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over reservation rows"})
			return
		}

		// Render the reservation page with customer ID and reservations
		c.HTML(http.StatusOK, "reservation.html", gin.H{
			"CustomerID":   customerID,
			"Reservations": reservations,
		})
	}
}

// ReservationHotelHandler renders the pet hotel reservation page with the user's animals
func ReservationHotelHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		customerID := session.Get("customerID").(string)

		// Pastikan customerID valid
		if customerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
			return
		}

		// Query untuk mendapatkan data hewan
		animalQuery := `
            SELECT animal_id, animal_name, species
            FROM Animal
            WHERE customer_id = ?
        `
		animalRows, err := db.Query(animalQuery, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch animals"})
			return
		}
		defer animalRows.Close()

		var animals []models.Animal
		for animalRows.Next() {
			var animal models.Animal
			if err := animalRows.Scan(&animal.AnimalID, &animal.AnimalName, &animal.Species); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan animal rows"})
				return
			}
			animals = append(animals, animal)
		}

		if err := animalRows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over animal rows"})
			return
		}

		// Query untuk mendapatkan data kamar yang available
		roomQuery := `
            SELECT room_id, type_id, room_name, deskripsi, available
            FROM Room
            WHERE available = 'yes'
        `
		roomRows, err := db.Query(roomQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available rooms"})
			return
		}
		defer roomRows.Close()

		var rooms []models.Room
		for roomRows.Next() {
			var room models.Room
			if err := roomRows.Scan(&room.RoomID, &room.TypeID, &room.RoomName, &room.Deskripsi, &room.Available); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan room rows"})
				return
			}
			rooms = append(rooms, room)
		}

		if err := roomRows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over room rows"})
			return
		}

		// Mengirimkan data ke template HTML
		c.HTML(http.StatusOK, "reservation_hotel.html", gin.H{
			"CustomerID": customerID,
			"Animals":    animals,
			"Rooms":      rooms,
		})
	}
}

// SubmitReservationHotelHandler handles the form submission for pet hotel reservation
func SubmitReservationHotelHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		customerID := session.Get("customerID").(string)

		animalID := c.PostForm("animal_id")
		roomID := c.PostForm("room_id")
		checkin := c.PostForm("hotelcheckin")
		checkout := c.PostForm("hotelcheckout")

		reservasiID := uuid.New().String()

		query := `
            INSERT INTO Reservasi (customer_id, animal_id, room_id, jenis_reservasi, hotelcheckin, hotelcheckout, finishreservasi)
            VALUES (?, ?, ?, 'Hotel', ?, ?, 'No')
        `
		_, err := db.Exec(query, customerID, animalID, roomID, checkin, checkout)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Reservasi berhasil dilakukan", "reservasi_id": reservasiID})
		c.HTML(http.StatusOK, "edit_animal.html", gin.H{"CustomerID": customerID})
	}
}

// ReservationDoctorHandler renders the doctor reservation page
func ReservationDoctorHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		customerID := session.Get("customerID").(string)

		// Fetch animals
		queryAnimals := `
            SELECT animal_id, animal_name, species
            FROM Animal
            WHERE customer_id = ?
        `
		rows, err := db.Query(queryAnimals, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch animals"})
			return
		}
		defer rows.Close()

		var animals []models.Animal
		for rows.Next() {
			var animal models.Animal
			if err := rows.Scan(&animal.AnimalID, &animal.AnimalName, &animal.Species); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan animal rows"})
				return
			}
			animals = append(animals, animal)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over animal rows"})
			return
		}

		// Fetch doctors
		queryDoctors := `
            SELECT dokter_id, dokter_name
            FROM Dokter
        `
		rows, err = db.Query(queryDoctors)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doctors"})
			return
		}
		defer rows.Close()

		var doctors []models.Dokter
		for rows.Next() {
			var doctor models.Dokter
			if err := rows.Scan(&doctor.DokterID, &doctor.DokterName); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan doctor rows"})
				return
			}
			doctors = append(doctors, doctor)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over doctor rows"})
			return
		}

		c.HTML(http.StatusOK, "reservation_doctor.html", gin.H{
			"Animals":    animals,
			"Doctors":    doctors,
			"CustomerID": customerID,
		})
	}
}

// SubmitReservationDoctorHandler handles the form submission for doctor reservation
func SubmitReservationDoctorHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		customerID := session.Get("customerID").(string)
		animalID := c.PostForm("animal_id")
		dokterID := c.PostForm("dokter_id")
		waktuKunjungan := c.PostForm("waktukunjungan")

		reservasiID := uuid.New().String()

		query := `
            INSERT INTO Reservasi (customer_id, animal_id, dokter_id, jenis_reservasi, waktukunjungan, finish_reservasi)
            VALUES (?, ?, ?, 'Dokter', ?, 'No')
        `
		_, err := db.Exec(query, customerID, animalID, dokterID, waktuKunjungan)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Reservasi berhasil dilakukan", "reservasi_id": reservasiID})
	}
}
