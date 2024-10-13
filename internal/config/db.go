package config

import (
	"fmt"
	"log"

	"github.com/acgyiyo/payment_api_test/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.Dbname, config.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	err = DB.AutoMigrate(&entity.Payment{})
	if err != nil {
		log.Fatalf("Failed to init migrations: %v", err)
		return nil, err
	}
	log.Print("Init DB successfully")

	// only for development porpouses
	err = InsertInitValues(DB)
	if err != nil {
		log.Fatalf("Failed to insert development values: %v", err)
	}

	return DB, nil
}

// only for development porpouses
func InsertInitValues(db *gorm.DB) error {
	merchants := []entity.Merchant{
		{
			Name: "Merchant 1",
		},
		{
			Name: "Merchant 2",
		},
	}

	// Insertar los registros en la base de datos
	if err := db.Create(&merchants).Error; err != nil {
		return err
	}

	//merchants will init with ID 1 and 2...for each app init

	return nil
}
