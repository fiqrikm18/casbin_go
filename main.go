package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var (
	DB *gorm.DB
)

func main() {
	CreateConnection()

	adapter, err := gormadapter.NewAdapterByDB(DB)
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer("rbac_model.conf", adapter)
	if err != nil {
		panic(err)
	}

	enforcer.LoadPolicy()
	status, err := enforcer.Enforce("alice", "data1", "read")
	if err != nil {
		panic(err)
	}

	fmt.Println(status)
	err = enforcer.SavePolicy()
	if err != nil {
		panic(err)
	}
}

func CreateConnection() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_casbin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}
