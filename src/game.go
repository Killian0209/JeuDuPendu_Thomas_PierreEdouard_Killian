package src

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var partie Game

var lettre string

func (j *Joueur) InitJoueur() {
	fmt.Println("Choisissez votre pseudo: ")
	fmt.Scan(&j.Pseudo)
	fmt.Println("Bonjour: ", j.Pseudo, "!")
	partie.Jouer()
}

type Game struct {
	Pseudo      string
	MotATrouver string
	Mot         string
	Essai       int
	Lettre      []string
}

func (g *Game) Jouer() error {
	//compteurEssai := 0
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

		if (lettre >= "A" && lettre <= "Z") {
			lettre = strings.ToLower(lettre)
		}

		//for k,_ := range g.MOt 

		if len(lettre) != 1 {
			fmt.Println("l'entrée doit être d'une seule lettre")
			fmt.Println("Veuiller entrer une nouvelle lettre: ")
			fmt.Scan(&lettre)
		}

		dejaUtilisee := false

		for l := range g.Lettre {
			if g.Lettre[l] == lettre {
				dejaUtilisee = true
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

		if g.Mot == g.MotATrouver {
			fmt.Println("Victoire ! Le mot a été trouvé.")
			break
		}
		
	}
	return nil
}
