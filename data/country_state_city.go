package data

import "github.com/jinzhu/gorm"

//CountryStateCity country state city
type CountryStateCity struct {
	gorm.Model
	Country string `gorm:"not null;index:countrystatecity"`
	State   string `gorm:"index:countrystatecity"`
	City    string `gorm:"index:countrystatecity"`

	Trulia []TruliaEntity
}
