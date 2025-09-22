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

func main() {
	nom := flag.String("nom", "", "Nom du contact")
	email := flag.String("email", "", "Email du contact")
	flag.Parse()

	contacts := make(map[int]Contact)
	prochainID := 1

	if *nom != "" && *email != "" {
		contact := Contact{ID: prochainID, Nom: *nom, Email: *email}
		contacts[prochainID] = contact
		fmt.Printf("contact ajouté: id %d\n", prochainID)
		return
	}

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

			contact := Contact{ID: prochainID, Nom: nom, Email: email}
			contacts[prochainID] = contact
			fmt.Printf("Contact ajouté avec ID: %d\n", prochainID)
			prochainID++
		case 2:
			if len(contacts) == 0 {
				fmt.Println("Aucun contact")
			} else {
				fmt.Println("\n--- Contacts ---")
				for _, contact := range contacts {
					fmt.Printf("id : %d | %s | %s\n", contact.ID, contact.Nom, contact.Email)
				}
			}
		case 3:
			fmt.Print("ID à supprimer: ")
			var id int
			fmt.Scan(&id)
			val, ok := contacts[id]

			if ok {
				delete(contacts, id)
				fmt.Println("Contact deleted:", val)
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
