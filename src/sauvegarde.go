package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Reprendre() {
	saveFile, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println("Aucune sauvegarde trouvée")
		return
	}
	lignes := strings.Split(string(saveFile), "\n")

	fmt.Print("Quel est ton pseudo ? ")
	var nom string
	fmt.Scan(&nom)

func save(play Game) {
	fmt.Printf("%v", play)
	file, err := os.OpenFile("save.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("saving game...")
	content := fmt.Sprintf("{\nname= %s\nlettres= %s\nmot= %s\ntries= %d\n}\n", play.Pseudo, play.Lettre, play.MotATrouver, play.Essai) // mis en forme pour que ce soit lisible
	file.WriteString(content)
}
