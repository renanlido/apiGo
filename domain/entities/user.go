package entities

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/renanlido/api-go/domain/contracts"
)

type todo struct {
	Id        *string    `json:"id"`
	Title     string     `json:"title"`
	Status    *bool      `json:"status"`
	CreatedAt *time.Time `json:"createdAt" `
	UpdatedAt *time.Time `json:"updatedAt"`
}

func Constructor(data []byte, UUIDGenerator contracts.UUIDGenerator) (*todo, error) {
	status := false
	newDate := time.Now()

	todo := &todo{}

	if err := json.Unmarshal(data, &todo); err != nil {
		return nil, err
	}

	if todo.Title == "" {
		return todo, fmt.Errorf("title is not provided")
	}

	if todo.Id == nil {
		id := UUIDGenerator()
		todo.Id = &id
	}

	if todo.Status == nil {
		todo.Status = &status
	}

	if todo.CreatedAt == nil {
		todo.CreatedAt = &newDate
	}

	todo.UpdatedAt = &newDate

	return todo, nil
}
