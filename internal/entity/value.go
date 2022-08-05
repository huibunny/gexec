// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type TableEntity struct {
	Table   string          `json:"table" example:"t_app"`
	Columns []string        `json:"columns" example:"[name,icon,cover]"`
	Values  [][]interface{} `json:"values" example:"[[note,https://bn.com/note.png,https://bn.com/note.png]],[[todo,https://bn.com/todo.png,https://bn.com/todo.png]]"`
}
