package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"snake/model"
)

// 生成假数据
func generateFakeData(db *gorm.DB) {
	users := []model.SysUser{
		{UserName: "admin", NickName: "管理员", Password: "123"},
		{UserName: "user1", NickName: "用户1", Password: "123"},
		{UserName: "user2", NickName: "用户2", Password: "123"},
	}
	result := db.Create(&users)
	if result.Error != nil {
		panic(result.Error)
	} else {
		fmt.Printf("插入成功, %d 条数据\n", result.RowsAffected)
	}
}

// 生成表结构
func main() {
	host := "127.0.0.1"
	user := "postgres"
	password := "888888"
	dbName := "snake"
	port := 5432
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.SysUser{})
	if err != nil {
		panic(err)
	}
	generateFakeData(db)

}
