package store

import "github.com/BanggEddy/golangCLIscratch/models"

type Storer interface {
    CreateContact(contact *models.Contact) error
    GetAllContacts() ([]models.Contact, error)
    GetContactByID(id uint) (*models.Contact, error)
    UpdateContact(contact *models.Contact) error
    DeleteContact(id uint) error
}
