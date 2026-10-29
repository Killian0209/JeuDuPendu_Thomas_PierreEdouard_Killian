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

	for {

		fmt.Print("\033[H\033[2J")

		fmt.Println("Mot à trouver:")
		fmt.Println("==>", g.Mot)
		fmt.Println("Entrez une lettre:")
		var input string
		if _, err := fmt.Scan(&input); err != nil {
			return err
		}

		if input == "!save" || input == "!s" {
			save(*g)
			return nil
		}

		guessRunes := []rune(strings.ToLower(input))
		if len(guessRunes) != 1 || !unicode.IsLetter(guessRunes[0]) {
			fmt.Println("L'entrée doit être une seule lettre")
			continue
		}

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

		if !trouve {
			g.Essai++
			fmt.Printf("Vous avez utilisé %d/6 essais\n", g.Essai)
		}
		fmt.Println("Lettres proposées:", g.Lettre)

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
