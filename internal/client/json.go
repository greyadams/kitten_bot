package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{
		"user": {
		  "id": 123,
		  "name": "Alice",
		  "email": "alice@example.com"
		},
		"status": "active"
	  }`

	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}
	userRaw, ok := data["user"]
	if !ok {
		fmt.Println("Ошибка")
		return
	}

	user, ok := userRaw.(map[string]interface{})
	if !ok {
		fmt.Println("Ошибка")
		return
	}

	name := user["name"]
	email := user["email"]

	fmt.Println(name)
	fmt.Println(email)
}
