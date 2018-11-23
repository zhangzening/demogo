package store

import (
	"BeyondMesh/src/api/bootstrap"
	"demogo/src/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func migrate() {
	db.AutoMigrate(&model.Animal{})
}

func Init() {
	var err error
	db, err = gorm.Open("mysql", "naftis:naftisIsAwesome@tcp(10.20.1.175:35689)/naftis?charset=utf8")
	if err != nil {
		panic(fmt.Errorf("failed to connect database %s", err))
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(bootstrap.Debug())

	fmt.Println("[db] init success")
	fmt.Println("table migrate init start ...")
	migrate()
	fmt.Println("table migrate init end ...")

}