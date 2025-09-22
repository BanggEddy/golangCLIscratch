package main

import (
	"flag"
	"fmt"
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

	}
}
