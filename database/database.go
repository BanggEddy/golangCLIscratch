package database

import (
	"log"
	
	"github.com/BanggEddy/golangCLIscratch/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Echec de la connexion bdd : ", err)
	}
	
	log.Println("Connexion à la bdd SQLite ok")
}

func AutoMigrate() {
	err := DB.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatal("Echec migration bdd : ", err)
	}
	
	log.Println("Migration bdd OK")
}