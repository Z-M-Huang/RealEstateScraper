package db

import (
	"github.com/jinzhu/gorm"
)

//Trulia database entity
type Trulia struct {
	gorm.Model
	URL   string `gorm:"unique_index;not null"`
	State string `gorm:"index:statecity;not null"`
	City  string `gorm:"index:statecity;not null"`
	Data  []byte
}

//Find populate current object
func (t *Trulia) Find() error {
	if db := dbContext.Where(*t).First(&t); db.Error != nil {
		return db.Error
	}
	return nil
}

//FindWithTx populate current object with transaction
func (t *Trulia) FindWithTx(tx *gorm.DB) error {
	if db := tx.Where(*t).First(&t); db.Error != nil {
		return db.Error
	}
	return nil
}

//Save save current user
func (t *Trulia) Save() error {
	if db := dbContext.Save(t).Scan(&t); db.Error != nil {
		return db.Error
	}
	return nil
}

//SaveWithTx save current user with transaction
func (t *Trulia) SaveWithTx(tx *gorm.DB) error {
	if db := tx.Save(t).Scan(&t); db.Error != nil {
		return db.Error
	}
	return nil
}
