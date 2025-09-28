package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var getID uint

var getCmd = &cobra.Command{
    Use:   "get",
    Short: "Afficher un contact par ID",
    Run: func(cmd *cobra.Command, args []string) {
        contact, err := storeInstance.GetContactByID(getID)
        if err != nil {
            fmt.Println("Erreur lors de la récupération :", err)
            return
        }

        fmt.Printf("ID: %d\nNom: %s\nEmail: %s\n", contact.ID, contact.Name, contact.Email)
    },
}

func init() {
    rootCmd.AddCommand(getCmd)
    getCmd.Flags().UintVarP(&getID, "id", "i", 0, "ID du contact à afficher")
    getCmd.MarkFlagRequired("id")
}
