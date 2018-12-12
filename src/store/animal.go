package store

import (
	"demogo/src/model"
	"errors"
)

func Save(a *model.Animal) error {
	if e := db.Limit(-1).Create(a).Error; e != nil {
		return errors.New("create Animal fail")
	}
	return nil
}
