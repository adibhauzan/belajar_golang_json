package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJsonObject(t *testing.T) {
	adib := Customer{
		FirstName:  "Adib",
		MiddleName: "Hauzan",
		LastName:   "Sofyan",
	}

	bytes, _ := json.Marshal(adib)
	fmt.Println(string(bytes))
}
