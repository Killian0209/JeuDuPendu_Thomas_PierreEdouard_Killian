package main

import (
	"Pendu/src"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "jouer":
			src.Jouer()
		case "reprendre":
			src.Reprendre()
		case "scores":
			src.Scores()
		default:
			fmt.Println("Commande incorrecte")
		}
	} else {
		fmt.Println("il faut écrire pour commencer: go run . jouer")
		fmt.Println("il faut écrire pour reprendre: go run . reprendre")
		fmt.Println("il faut écrire pour scores: go run . scores")
	}
}
