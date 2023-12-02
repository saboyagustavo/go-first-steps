package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func (c Product) GetFullInfo() string {
	return fmt.Sprintf(
		"Name: %s,\nDescription: %s,\nPrice: %d,\n",
		c.Name, c.Description, c.Price,
	)
}

func home(w http.ResponseWriter, r *http.Request) {
	course := Product{
		Name:        "Golang Course",
		Description: "Learn Golang fast and deploy scalable projects",
		Price:       100,
	}

	json.NewEncoder(w).Encode(course)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server listening on localhost:8080")
}
