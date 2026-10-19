package main

import (
	"fmt"
	"os"
	//"pendu/src"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "jouer":
			jouer()
		default:
			fmt.Println("Commande incorrecte")
		}
	} else {
		fmt.Println("il faut écrire pour commencer: go run . jouer")
	}
}

func jouer() 
