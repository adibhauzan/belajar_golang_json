package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamingDecoder(t *testing.T) {
	reader, err := os.Open("Product.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(reader)

	// Inisialisasi objek Customer
	product := &Product{}

	// Gunakan pointer ke customer saat melakukan Decode
	err = decoder.Decode(product)
	if err != nil {
		panic(err)
	}

	// Tampilkan hasil
	fmt.Println(product)
}
