# CRM Golang (CLI + API REST)

Un projet en Go 1.25+ permettant de gérer une base de contacts :

* Mode CLI : ajout, suppression, mise à jour, affichage interactif des contacts
* Mode Serveur REST : API RESTful exposée avec Gin sur localhost:8080

Les données sont stockées dans SQLite (via GORM) ou au format JSON (selon la config).
---

### Prérequis (dév sous windows)

- [Go](https://go.dev/) 

## Arborescence du projet
```bash
golangCLIscratch/
│   contacts.db (générée automatiquement via GORM)
│   go.mod
│   go.sum
│   golangCLIscratch.exe
│   main.go
│   README.md
│
├───cmd/
│    └─ add.go
│    └─ delete.go
│    └─ get.go
│    └─ interactive.go
│    └─ list.go
│    └─ root.go
│    └─ update.go  
├───config/
│    └─ config.go      
├───database/
│    └─ database.go
│    └─ gorm_store.go    
├───dto/
│    └─ contact_input.go
│    └─ contact_output.go    
├───handlers/
│    └─ contact_handler.go  
├───models/
│    └─ contact.go  
├───server/
│    └─ server.go        
└───store/
     └─ storer.go
     └─ jsonstore.go         

```

## Cloner le projet :

```bash
git clone https://github.com/BanggEddy/golangCLIscratch.git
cd golangCLIscratch
```

## Voir commandes pour faires les opérations CRUD directement par ligne de commande:

```bash
go run main.go --mode cli
```

## Résultat :

```bash
Usage:
  mini-crm [command]

Available Commands:
  add         Ajouter un contact
  completion  Generate the autocompletion script for the specified shell
  delete      Supprimer un contact
  get         Afficher un contact par ID
  help        Help about any command
  interactive Mode interactif pour gérer les contacts
  list        Lister tous les contacts
  update      Mettre à jour un contact
```

## Commandes disponibles :
```
go run main.go add --name "Alice" --email "alice@example.com"
go run main.go list
go run main.go get --id 1
go run main.go update --id 1 --name "Alice B."
go run main.go delete --id 1
go run main.go interactive  
```

## Lancer en mode Serveur API (dans la racine du projet quand même) :

```bash
go run main.go --mode server
```

## Affiche dans la console :

```bash
Fichier de configuration chargé avec succès.
Environnement: development
Port du serveur: 8080
Connexion à la bdd SQLite ok
Migration bdd OK
Serveur démarré sur le port :8080
```
L’API est dispo sur http://localhost:8080/api/v1/contacts.

## Lancer en mode CLI
```bash
go run main.go interactive
```

## Affiche le menu interactif :
```bash
-Contacts version CLI -
1. Ajouter un contact
2. Lister les contacts
3. Mettre à jour un contact
4. Supprimer un contact
5. Quitter
Choix:
```

## API REST avec Postman
Base URL :
```bash
http://localhost:8080/api/v1/contacts
```

## Endpoints disponibles
| Méthode | URL                  | Description                  |
|---------|----------------------|------------------------------|
| POST    | /api/v1/contacts/    | Créer un nouveau contact     |
| GET     | /api/v1/contacts/    | Récupérer tous les contacts  |
| GET     | /api/v1/contacts/:id | Récupérer un contact par ID  |
| PUT     | /api/v1/contacts/:id | Mettre à jour un contact     |
| DELETE  | /api/v1/contacts/:id | Supprimer un contact         |

## Exemple POST (création d’un contact):
1. URL :
```bash
POST http://localhost:8080/api/v1/contacts/
```

2. Body (JSON) :
```bash
{
  "name": "Alice Dupont",
  "email": "alice@example.com"
}
```

Réponse (201) :
```bash
{
  "id": 1,
  "name": "Alice Dupont",
  "email": "alice@example.com"
}
```
Côté Terminal:
```bash
[GIN-debug] Listening and serving HTTP on :8080
[GIN-debug] redirecting request 301: /api/v1/contacts/ --> /api/v1/contacts/
2025/09/28 22:47:42 GET /api/v1/contacts/
[GIN] 2025/09/28 - 22:47:42 | 200 |      1.0052ms |             ::1 | GET      "/api/v1/contacts/"
[GIN-debug] redirecting request 307: /api/v1/contacts/ --> /api/v1/contacts/
2025/09/28 22:47:57 POST /api/v1/contacts/
[GIN] 2025/09/28 - 22:47:57 | 201 |     14.1889ms |             ::1 | POST     "/api/v1/contacts/"
```







