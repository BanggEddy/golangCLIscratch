package cmd

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

	"github.com/BanggEddy/golangCLIscratch/models"
    "github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
    Use:   "interactive",
    Short: "Mode interactif pour gérer les contacts",
    Run: func(cmd *cobra.Command, args []string) {
        reader := bufio.NewReader(os.Stdin)

        for {
            fmt.Println("\n-Contacts version CLI -")
            fmt.Println("1. Ajouter un contact")
            fmt.Println("2. Lister les contacts")
            fmt.Println("3. Mettre à jour un contact")
            fmt.Println("4. Supprimer un contact")
            fmt.Println("5. Quitter")
            fmt.Print("Choix: ")

            choixStr, _ := reader.ReadString('\n')
            choixStr = strings.TrimSpace(choixStr)
            choix, _ := strconv.Atoi(choixStr)

            switch choix {
            case 1:
                fmt.Print("Nom: ")
                nom, _ := reader.ReadString('\n')
                nom = strings.TrimSpace(nom)

                fmt.Print("Email: ")
                email, _ := reader.ReadString('\n')
                email = strings.TrimSpace(email)

                contact := &models.Contact{Name: nom, Email: email}
                if err := storeInstance.CreateContact(contact); err != nil {
                    fmt.Println("Erreur:", err)
                } else {
                    fmt.Println("Contact ajouté, ID:", contact.ID)
                }

            case 2:
                contacts, _ := storeInstance.GetAllContacts()
                if len(contacts) == 0 {
                    fmt.Println("Aucun contact")
                } else {
                    for _, c := range contacts {
                        fmt.Printf("ID: %d | %s | %s\n", c.ID, c.Name, c.Email)
                    }
                }

            case 3:
                fmt.Print("ID du contact à modifier: ")
                idStr, _ := reader.ReadString('\n')
                id, _ := strconv.Atoi(strings.TrimSpace(idStr))

                contact, err := storeInstance.GetContactByID(uint(id))
                if err != nil {
                    fmt.Println("Contact non trouvé")
                    continue
                }

                fmt.Print("Nouveau nom (laisser vide pour ne pas changer): ")
                nom, _ := reader.ReadString('\n')
                nom = strings.TrimSpace(nom)
                if nom != "" {
                    contact.Name = nom
                }

                fmt.Print("Nouvel email (laisser vide pour ne pas changer): ")
                email, _ := reader.ReadString('\n')
                email = strings.TrimSpace(email)
                if email != "" {
                    contact.Email = email
                }

                if err := storeInstance.UpdateContact(contact); err != nil {
                    fmt.Println("Erreur lors de la mise à jour:", err)
                } else {
                    fmt.Println("Contact mis à jour")
                }

            case 4:
                fmt.Print("ID du contact à supprimer: ")
                idStr, _ := reader.ReadString('\n')
                id, _ := strconv.Atoi(strings.TrimSpace(idStr))
                if err := storeInstance.DeleteContact(uint(id)); err != nil {
                    fmt.Println("Erreur:", err)
                } else {
                    fmt.Println("Contact supprimé")
                }

            case 5:
                fmt.Println("Au revoir!")
                os.Exit(0)

            default:
                fmt.Println("Choix invalide")
            }
        }
    },
}

func init() {
    rootCmd.AddCommand(interactiveCmd)
}
