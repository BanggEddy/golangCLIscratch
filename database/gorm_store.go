package database

import (
    "log"

    "github.com/BanggEddy/golangCLIscratch/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type GORMStore struct {
    db *gorm.DB
}

func NewGORMStore(dsn string) *GORMStore {
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Impossible de se connecter à la base :", err)
    }
    db.AutoMigrate(&models.Contact{})
    return &GORMStore{db: db}
}

func (s *GORMStore) CreateContact(contact *models.Contact) error {
    return s.db.Create(contact).Error
}

func (s *GORMStore) GetAllContacts() ([]models.Contact, error) {
    var contacts []models.Contact
    err := s.db.Find(&contacts).Error
    return contacts, err
}

func (s *GORMStore) GetContactByID(id uint) (*models.Contact, error) {
    var c models.Contact
    err := s.db.First(&c, id).Error
    if err != nil {
        return nil, err
    }
    return &c, nil
}

func (s *GORMStore) UpdateContact(contact *models.Contact) error {
    return s.db.Save(contact).Error
}

func (s *GORMStore) DeleteContact(id uint) error {
    return s.db.Delete(&models.Contact{}, id).Error
}
