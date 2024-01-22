package controller

import (
	"fmt"
	templates "koh-lanta/templates"
	"net/http"
)

func PrintColorResult(color string, message string) {
	colorCode := ""
	switch color {
	case "blue":
		colorCode = "\033[34m"
	default:
		colorCode = "\033[0m"
	}
	fmt.Printf("%s%s\033[0m", colorCode, message)
}

func ConnexionHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "connexion", nil)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "create", nil)
}

func AventurierHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "aventurier", nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	templates.Temp.ExecuteTemplate(w, "404", nil)
}
