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
		default:
			fmt.Println("Commande incorrecte")
		}
	} else {
		fmt.Println("il faut écrire pour commencer: go run . jouer")
	}
}
