package database

import (
	_ "database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	fmt.Println("koneksi database terpanggil")
	// Sesuaikan dengan user/password/database MySQL kamu
	dsn := "root:@tcp(localhost:3306)/db_web_gudang?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("gagal koneksi ke database:", err)
	} else {
		DB = db
		DB.SingularTable(true) //Ini config dari GORM
		fmt.Println("koneksi ke database sukses")
	}
}
