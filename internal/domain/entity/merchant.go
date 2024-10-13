package entity

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name     string
	Payments []Payment `gorm:"foreignKey:MerchantID"`
}
