package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Adib",
		MiddleName: "Hauzan",
		LastName:   "Sofyan",
		Hobbies: []string{
			"Gaming",
			"Reading",
			"Coding",
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	request := `{"FirstName":"Adib","MiddleName":"Hauzan","LastName":"Sofyan","Hobbies":["Gaming","Reading","Coding"]}`
	jsonByte := []byte(request)

	customer := &Customer{}

	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestJsonArrayCompleEncode(t *testing.T) {
	customer := Customer{
		FirstName: "Adib",
		Hobbies:   []string{"Coding", "Gaming", "Reading"},
		Addresses: []Address{
			{
				Street:     "Belum Ada",
				Country:    "Indonesia",
				PostalCode: "6969",
			},
			{
				Street:     "Cijerokaso",
				Country:    "Australia",
				PostalCode: "696969",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	request := `{"FirstName":"Adib","MiddleName":"","LastName":"","Hobbies":["Coding","Gaming","Reading"],"Addresses":[{"Street":"Belum Ada","Country":"Indonesia","PostalCode":"6969"},{"Street":"Cijerokaso","Country":"Australia","PostalCode":"696969"}]}`
	bytes := []byte(request)

	customer := &Customer{}
	err := json.Unmarshal(bytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}
