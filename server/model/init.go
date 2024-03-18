package model

import "gorm.io/gorm"

var DB *gorm.DB

func Init() {

}

func CreateTables(db *gorm.DB) {
	var err error
	err = db.AutoMigrate(&SysUser{})

	if err != nil {
		panic("创建表失败！")
	}
}
