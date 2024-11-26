package main

import (
	"database/sql"
	"fmt"
	"log"

	"petcares/routes" // Ensure this path is correct

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/databasepetcare")
	if err != nil {
		log.Fatal(err)
	} else {
		defer db.Close()
		fmt.Println("Connected to the database!")
	}

	router := gin.Default()

	// Mengatur session store
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Set up routes
	routes.SetupRoutes(router, db)

	// Serve the index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
