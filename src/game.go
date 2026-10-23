package src

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var partie Game

var lettre string
var list_Lettres []string

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
	fmt.Println("Bonjour: ", g.Pseudo, "!")
	partie.Jouer(g.Pseudo)
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
			save(pseudo, list_Lettres, compteurEssai) // envoie les donnée de la partie
			return err
		}

		if lettre >= "A" && lettre <= "Z" {
			lettre = strings.ToLower(lettre) // met en minuscule les lettres en maj
		}
		
		if len(lettre) != 1 {
			fmt.Println("l'entrée doit être d'une seule lettre") //verifie si l'entrée est de qu'une seule lettre
			compteurEssai --
			continue
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

		if compteurEssai < 0 {
			compteurEssai = 0
		}

		if compteurEssai >= 6 {
			fmt.Println("Tous les essais ont été utilisés")
			fmt.Println("Partie Perdu !")
			break
		}

		for k := range g.Mot {
			if lettre == string(g.Mot[k]) {
				compteurEssai --
			}
			if lettre[0] != g.Mot[k] {
				list_Lettres = append(list_Lettres, lettre)
				compteurEssai ++
				fmt.Printf("Vous avez utilisé %d/6 essais ", compteurEssai)
				break
			}
		}

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
