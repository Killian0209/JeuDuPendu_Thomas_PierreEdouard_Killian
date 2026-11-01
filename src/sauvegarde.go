package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Reprendre() {
	contenu, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("Aucune sauvegarde trouvée")
		return
	}
	lignes := strings.Split(string(contenu), "\n")

	fmt.Print("Quel est ton pseudo ? ")
	var nom string
	fmt.Scan(&nom)

		var partie Game
	trouve := false

	for i := 0; i < len(lignes)-3; i++ {
		mots := strings.Fields(lignes[i])
		if len(mots) == 2 && mots[0] == "name=" && mots[1] == nom {
			partie.Pseudo = nom

			// ligne lettres
			for _, m := range strings.Fields(lignes[i+1])[1:] {
				m = strings.Trim(m, "[]")
				if m != "" {
					partie.Lettre = append(partie.Lettre, m)
				}
			}

			// ligne mot
			position, _ := strconv.Atoi(strings.Fields(lignes[i+2])[1])
			motFichier, _ := os.ReadFile("mot.txt")
			listeMots := strings.Fields(string(motFichier))
			if position < len(listeMots) {
				partie.MotATrouver = listeMots[position]
			}

			// ligne tries
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
		if dejaTrouvee || !unicode.IsLetter(l) {
			partie.Mot += string(l) // tt ce qui est - ou ' etc...
		} else {
			partie.Mot += "_"
		}
	}

	CurrentGame = partie
	CurrentGame.Jouer()
}
	

func Save(play Game) {
	contenu, err := os.ReadFile("mot.txt")
	if err != nil {
		panic(err)
	}
	listeMots := strings.Fields(string(contenu))
	position := 0
	for i, m := range listeMots {
		if m == play.MotATrouver {
			position = i
		}
	}

	file, err := os.OpenFile("save.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Printf("saving game...")
	content := fmt.Sprintf("{\nname= %s\nlettres= %s\nmot= %d\ntries= %d\n}\n", play.Pseudo, play.Lettre, position, play.Essai) // mot= devient un numéro
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