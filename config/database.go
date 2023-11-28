package config

import (
	"github.com/fawzy17/gin-relation-table/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gin-relation-table-fix?parseTime=true"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Kelas{})
	database.AutoMigrate(&models.Mahasiswa{})

	DB = database
}
