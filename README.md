# Mini-CRM en Go (La gestion des contacts)

Un simple CRM (Customer Relationship Management) en ligne de commande développé en Go from scratch

---

## Fonctionnalités

- **Ajouter un contact** : Créer un nouveau contact avec ID, nom et email  
- **Lister les contacts** : Afficher tous les contacts enregistrés  
- **Supprimer un contact** : Supprimer un contact par son ID  
- **Mettre à jour un contact** : Modifier le nom et/ou l'email d'un contact existant  
- **Support des flags** : Ajouter un contact directement via la ligne de commande

---

## Installation et utilisation

### Prérequis (dév sous windows)

- [Go](https://go.dev/) 

### Installation

1. Clonez le projet :
```bash
git clone https://github.com/BanggEddy/golangCLIscratch.git
cd golangCLIscratch
```

2. Lancez le programme dans le dossier du projet :
```bash
go run .
```

3. Utilisation rapide avec flags
```bash
go run . --nom="Alice" --email="alice@example.com"
```

