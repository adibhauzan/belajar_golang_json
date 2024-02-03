package belajar_golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestEncode(t *testing.T) {
	logJson("Adib")
	logJson(22)
	logJson(true)
	logJson([]string{"Adib", "Hauzan", "Sofyan"})
}