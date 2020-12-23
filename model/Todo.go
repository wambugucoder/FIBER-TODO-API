package model

import (
	"github.com/Kamva/mgm/v2"
)

// Todo is the model that defines a todo entry
type Todo struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string `json:"title "bson:"title"`
	Description      string `json:"description "bson:"description"`
	Done             bool   `json:"done" bson:"done"`
}

// CreateTodo is the function that creates a Todo
func CreateTodo(title string, description string) *Todo {
	return &Todo{
		Title:       title,
		Description: description,
		Done:        false,
	}
}
