// controllers/controllers.go

package controllers

import (
	"database/sql"
	"net/http"
	"petcares/models"

	"github.com/gin-gonic/gin"
)

// func EditProfileHandler(db *sql.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		customerID := c.Param("customer_id")

// 		if c.Request.Method == http.MethodPost {
// 			// Handle form submission for profile update
// 			firstName := c.PostForm("first_name")
// 			lastName := c.PostForm("last_name")
// 			email := c.PostForm("email")
// 			phone := c.PostForm("phone")

// 			// Update query
// 			query := `
//                 UPDATE Customer
//                 SET first_name = ?, last_name = ?, email = ?, phone = ?
//                 WHERE customer_id = ?
//             `
// 			_, err := db.Exec(query, firstName, lastName, email, phone, customerID)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
// 				return
// 			}

// 			// Redirect to profile page after successful update
// 			c.Redirect(http.StatusSeeOther, "/profil") // Assuming /profil is your profile page URL
// 		} else {
// 			// Fetch existing customer details to populate the form
// 			var customer models.Customer
// 			query := `
//                 SELECT customer_id, first_name, last_name, email, phone
//                 FROM Customer
//                 WHERE customer_id = ?
//             `
// 			err := db.QueryRow(query, customerID).Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone)
// 			if err != nil {
// 				c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
// 				return
// 			}

//				// Render edit profile form with current customer data
//				c.HTML(http.StatusOK, "edit_profile.html", gin.H{
//					"Customer": customer,
//				})
//			}
//		}
//	}
func ViewProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")

		var customer models.Customer
		query := "SELECT customer_id, firstname, lastname, email, phone FROM Customer WHERE customer_id = ?"
		err := db.QueryRow(query, customerID).Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}

		// Render profile template with customer data
		c.HTML(http.StatusOK, "profil.html", gin.H{
			"Customer": customer,
		})
	}
}
func EditProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")

		var customer models.Customer
		query := "SELECT customer_id, firstname, lastname, email, phone FROM Customer WHERE customer_id = ?"
		err := db.QueryRow(query, customerID).Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
			return
		}

		c.HTML(http.StatusOK, "edit_profile.html", gin.H{
			"Customer": customer,
		})
	}
}

// Handler untuk menghandle update profil
func UpdateProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")

		// Ambil data dari form yang di-submit
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		email := c.PostForm("email")
		phone := c.PostForm("phone")

		// Update data di database
		query := "UPDATE Customer SET firstname = ?, lastname = ?, email = ?, phone = ? WHERE customer_id = ?"
		_, err := db.Exec(query, firstName, lastName, email, phone, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer profile"})
			return
		}

		// Redirect kembali ke halaman profil setelah berhasil update
		c.Redirect(http.StatusSeeOther, "/profile/"+customerID)
	}
}

func ViewPaymentsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")

		// Query untuk mengambil pembayaran dan tanggal checkin atau waktu kunjungan jika checkin null
		query := `
			SELECT p.amount, COALESCE(r.hotelcheckin, r.waktukunjungan) AS checkin_date
			FROM Payment p
			JOIN Reservasi r ON p.reservasi_id = r.reservasi_id
			JOIN Animal a ON r.animal_id = a.animal_id
			WHERE a.customer_id = ?
		`
		rows, err := db.Query(query, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
			return
		}
		defer rows.Close()

		var payments []struct {
			Amount      int
			CheckinDate string
		}

		for rows.Next() {
			var payment struct {
				Amount      int
				CheckinDate string
			}
			if err := rows.Scan(&payment.Amount, &payment.CheckinDate); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan payments"})
				return
			}
			payments = append(payments, payment)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over payments"})
			return
		}

		c.HTML(http.StatusOK, "view_payments.html", gin.H{
			"Payments": payments,
		})
	}
}
