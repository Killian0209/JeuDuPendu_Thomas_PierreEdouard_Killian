package src

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var CurrentGame Game

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
	fmt.Println("Bonjour:", g.Pseudo, "!")
	g.Jouer()
}

func (g *Game) Jouer() error {
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

	fmt.Printf("1 : %v\n", g)

	for {
		currentPlay := Game{g.Pseudo, g.MotATrouver, g.Mot, compteurEssai, list_Lettres}
		//debug(
		fmt.Printf("for 1 : %v\n", g)
		fmt.Printf("for 2 : %v\n", currentPlay)
		//debug)

		//ajoute toutes les données a la structure
		fmt.Println("Mot à trouver: ")
		fmt.Println("==>", g.Mot)
		fmt.Println("Entrez une lettre: ")
		fmt.Scan(&lettre)
		if lettre == "!save" || lettre == "!s" {
			save(currentPlay) // envoie les donnée de la partie
			return err
		}

		if lettre >= "A" && lettre <= "Z" {
			lettre = strings.ToLower(lettre) // met en minuscule les lettres en maj
		}

		if len(lettre) != 1 {
			fmt.Println("l'entrée doit être d'une seule lettre") //verifie si l'entrée est de qu'une seule lettre
			continue
		}

		dejaUtilisee := false

		for l := range g.Lettre {
			if g.Lettre[l] == lettre {
				dejaUtilisee = true
			}
		}

		for k := range g.MotATrouver {
			if lettre[0] != g.MotATrouver[k] {
				list_Lettres = append(list_Lettres, lettre)
				compteurEssai++
				fmt.Printf("Vous avez utilisé %d/6 essais ", compteurEssai)
				break
			}
		}

		if dejaUtilisee == false && (rune(lettre[0]) >= 'a' && rune(lettre[0]) <= 'z') {
			currentPlay.Lettre = append(g.Lettre, lettre)
		}
		fmt.Println(list_Lettres) //liste de lettre deja proposé

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
				compteurEssai--
			}
			if lettre[0] != g.Mot[k] {
				list_Lettres = append(list_Lettres, lettre)
				compteurEssai++
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
