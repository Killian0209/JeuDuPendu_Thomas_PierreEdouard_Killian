# Jeu du Pendu

Un jeu du pendu en ligne de commande, écrit en Go.
Projet réalisé par Thomas, Pierre-Édouard et Killian.

## Fonctionnalités

- Un mot est tiré au hasard dans `mot.txt`
- 6 essais maximum avant de perdre
- Proposition d'une lettre ou du mot entier
- Sauvegarde et reprise d'une partie en cours
- Pseudo unique pour chaque joueur
- Historique des scores (victoire ou défaite, essais, mot, date)
- Les caractères `-` et `'` sont affichés directement dans le mot

## Prérequis

- [Go] installé sur ta machine

## Lancer le jeu

Depuis le dossier du projet :

| Commande | Action |
|---|---|
| `go run . jouer` | Lance une nouvelle partie |
| `go run . reprendre` | Reprend une partie sauvegardée |
| `go run . scores` | Affiche les scores |

Sans argument, le programme affiche les commandes disponibles.

## Règles du jeu

1. Entre ton pseudo (il doit être différent de ceux déjà utilisés dans les sauvegardes).
2. Le mot à trouver s'affiche avec des `_` à la place des lettres.
3. À chaque tour, propose **une lettre** :
   - bonne lettre : elle apparaît dans le mot
   - mauvaise lettre : tu perds **1 essai**
   - lettre déjà proposée : rien ne se passe, tu peux rejouer
4. Tu peux aussi proposer **le mot entier**, à condition qu'il ait le même nombre de lettres que le mot à trouver :
   - mot correct : victoire immédiate
   - mot faux : tu perds **2 essais**
5. Tu perds quand tu as utilisé les **6 essais**.

## Sauvegarder et reprendre

- Pendant une partie, tape `!save` (ou `!s`) à la place d'une lettre pour sauvegarder et quitter.
- La sauvegarde est écrite dans `save.json`.
- Pour reprendre : `go run . reprendre`, puis entre ton pseudo. Si ce pseudo a plusieurs sauvegardes, c'est la dernière qui est reprise.

## Scores

À la fin de chaque partie (victoire ou défaite), une ligne est ajoutée dans `scores.csv` :

```
pseudo;statut;essais;mot;date
kiki;victoire;3;banane;01/11/2026 14:32
```

`go run . scores` affiche toutes les parties enregistrées. Une partie sauvegardée avec `!save` n'est pas comptée tant qu'elle n'est pas terminée.

## Ajouter des mots

Ouvre `mot.txt` et ajoute un mot par ligne, sans espace. Les mots peuvent contenir des `-` ou des `'`. Évite les accents pour que les joueurs n'aient pas à les taper.

## Structure du projet

```
.
├── main.go            # point d'entrée, gère les commandes
├── mot.txt            # liste des mots possibles
├── save.json          # parties sauvegardées
├── scores.csv         # historique des scores
├── README.md
└── src/
    ├── game.go        # logique de la partie
    ├── sauvegarde.go  # sauvegarde et reprise
    └── scores.go      # enregistrement et affichage des scores
```

## Auteurs

- Thomas
- Pierre-Édouard
- Killian