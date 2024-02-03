package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingEncoder(t *testing.T) {
	writer, err := os.Create("sample_output.json")
	if err != nil {
		panic(err)
	}
	encoder := json.NewEncoder(writer)

	product := Product{
		Id:    "p001",
		Name:  "Macbook Pro Ahayy",
		Image: "http://bujang.com",
		Price: 1000000,
	}
	err = encoder.Encode(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
