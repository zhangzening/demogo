package service

import (
	"demogo/src/model"
	"demogo/src/store"
)

var DefaultAnimalServer animalServer

type animalServer interface {
	Create(param model.CreateParam) error
}

type animal struct {}

func InitAnimalServer() {
	DefaultAnimalServer =  &animal{}
}

func (a *animal) Create(param model.CreateParam) error {
	newAnimal := model.NewAnimal(param)
	return store.Save(newAnimal)
}