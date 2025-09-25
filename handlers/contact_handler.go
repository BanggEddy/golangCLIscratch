package handlers

import (
	"net/http"
	"strconv"

	"github.com/BanggEddy/golangCLIscratch/database"
	"github.com/BanggEddy/golangCLIscratch/dto"
	"github.com/BanggEddy/golangCLIscratch/models"
	"github.com/gin-gonic/gin"
)

// helper pour éviter d'avoir le même code tout le temps partout
func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return 0, false
	}
	return uint(id), true
}

func contactToDTO(contact models.Contact) dto.ContactOutputDTO {
	return dto.ContactOutputDTO{
		ID:    contact.ID,
		Name:  contact.Name,
		Email: contact.Email,
	}
}

// fonction add
func CreateContact(c *gin.Context) {
	var input dto.ContactInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact := models.Contact{Name: input.Name, Email: input.Email}
	if err := database.DB.Create(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	c.JSON(http.StatusCreated, contactToDTO(contact))
}