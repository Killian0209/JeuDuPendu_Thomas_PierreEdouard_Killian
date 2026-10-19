package src

import (
	"fmt"
	"os"
)

func Jouer() {
	fmt.Printf("Partie crée\n")
	mot, err := os.ReadFile("mot.txt")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", mot)
}
