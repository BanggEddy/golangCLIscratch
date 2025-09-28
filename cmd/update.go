package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var updateID uint
var updateName, updateEmail string

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "Mettre à jour un contact",
    Run: func(cmd *cobra.Command, args []string) {
        contact, err := storeInstance.GetContactByID(updateID)
        if err != nil {
            fmt.Println("Contact non trouvé:", err)
            return
        }

        if updateName != "" {
            contact.Name = updateName
        }
        if updateEmail != "" {
            contact.Email = updateEmail
        }

        if err := storeInstance.UpdateContact(contact); err != nil {
            fmt.Println("Erreur lors de la mise à jour :", err)
            return
        }

        fmt.Println("Contact mis à jour avec succès !")
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
    updateCmd.Flags().UintVarP(&updateID, "id", "i", 0, "ID du contact à mettre à jour")
    updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "Nouveau nom")
    updateCmd.Flags().StringVarP(&updateEmail, "email", "e", "", "Nouvel email")
    updateCmd.MarkFlagRequired("id")
}
