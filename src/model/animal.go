package model

import "github.com/jinzhu/gorm"

type Animal struct {
	gorm.Model
	Name string `json:"name" gorm:"default:'galeone'"`
	Age  int64
}

type CreateParam struct {
	Name string
	Age int64
}

func NewAnimal(c CreateParam) *Animal {
	return &Animal{
		Name: c.Name,
		Age: c.Age,
	}
}

