
package database

import (
	"fmt"
	"log"

	"github.com/nuty/simple-blog/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite" 
	"github.com/nuty/simple-blog/models"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Dbname,
		config.Database.Sslmode)

	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
}


func ConnectDBTest() {
	var err error
	dsn := "file:test.db?cache=shared&mode=rwc"
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error to connect to database:", err)
	}

	err = DB.AutoMigrate(
		&models.Article{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Connected to SQLite database successfully!")
}