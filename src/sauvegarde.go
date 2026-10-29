package src

import (
	"fmt"
	"os"
)

func Reprendre() /*Game*/ {
	var save []byte
	i := 0
	saveFile, err := os.ReadFile("save.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("quel nom ?") // todo, changere le text
	fmt.Scanf("%s", &save)
	for string(saveFile[i:i+5]) != "name=" { //check jusqu'a "name=" c'est le nom
		i++
	}
	i += 6
	for saveFile[i] != '\n' {
		fmt.Printf("-%c-", saveFile[i])
		i++
	}

	fmt.Printf("name = %s", save)
	//return play
}

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
