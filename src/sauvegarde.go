package src

import (
	"fmt"
	"os"
)

func Reprendre() {
	fmt.Println("Bonjour")
}

func save(pseudo string, lettres []string, tries int, mot string) {
	// fichier, err := os.ReadFile("save.json")

	file, err := os.OpenFile("save.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("saving game...")
	content := fmt.Sprintf("{\nName = %s\n	Lettres = %s\n	Mot = %s\nTries = %d\n}\n", pseudo, lettres, mot, tries) // mis en forme pour que ce soit lisible
	file.WriteString(content)
}
