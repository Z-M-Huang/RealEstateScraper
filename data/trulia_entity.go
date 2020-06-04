package data

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//TruliaEntity data from Trulia
type TruliaEntity struct {
	gorm.Model
	URL                  string `gorm:"not null;unique_index"`
	CountryStateCityID   uint   `gorm:"not null"`
	Address              string `gorm:"unique"`
	Bedrooms             float32
	Bathrooms            float32
	Area                 float32
	Description          string
	IsForRent            bool    `gorm:"not null;default:false"`
	IsForSale            bool    `gorm:"not null;default:false"`
	UnitPrice            float32 `gorm:"not null;default:0"`
	Type                 string  `gorm:"not null"`
	HOAFee               float32 `gorm:"not null;default:0"`
	TaxRateCodeArea      string
	YearBuilt            int
	Heating              string
	HeatingFuel          string
	CoolingSystem        string
	AC                   bool           `gorm:"not null;default:false"`
	Washer               bool           `gorm:"not null;default:false"`
	Dryer                bool           `gorm:"not null;default:false"`
	Refrigerator         bool           `gorm:"not null;default:false"`
	Microwave            bool           `gorm:"not null;default:false"`
	Dishwasher           bool           `gorm:"not null;default:false"`
	Disposal             bool           `gorm:"not null;default:false"`
	Floors               pq.StringArray `gorm:"type:text[]"`
	Garden               bool           `gorm:"not null;default:false"`
	Parking              string
	Garage               bool `gorm:"not null;default:false"`
	ParkingSpace         int  `gorm:"not null;default:0"`
	FitnessCenter        bool `gorm:"not null;default:false"`
	Exterior             string
	Foundations          string
	Roof                 string
	Images               pq.StringArray `gorm:"type:text[]"`
	AmtElementarySchools int            `gorm:"not null;default:0"`
	AmtMiddleSchools     int            `gorm:"not null;default:0"`
	AmtHighSchools       int            `gorm:"not null;default:0"`
	CrimeRate            string
	Restaurant           int `gorm:"not null;default:0"`
	Groceires            int `gorm:"not null;default:0"`
	Nightlife            int `gorm:"not null;default:0"`
	LastModifiedDate     time.Time
}
