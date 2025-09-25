package handlers

import (
	"net/http"
	"strconv"

	"github.com/BanggEddy/golangCLIscratch/database"
	"github.com/BanggEddy/golangCLIscratch/dto"
	"github.com/BanggEddy/golangCLIscratch/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// fonction récup all contacts
func GetAllContacts(c *gin.Context) {
	var contacts []models.Contact
	if err := database.DB.Find(&contacts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	response := make([]dto.ContactOutputDTO, len(contacts))
	for i, contact := range contacts {
		response[i] = contactToDTO(contact)
	}

	c.JSON(http.StatusOK, response)
}

//fonction récup un id contact
func GetContactByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var contact models.Contact
	if err := database.DB.First(&contact, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contact non trouvé"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	c.JSON(http.StatusOK, contactToDTO(contact))
}

// fonction update un contact
func UpdateContact(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var input dto.ContactInputDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var contact models.Contact
	if err := database.DB.First(&contact, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contact non trouvé"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	contact.Name = input.Name
	contact.Email = input.Email
	if err := database.DB.Save(&contact).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}

	c.JSON(http.StatusOK, contactToDTO(contact))
}

func DeleteContact(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	result := database.DB.Delete(&models.Contact{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur serveur"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact non trouvé"})
		return
	}

	c.Status(http.StatusNoContent)
}
