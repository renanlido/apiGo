package main

import (
	"encoding/json"
	"fmt"

	"github.com/renanlido/api-go/domain/entities"
	uuid "github.com/satori/go.uuid"
)

func main() {

	todoData := []byte(`{"id":"9d8733a2-fa8d-4b01-adea-f471eb42d8ac", "title":"fazer compras","status":true, "createdAt":"2023-01-29T21:52:41.757197972-03:00"}`)

	todo, err := entities.Constructor(todoData, uuid.NewV4().String)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoJson, err := json.Marshal(todo)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(todoJson))
}
