package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//enregistrement du score
func enregistrerScore(play Game, resultat string) {
	file, err := os.OpenFile("scores.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Impossible d'enregistrer le score")
		return
	}
	defer file.Close()

	date := time.Now().Format("02/01/2006 15:04") // date
	ligne := play.Pseudo + ";" + resultat + ";" + strconv.Itoa(play.Essai) + ";" + play.MotATrouver + ";" + date + "\n"
	file.WriteString(ligne)
}

// affiche tous les scores (go run . scores)
func Scores() {
	contenu, err := os.ReadFile("scores.csv")
	if err != nil || strings.TrimSpace(string(contenu)) == "" {
		fmt.Println("Aucun score pour le moment")
		return
	}

	fmt.Println("Scores: ")
	lignes := strings.Split(string(contenu), "\n")
	for _, ligne := range lignes {
		morceaux := strings.Split(ligne, ";") // [pseudo, statut, essais, mot, date]
		if len(morceaux) == 5 {
			fmt.Printf("Pseudo: %s | Statut: %s | Essais: %s | Mot: %s | Date: %s\n",
				morceaux[0], morceaux[1], morceaux[2], morceaux[3], morceaux[4])
		}
	}
}