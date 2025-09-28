package cmd

import (
    "fmt"
    "log"
    "github.com/BanggEddy/golangCLIscratch/database"
    "github.com/BanggEddy/golangCLIscratch/store"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var storeInstance store.Storer

var rootCmd = &cobra.Command{
    Use:   "mini-crm",
    Short: "Mini-CRM CLI : Gestion des contacts",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        t := viper.GetString("storage.type")
        switch t {
        case "gorm":
            storeInstance = database.NewGORMStore(viper.GetString("database.dsn"))
        case "json":
            storeInstance = store.NewJSONStore(viper.GetString("database.file"))
        default:
            log.Fatal("Type de stockage inconnu:", t)
        }
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }
}
