package db

import (
	"os"

	"github.com/Z-M-Huang/RealEstateScraper/utils"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mssql" //supporting packages
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//dbContext database connection
var dbContext *gorm.DB

//InitDB init db
func InitDB() {
	var err error
	dbContext, err = gorm.Open(os.Getenv("DB_DRIVER"), os.Getenv("CONNECTION_STRING"))
	if err != nil {
		utils.Logger.Sugar().Fatalf("failed to open database %s", err.Error())
	}
	migrate()
}
func migrate() {
	dbContext.AutoMigrate(&Trulia{})
}

//DoTransaction do transaction
func DoTransaction(fc func(tx *gorm.DB) error) error {
	return dbContext.Transaction(fc)
}

//Disconnect disconnect
func Disconnect() error {
	return dbContext.Close()
}
