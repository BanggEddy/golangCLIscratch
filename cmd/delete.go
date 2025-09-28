package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var deleteID uint

var deleteCmd = &cobra.Command{
    Use:   "delete",
    Short: "Supprimer un contact",
    Run: func(cmd *cobra.Command, args []string) {
        if err := storeInstance.DeleteContact(deleteID); err != nil {
            fmt.Println("Erreur lors de la suppression :", err)
            return
        }
        fmt.Println("Contact supprimé avec succès !")
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
    deleteCmd.Flags().UintVarP(&deleteID, "id", "i", 0, "ID du contact à supprimer")
    deleteCmd.MarkFlagRequired("id")
}
