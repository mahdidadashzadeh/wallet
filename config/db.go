package config

import (
	"wallet-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "root:mahdidatabase1383525@gmail.com@tcp(127.0.0.1:3306)/wallet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Balance{}, &models.TransactionHistory{})
	return db, nil
}
