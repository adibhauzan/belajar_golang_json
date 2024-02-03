package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image_url"`
	Price int    `json:"price"`
}

func TestJsonTagEncode(t *testing.T) {
	product := Product{
		Id:    "1",
		Name:  "Macbook Pro Ahayy",
		Image: "http://example.com",
	}

	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	request := `{"id":"1","name":"Macbook Pro Ahayy","image_url":"http://example.com"}`
	bytes := []byte(request)

	product := &Product{}
	err := json.Unmarshal(bytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
