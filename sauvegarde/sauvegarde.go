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

func Ajout(name string, lastname string, password string) {
	Nom.append(Nom, name)
	Prenom.append(Prenom, lastname)
	Password.append(Password, password)

	file, err := os.Create("sauvegarde.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	ecrit := bufio.NewWriter(file)

	for cle, valeur := range Save {
		_, err := fmt.Fprintf(ecrit, "%s: %s\n", cle, valeur)
		if err != nil {
			fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
			return
		}
	}

	ecrit.Flush()
}

func Verif(name string, lastname string, password string) {
	utilisateurs, err := chargerUtilisateursDepuisFichier("sauvegarde.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement des utilisateurs :", err)
		return
	}

	// Exemple d'utilisation de la fonction pour vérifier si un utilisateur est enregistré
	ok := verifierUtilisateur(utilisateurs, name, lastname)
	if ok {
		fmt.Println("L'utilisateur est enregistré.")
	} else {
		fmt.Println("L'utilisateur n'est pas enregistré.")
	}
}

func ChargerUtilisateursDepuisFichier(nomFichier string) (map[string]Utilisateur, error) {
	utilisateurs := make(map[string]Utilisateur)

	fichier, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		ligne := scanner.Text()
		elements := strings.Split(ligne, ": ")

		if len(elements) == 2 {
			nomPrenom := elements[0]
			password := elements[1]
			nomPrenomElements := strings.Split(nomPrenom, " ")
			if len(nomPrenomElements) == 2 {
				prenom := nomPrenomElements[0]
				nom := nomPrenomElements[1]
				utilisateur := Utilisateur{Nom: nom, Prenom: prenom, Password: password}
				utilisateurs[nomPrenom] = utilisateur
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return utilisateurs, nil
}

// Vérifier si un utilisateur est enregistré dans la carte
func VerifierUtilisateur(utilisateurs map[string]Utilisateur, prenom, nom, password string) bool {
	cle := fmt.Sprintf("%s %s", prenom, nom)
	utilisateur, ok := utilisateurs[cle]
	return ok && utilisateur.Password == password
}
