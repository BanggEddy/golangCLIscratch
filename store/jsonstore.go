package store

import (
    "encoding/json"
    "errors"
    "os"
    "sync"

    "github.com/BanggEddy/golangCLIscratch/models"
)

type JSONStore struct {
    file     string
    contacts []models.Contact
    mu       sync.Mutex
}

func NewJSONStore(file string) *JSONStore {
    s := &JSONStore{file: file}
    s.load()
    return s
}

func (s *JSONStore) load() {
    f, err := os.Open(s.file)
    if err != nil {
        return
    }
    defer f.Close()
    json.NewDecoder(f).Decode(&s.contacts)
}

func (s *JSONStore) save() error {
    f, err := os.Create(s.file)
    if err != nil {
        return err
    }
    defer f.Close()
    return json.NewEncoder(f).Encode(s.contacts)
}

func (s *JSONStore) CreateContact(c *models.Contact) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    c.ID = uint(len(s.contacts) + 1)
    s.contacts = append(s.contacts, *c)
    return s.save()
}

func (s *JSONStore) GetAllContacts() ([]models.Contact, error) {
    return s.contacts, nil
}

func (s *JSONStore) GetContactByID(id uint) (*models.Contact, error) {
    for _, c := range s.contacts {
        if c.ID == id {
            return &c, nil
        }
    }
    return nil, errors.New("contact non trouvé")
}

func (s *JSONStore) UpdateContact(contact *models.Contact) error {
    for i, c := range s.contacts {
        if c.ID == contact.ID {
            s.contacts[i] = *contact
            return s.save()
        }
    }
    return errors.New("contact non trouvé")
}

func (s *JSONStore) DeleteContact(id uint) error {
    for i, c := range s.contacts {
        if c.ID == id {
            s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
            return s.save()
        }
    }
    return errors.New("contact non trouvé")
}
