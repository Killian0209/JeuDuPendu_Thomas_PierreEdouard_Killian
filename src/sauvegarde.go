package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Reprendre() {
	contenu, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("Aucune sauvegarde trouvée")
		return
	}
	lignes := strings.Split(string(contenu), "\n") // une case du tableau = une ligne du fichier

	fmt.Print("Quel est ton pseudo ? ")
	var nom string
	fmt.Scan(&nom)

		var partie Game
	trouve := false

	for i := 0; i < len(lignes)-3; i++ {
		mots := strings.Fields(lignes[i]) // "name= exemple" devient ["name=", "exemple"]
		if len(mots) == 2 && mots[0] == "name=" && mots[1] == nom {
			partie.Pseudo = nom

			// ligne lettres= [a b c d e]
			for _, m := range strings.Fields(lignes[i+1])[1:] {
				m = strings.Trim(m, "[]") // on enlève les crochets
				if m != "" {
					partie.Lettre = append(partie.Lettre, m)
				}
			}

			// ligne mot= banane
			partie.MotATrouver = strings.Fields(lignes[i+2])[1]

			// ligne tries= 3
			partie.Essai, _ = strconv.Atoi(strings.Fields(lignes[i+3])[1])

			trouve = true
		}
	}

	if !trouve {
		fmt.Println("Aucune sauvegarde pour ce pseudo")
		return
	}

		// pour chaque lettre du mot : si elle a été proposée on l'affiche, sinon "_"
	partie.Mot = ""
	for _, l := range partie.MotATrouver {
		dejaTrouvee := false
		for _, lettre := range partie.Lettre {
			if lettre == string(l) {
				dejaTrouvee = true
			}
		}
		if dejaTrouvee {
			partie.Mot += string(l)
		} else {
			partie.Mot += "_"
		}
	}

	CurrentGame = partie
	CurrentGame.Jouer()
}
	

func Save(play Game) {
	file, err := os.OpenFile("save.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("saving game...")
	content := fmt.Sprintf("{\nname= %s\nlettres= %s\nmot= %s\ntries= %d\n}\n", play.Pseudo, play.Lettre, play.MotATrouver, play.Essai) // mis en forme pour que ce soit lisible
	file.WriteString(content)
}

// renvoie true si le pseudo est déjà dans save.json
func pseudoUtilise(pseudo string) bool {
	contenu, err := os.ReadFile("save.json")
	if err != nil {
		return false
	}

	lignes := strings.Split(string(contenu), "\n")
	for _, ligne := range lignes {
		mots := strings.Fields(ligne)
		if len(mots) == 2 && mots[0] == "name=" && mots[1] == pseudo {
			return true
		}
	}
	return false
}