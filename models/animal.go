package models

type Animal struct {
	AnimalID   string `json:"animal_id"`
	CustomerID string `json:"customer_id"`
	AnimalName string `json:"animal_name"`
	Species    string `json:"species"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
}
