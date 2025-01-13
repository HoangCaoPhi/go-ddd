package persistence

import (
	"fmt"
	"phihc116/go-task/backend/domain/tags"
	"phihc116/go-task/backend/shared"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDatabase() {
	dsn := "sqlserver://%s:%s@%s?database=%s&encrypt=disable&trustServerCertificate=true"

	sqlServerOptions := shared.Config.SqlServerOptions
	connectionString := fmt.Sprintf(dsn,
		sqlServerOptions.UserName,
		sqlServerOptions.Password,
		sqlServerOptions.Host,
		sqlServerOptions.DbName)

	fmt.Print(connectionString)

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err != nil {
		fmt.Print("Init db error")
	}

	db.AutoMigrate(&tags.Tag{})

	shared.DbContext = db
}
