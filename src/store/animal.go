package store

import (
	"demogo/src/model"
	"errors"
)

func Save(a *model.Animal) error {

	animal := model.Animal{Age: 99}
	if e := db.Limit(-1).Create(&animal).Error; e != nil {
		return errors.New("create Animal fail")
	}
	return nil
}


