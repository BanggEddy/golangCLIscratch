package main

import (
	"fmt"
	"log"

	"github.com/BanggEddy/golangCLIscratch/database"
	"github.com/BanggEddy/golangCLIscratch/handlers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// middleware
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func InitConfig() {
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.name", "contacts.db")
	viper.SetDefault("database.dsn", "./contacts.db")
	viper.SetDefault("app.environment", "development")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Fichier de configuration config.yaml non trouvé, utilisation des valeurs par défaut")
		} else {
			log.Fatal("Erreur lors de la lecture du fichier de configuration: ", err)
		}
	} else {
		log.Println("Fichier de configuration chargé avec succès.")
	}

	fmt.Printf("Environnement: %s\n", viper.GetString("app.environment"))
	fmt.Printf("Port du serveur: %d\n", viper.GetInt("server.port"))
}

func main() {
	InitConfig()
	database.ConnectDB()
	database.AutoMigrate()
	router := gin.Default()
	router.Use(LoggerMiddleware())

	// pour tester
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	// groupe de routes api
	api := router.Group("/api/v1")
	{
		contacts := api.Group("/contacts")
		{
			contacts.POST("/", handlers.CreateContact)     
			contacts.GET("/", handlers.GetAllContacts)     
			contacts.GET("/:id", handlers.GetContactByID)   
			contacts.PUT("/:id", handlers.UpdateContact)    
			contacts.DELETE("/:id", handlers.DeleteContact) 
		}
	}

	// Démarre serveur
	port := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	log.Printf("Serveur démarré sur le port %s", port)
	router.Run(port)
}