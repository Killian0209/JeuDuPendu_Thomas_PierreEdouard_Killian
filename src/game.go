package src

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var partie Game

var lettre string
var Lettres []byte

type Joueur struct {
	Pseudo string
}

func (j *Joueur) InitJoueur() {
	fmt.Print("Choisissez votre pseudo: ")
	fmt.Scan(&j.Pseudo)
	fmt.Println("Bonjour: ", j.Pseudo, "!")
	partie.Jouer(j.Pseudo)
}

type Game struct {
	Player      Joueur
	MotATrouver string
	Mot         string
	Essai       int
	Lettre      []string
}

func (g *Game) Jouer(pseudo string) error {
	compteurEssai := 0
	fmt.Printf("Partie crée\n")
	mot, err := os.ReadFile("mot.txt")
	if err != nil {
		return err
	}

	ListMot := strings.Fields(string(mot))
	randIndex := rand.Intn(len(ListMot))
	g.MotATrouver = ListMot[randIndex]
	g.Mot = strings.Repeat("_", len(g.MotATrouver))

	for {
		fmt.Println("Mot à trouver: ")
		fmt.Println("==>", g.Mot)
		fmt.Println("Entrez une lettre: ")
		fmt.Scan(&lettre)

		if lettre == "!save" || lettre == "!s" {
			strings.ToLower(lettre)

			save(pseudo, lettre, compteurEssai) // envoie les donnée de la partie
			return err
		}
		if len(lettre) != 1 {
			fmt.Println("l'entrée doit être d'une seule lettre")
			compteurEssai -= 1
			continue
		}

		if lettre >= "A" && lettre <= "Z" {
			lettre = strings.ToLower(lettre)
		}

		dejaUtilisee := false

		for l := range g.Lettre {
			if g.Lettre[l] == lettre {
				dejaUtilisee = true
			}
		}

		for k := range g.MotATrouver {
			if lettre[0] != g.MotATrouver[k] {
				compteurEssai += 1
				fmt.Printf("Vous avez utilisé %d/6 essais ", compteurEssai)
				break
			}
		}

		if dejaUtilisee == false && (rune(lettre[0]) >= 'a' && rune(lettre[0]) <= 'z') {
			g.Lettre = append(g.Lettre, lettre)
		}
		fmt.Println(g.Lettre)

		for i := 0; i < len(g.Mot); i++ {
			for j := 0; j < len(g.Lettre); j++ {

				if g.Lettre[j] == string(g.MotATrouver[i]) {
					motTemp := []rune(g.Mot)
					motTemp[i] = rune(g.Lettre[j][0])
					g.Mot = string(motTemp)
				}
			}
		}
		fmt.Printf("%d", compteurEssai)
		if g.Mot == g.MotATrouver {
			fmt.Println("Victoire ! Le mot a été trouvé.")
			break
		}

	}
	return nil
}
