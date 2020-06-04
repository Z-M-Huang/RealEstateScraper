package data

import (
	"os"
	"strings"

	"github.com/Z-M-Huang/RealEstateScraper/utils"
	"github.com/jinzhu/gorm"
)

//dbContext database connection
var dbContext *gorm.DB

//InitDB init db
func InitDB() {
	driver := strings.TrimSpace(os.Getenv("DB_DRIVER"))
	connStr := strings.TrimSpace(os.Getenv("CONNECTION_STRING"))
	if driver == "" || connStr == "" {
		utils.Logger.Fatal("No DB connection string")
	}
	var err error
	dbContext, err = gorm.Open(driver, connStr)
	if err != nil {
		utils.Logger.Sugar().Fatalf("failed to open database %s", err.Error())
	}
	migrate()
}
func migrate() {
	dbContext.AutoMigrate(&TruliaEntity{})
}

//DoTransaction do transaction
func DoTransaction(fc func(tx *gorm.DB) error) error {
	return dbContext.Transaction(fc)
}

//Disconnect disconnect
func Disconnect() error {
	return dbContext.Close()
}
