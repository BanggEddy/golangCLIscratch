package cmd

import (
    "fmt"
    "github.com/BanggEddy/golangCLIscratch/models"
    "github.com/spf13/cobra"
)

var name, email string

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Ajouter un contact",
    Run: func(cmd *cobra.Command, args []string) {
        contact := &models.Contact{Name: name, Email: email}
        if err := storeInstance.CreateContact(contact); err != nil {
            fmt.Println("Erreur lors de la création:", err)
            return
        }
        fmt.Println("Contact ajouté avec succès, ID:", contact.ID)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
    addCmd.Flags().StringVarP(&name, "name", "n", "", "Nom du contact")
    addCmd.Flags().StringVarP(&email, "email", "e", "", "Email du contact")
    addCmd.MarkFlagRequired("name")
    addCmd.MarkFlagRequired("email")
}
