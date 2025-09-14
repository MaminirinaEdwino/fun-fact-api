package controller

import (
	"fmt"
	"os"
)

func ReadTemplate(templateDir string) string{
	file, err := os.ReadFile(templateDir)
	if err != nil {
		fmt.Println("Erreur pendant l'ouverture duf fichier")
		return "<div>Erreur</div>"
	}
	return  string(file[:])
}