package main

import (
    "flag"
    "log"

    "github.com/BanggEddy/golangCLIscratch/cmd"
    "github.com/BanggEddy/golangCLIscratch/server" 
    "github.com/spf13/viper"
)

func main() {
	mode := flag.String("mode", "cli", "Mode d'exécution: cli ou server")
	flag.Parse()

    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./config")
    if err := viper.ReadInConfig(); err != nil {
        log.Println("Fichier config non trouvé, valeurs par défaut utilisées")
    }

    switch *mode {
    case "cli":
        cmd.Execute()  
    case "server":
        server.Run()   
    default:
        log.Fatal("Mode inconnu: ", *mode)
    }
}
