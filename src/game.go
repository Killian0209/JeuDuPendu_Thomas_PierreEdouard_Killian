package src

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var CurrentGame Game

type Game struct {
	Pseudo      string
	MotATrouver string
	Mot         string
	Essai       int
	Lettre      []string
}

func (g *Game) InitJoueur() {
	fmt.Print("Choisissez votre pseudo: ")
	fmt.Scan(&g.Pseudo)
	fmt.Println("Bonjour:", g.Pseudo, "!")
	g.Jouer()
}

func (g *Game) Jouer() error {
	fmt.Printf("Partie crée\n")

	//	 Choix du mot
	if g.MotATrouver == "" {
		mot, err := os.ReadFile("mot.txt")
		if err != nil {
			return err
		}
		ListMot := strings.Fields(string(mot))
		if len(ListMot) == 0 {
			return fmt.Errorf("mot.txt ne contient aucun mot")
		}

		randIndex := rand.Intn(len(ListMot))
		g.MotATrouver = ListMot[randIndex]
		g.Mot = strings.Repeat("_", len([]rune(g.MotATrouver)))
		g.Essai = 0
		g.Lettre = nil
	}

	//	Boucle de jeu
	for {
		fmt.Println("Mot à trouver:")
		fmt.Println("==>", g.Mot)
		fmt.Println("Entrez une lettre:")

		// Lecture de l'entrée
		var input string
		if _, err := fmt.Scan(&input); err != nil {
			return err
		}

		// Sauvegarde
		if input == "!save" || input == "!s" {
			Save(*g)
			return nil
		}

		guessRunes := []rune(strings.ToLower(input))

		// Tentative sur le mot entier (même longueur que le mot à trouver)
		if len(guessRunes) > 1 && len(guessRunes) == len([]rune(g.MotATrouver)) {
			if strings.ToLower(input) == strings.ToLower(g.MotATrouver) {
				g.Mot = g.MotATrouver
				fmt.Println("Victoire ! Le mot a été trouvé.")
				break
			}
			g.Essai += 2
			fmt.Printf("Mauvais mot ! Vous avez utilisé %d/6 essais\n", g.Essai)
			if g.Essai >= 6 {
				fmt.Println("Tous les essais ont été utilisés")
				fmt.Println("Partie Perdu !") //skill issue
				break
			}
			continue
		}

		// Vérification : une seule lettre
		if len(guessRunes) != 1 || !unicode.IsLetter(guessRunes[0]) {
			fmt.Println("L'entrée doit être une seule lettre ou essayez d'écrire le mot en entier avec le bon nombre de lettres")
			continue
		}

		// Vérification : lettre déjà proposée
		guess := string(guessRunes[0])
		dejaUtilisee := false
		for _, lettre := range g.Lettre {
			if lettre == guess {
				dejaUtilisee = true
				break
			}
		}
		if dejaUtilisee {
			fmt.Println("Vous avez déjà proposé cette lettre")
			continue
		}

		// Ajout de la lettre et affichage dans le mot
		g.Lettre = append(g.Lettre, guess)
		motTrouve := []rune(g.MotATrouver)
		motAffiche := []rune(g.Mot)
		trouve := false
		for i, lettre := range motTrouve {
			if unicode.ToLower(lettre) == guessRunes[0] {
				motAffiche[i] = lettre
				trouve = true
			}
		}
		g.Mot = string(motAffiche)

		// Mauvaise lettre : un essai en moins
		if !trouve {
			g.Essai++
			fmt.Printf("Vous avez utilisé %d/6 essais\n", g.Essai)
		}
		fmt.Println("Lettres proposées:", g.Lettre)

		// Fin de partie
		if g.Mot == g.MotATrouver {
			fmt.Println("Victoire ! Le mot a été trouvé.")
			break
		}
		if g.Essai >= 6 {
			fmt.Println("Tous les essais ont été utilisés")
			fmt.Println("Partie Perdu !") //skill issue
			break
		}
	}
	return nil
}
