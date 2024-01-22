package sauvegarde

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Nom []string
var Prenom []string
var Password []string

var Nom_test []string
var Prenom_test []string
var Password_test []string

var Ntest []string
var Ptest []string
var Patest []string

func Ajout(name string, lastname string, password string) {
	Nom.append(Nom, name)
	Prenom.append(Prenom, lastname)
	Password.append(Password, password)
	SauvegarderUtilisateurs(Nom, Prenom, Password)
}

func Test(name string, lastname string, password string) {
	Ntest.append(Ntest, name)
	Ptest.append(Ptest, lastname)
	Patest.append(Patest, password)
	Verif()
	Verification()
}

func SauvegarderUtilisateurs(noms, prenoms, passwords []string) error {
	// Vérifier si la longueur des trois listes est la même
	if len(noms) != len(prenoms) || len(prenoms) != len(passwords) {
		return fmt.Errorf("Les listes de noms, prénoms et mots de passe doivent avoir la même longueur")
	}

	// Ouvrir le fichier en mode écriture, le créer s'il n'existe pas
	fichier, err := os.OpenFile("sauvegarde.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fichier.Close()

	// Parcourir les listes et écrire chaque entrée dans le fichier
	for i := 0; i < len(noms); i++ {
		// Construire la ligne à écrire dans le fichier
		ligne := fmt.Sprintf("Nom: %s, Prenom: %s, Password: %s\n", noms[i], prenoms[i], passwords[i])

		// Écrire la ligne dans le fichier
		_, err := fichier.WriteString(ligne)
		if err != nil {
			return err
		}
	}
	control.ConnexionHandler()
	fmt.Println("Sauvegarde réussie dans le fichier sauvegarde.txt")
	return nil
}

func Verif() {
	fichier, err := os.Open("sauvegarde.txt")
	if err != nil {
		return err
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		// Diviser la ligne en utilisant la virgule comme séparateur
		parts := strings.Split(scanner.Text(), ", ")

		// Extraire les valeurs de Nom, Prenom et Password
		for _, part := range parts {
			if strings.HasPrefix(part, "Nom: ") {
				Nom_test = append(Nom_test, strings.TrimPrefix(part, "Nom: "))
			} else if strings.HasPrefix(part, "Prenom: ") {
				Prenom_test = append(Prenom_test, strings.TrimPrefix(part, "Prenom: "))
			} else if strings.HasPrefix(part, "Password: ") {
				Password_test = append(Password_test, strings.TrimPrefix(part, "Password: "))
			}
		}
	}
}

func Verification() bool {
	for i := 0; i < len(Nom_test); i++ {
		if Ntest[i] == Nom_test[i] {
			Info.N = Ntest[i]
			Info.P = Ptest[i]
			Info.Pass = Patest[i]
			return true
		}
	}
}

type Info struct {
	N    string
	P    string
	Pass string
}
