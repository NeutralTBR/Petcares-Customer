package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"petcares/models"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type EditProfileForm struct {
	FirstName  string `form:"firstName" binding:"required"`
	LastName   string `form:"lastName" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Phone      string `form:"phone" binding:"required"`
	Password   string `form:"password"`
	CustomerID string `form:"customer_id"`
}

// CheckCustomerCredentials queries the database for a customer with the given email and password
func CheckCustomerCredentials(db *sql.DB, email, password string) (models.Customer, error) {
	var customer models.Customer
	err := db.QueryRow("SELECT customer_id, firstname, lastname, email, phone, password FROM Customer WHERE email = ? AND password = ?", email, password).
		Scan(&customer.CustomerID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Phone, &customer.Password)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

// CustomerLoginHandler handles customer login requests
func CustomerLoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		customer, err := CheckCustomerCredentials(db, email, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		session := sessions.Default(c)
		session.Set("customerID", customer.CustomerID)
		session.Set("customerFirstName", customer.FirstName)
		session.Set("customerLastName", customer.LastName)
		session.Set("customerEmail", customer.Email)
		session.Set("customerPhone", customer.Phone)
		session.Save()

		custID := session.Get("customerID")
		custStr := fmt.Sprint(custID) // Konversi ke string
		c.Redirect(http.StatusFound, "/dashboard/"+custStr)

		// c.Set("customer", customer)

		// c.HTML(http.StatusOK, "indexcustomer.html", gin.H{
		// 	"CustomerID": customer.CustomerID,
		// 	"FirstName":  customer.FirstName,
		// 	"LastName":   customer.LastName,
		// 	"Email":      customer.Email,
		// 	"Phone":      customer.Phone,
		// })

		// c.HTML(http.StatusOK, "reservation.html", gin.H{"CustomerID": customer.CustomerID})

	}
}

func CustomerDashboardHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// customerID := c.Param("customer_id")
		session := sessions.Default(c)

		c.HTML(http.StatusOK, "indexcustomer.html", gin.H{
			"CustomerID": session.Get("customerID"),
			"FirstName":  session.Get("customerFirstName"),
			"LastName":   session.Get("customerLastName"),
			"Email":      session.Get("customerEmail"),
			"Phone":      session.Get("customerPhone")})

	}
}

// GetCustomerProfileHandler retrieves and returns the logged-in customer's profile information
func GetCustomerProfileHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customer, exists := c.Get("customer")
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer information not found"})
			return
		}

		c.JSON(http.StatusOK, customer)
	}
}

func HandleRegisterForm(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil data dari formulir
		if c.Request.Method == http.MethodPost {
			email := c.PostForm("email")
			firstName := c.PostForm("firstName")
			lastName := c.PostForm("lastName")
			phone := c.PostForm("phone")
			password := c.PostForm("password")

			// TODO: Validasi data formulir sesuai kebutuhan aplikasi Anda

			// Simpan data ke dalam database
			query := `
            INSERT INTO customer (email, firstname, lastname, phone, password)
            VALUES (?, ?, ?, ?, ?)
        `
			// Eksekusi query untuk menyimpan data
			result, err := db.Exec(query, email, firstName, lastName, phone, password)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "Failed to register user",
					"message": err.Error(), // Sertakan pesan error dari database jika ada
					"result":  result,
				})
				return
			}
			c.Redirect(http.StatusFound, "/")

		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"email": c.Param("email"),
			})
		}

		// Dapatkan ID pengguna yang baru saja dibuat jika perlu
		// userID, _ := result.LastInsertId()
		// c.Redirect(http.StatusFound, "/login")

		// // Berhasil mendaftar, arahkan ke halaman login atau berikan respons sukses
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "User registered successfully",
		// 	"email":   email,
		// 	"user_id": userID, // Sertakan ID pengguna jika diperlukan
		// })

	}
}
