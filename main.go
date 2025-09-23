package main

import (
    "flag"
    "fmt"
    "os"
)

type Contact struct {
    ID    int
    Nom   string
    Email string
}

func NewContact(id int, nom, email string) *Contact {
    return &Contact{
        ID:    id,
        Nom:   nom,
        Email: email,
    }
}

func (c *Contact) Ajouter(contacts map[int]*Contact) {
    contacts[c.ID] = c
    fmt.Printf("Contact ajouté avec ID: %d\n", c.ID)
}

func SupprimerContact(contacts map[int]*Contact, id int) {
    if contact, ok := contacts[id]; ok {
        delete(contacts, id)
        fmt.Printf("Contact supprimé: %s\n", contact.Nom)
    } else {
        fmt.Println("Contact non trouvé")
    }
}

func main() {
    nom := flag.String("nom", "", "Nom du contact")
    email := flag.String("email", "", "Email du contact")
    flag.Parse()

    contacts := make(map[int]*Contact)
    prochainID := 1

    // Ajout via ligne de commande
    if *nom != "" && *email != "" {
        contact := NewContact(prochainID, *nom, *email)
        contact.Ajouter(contacts)
        return
    }

    // Menu principal
    for {
        fmt.Println("\n=== Menu ===")
        fmt.Println("|| 1. Ajouter un contact")
        fmt.Println("|| 2. Lister les contacts")
        fmt.Println("|| 3. Supprimer un contact")
        fmt.Println("|| 4. Mettre à jour un contact")
        fmt.Println("|| 5. Quitter")
        fmt.Print("Choix: ")

        var choix int
        fmt.Scan(&choix)

        switch choix {
        case 1:
            var nom, email string
            fmt.Print("Nom: ")
            fmt.Scan(&nom)
            fmt.Print("Email: ")
            fmt.Scan(&email)
            
            contact := NewContact(prochainID, nom, email)
            contact.Ajouter(contacts)
            prochainID++

        case 2:
            if len(contacts) == 0 {
                fmt.Println("Aucun contact")
            } else {
                fmt.Println("\n--- Contacts ---")
                for _, contact := range contacts {
                    fmt.Printf("ID: %d | %s | %s\n", contact.ID, contact.Nom, contact.Email)
                }
            }

        case 3:
            fmt.Print("ID à supprimer: ")
            var id int
            fmt.Scan(&id)
            SupprimerContact(contacts, id)

        case 4:
            fmt.Print("ID à modifier: ")
            var id int
            fmt.Scan(&id)
            
            if contact, ok := contacts[id]; ok {
                fmt.Print("Nouveau nom: ")
                fmt.Scan(&contact.Nom)
                fmt.Print("Nouvel email: ")
                fmt.Scan(&contact.Email)
                fmt.Println("Contact mis à jour")
            } else {
                fmt.Println("Contact non trouvé")
            }

        case 5:
            fmt.Println("Au revoir!")
            os.Exit(0)

        default:
            fmt.Println("Choix invalide")
        }
    }
}