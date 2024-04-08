package db

import (
	"eatry/pkg/config"
	"eatry/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&domain.AdminDetails{})
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Coupon{})
	db.AutoMigrate(&domain.UsedCoupon{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.UserOrderItem{})
	db.AutoMigrate(&domain.OrderSuccessResponse{})
	db.AutoMigrate(&domain.RazorPay{})
	db.AutoMigrate(&domain.WishList{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.WalletHistory{})
	db.AutoMigrate(&domain.NewWalletHistory{})

	return db, dbErr

}

// test comment
