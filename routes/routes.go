package routes

import (
	"database/sql"
	"net/http"
	"petcares/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/index", controllers.CustomerLoginHandler(db))
	router.POST("/register", controllers.HandleRegisterForm(db)) // Add this line
	router.GET("/dashboard/:customer_id", controllers.CustomerDashboardHandler(db))
	router.GET("/profil", controllers.GetCustomerProfileHandler(db))
	router.GET("/profil/:customer_id", controllers.ViewProfileHandler(db))
	router.GET("/profile/edit/:customer_id", controllers.EditProfileHandler(db))
	router.POST("/profile/edit/:customer_id", controllers.UpdateProfileHandler(db))
	// Routes for other functionalities
	router.GET("/animals/:customer_id", controllers.GetAnimalsHandler(db))
	router.POST("/add_animal/:customer_id", controllers.AddAnimalHandler(db))
	router.GET("/add_animal/:customer_id", controllers.AddAnimalHandler(db))
	router.GET("/edit_animal/:customer_id/:animal_id", controllers.EditAnimalHandler(db))
	router.POST("/edit_animal/:customer_id/:animal_id", controllers.EditAnimalHandler(db))
	router.GET("/delete_animal/:customer_id/:animal_id", controllers.DeleteAnimalHandler(db))

	router.GET("/reservation/:customer_id", controllers.ReservationHandler(db))
	router.GET("/reservation/hotel/:customer_id", controllers.ReservationHotelHandler(db))
	router.POST("/reservation/hotel/:customer_id", controllers.SubmitReservationHotelHandler(db))
	router.GET("/reservation/doctor/:customer_id", controllers.ReservationDoctorHandler(db))
	router.POST("/reservation/doctor/:customer_id", controllers.SubmitReservationDoctorHandler(db))

	router.GET("/payment/:customer_id", controllers.ViewPaymentsHandler(db))
	// Handle Not Found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
	})
}

func RegisterLoginRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", controllers.CustomerLoginHandler(db))
}

func RegisterAnimalRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/animals/:customer_id", controllers.GetAnimalsHandler(db))
	router.POST("/animals/:customer_id", controllers.AddAnimalHandler(db))
	router.PUT("/animals/:animal_id/:customer_id", controllers.EditAnimalHandler(db))
	router.DELETE("/animals/:animal_id/:customer_id", controllers.DeleteAnimalHandler(db))
}

func RegisterAPIRoutes(router *gin.Engine, db *sql.DB) {
	// Define your API routes here if needed
}
