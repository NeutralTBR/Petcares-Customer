package controllers

import (
	"database/sql"
	"net/http"
	"petcares/models"

	"github.com/gin-gonic/gin"
)

// GetAnimalsHandler retrieves and returns the animals for a specific customer
func GetAnimalsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")

		// Query to fetch animals for the given customer ID
		query := `
            SELECT animal_id, customer_id, animal_name, species, age, gender 
            FROM Animal
            WHERE customer_id = ?
        `
		rows, err := db.Query(query, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch animals"})
			return
		}
		defer rows.Close()

		var animals []models.Animal
		for rows.Next() {
			var animal models.Animal
			if err := rows.Scan(&animal.AnimalID, &animal.CustomerID, &animal.AnimalName, &animal.Species, &animal.Age, &animal.Gender); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan animal rows"})
				return
			}
			animals = append(animals, animal)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over animal rows"})
			return
		}

		// Render animals page with fetched data
		c.HTML(http.StatusOK, "animals.html", gin.H{
			"CustomerID": customerID,
			"Animals":    animals,
		})
	}
}

// AddAnimalHandler handles adding a new animal for a specific customer
func AddAnimalHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			customerID := c.Param("customer_id")
			animalName := c.PostForm("animal_name")
			species := c.PostForm("species")
			age := c.PostForm("age")
			gender := c.PostForm("gender")

			query := `
                INSERT INTO Animal (customer_id, animal_name, species, age, gender)
                VALUES (?, ?, ?, ?, ?)
            `
			_, err := db.Exec(query, customerID, animalName, species, age, gender)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add animal"})
				return
			}

			c.Redirect(http.StatusSeeOther, "/animals/"+customerID)
		} else {
			c.HTML(http.StatusOK, "add_animal.html", gin.H{
				"CustomerID": c.Param("customer_id"),
			})
		}
	}
}

// EditAnimalHandler handles editing an existing animal's details
func EditAnimalHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")
		animalID := c.Param("animal_id")

		if c.Request.Method == http.MethodPost {
			animalName := c.PostForm("animal_name")
			species := c.PostForm("species")
			age := c.PostForm("age")
			gender := c.PostForm("gender")

			query := `
                UPDATE Animal
                SET animal_name = ?, species = ?, age = ?, gender = ?
                WHERE animal_id = ? AND customer_id = ?
            `
			_, err := db.Exec(query, animalName, species, age, gender, animalID, customerID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit animal"})
				return
			}

			c.Redirect(http.StatusSeeOther, "/animals/"+customerID)
		} else {
			// Fetch the existing animal details to populate the form
			var animal models.Animal
			query := `
                SELECT animal_id, customer_id, animal_name, species, age, gender 
                FROM Animal
                WHERE animal_id = ? AND customer_id = ?
            `
			err := db.QueryRow(query, animalID, customerID).Scan(&animal.AnimalID, &animal.CustomerID, &animal.AnimalName, &animal.Species, &animal.Age, &animal.Gender)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Animal not found"})
				return
			}

			c.HTML(http.StatusOK, "edit_animal.html", gin.H{
				"CustomerID": customerID,
				"Animal":     animal,
			})
		}
	}
}

// DeleteAnimalHandler handles deleting an animal
func DeleteAnimalHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerID := c.Param("customer_id")
		animalID := c.Param("animal_id")

		query := `
            DELETE FROM Animal
            WHERE animal_id = ? AND customer_id = ?
        `
		_, err := db.Exec(query, animalID, customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete animal"})
			return
		}

		c.Redirect(http.StatusSeeOther, "/animals/"+customerID)
	}
}
