package root

import (
	control "koh-lanta/controller"
	templates "koh-lanta/templates"
	"net/http"
)

func InitServer() {
	FileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", FileServer))
	http.HandleFunc("/connexion", control.ConnexionHandler)
	http.HandleFunc("/create", control.CreateHandler)
	http.HandleFunc("/aventurier", control.AventurierHandler)
	http.HandleFunc("/404", control.NotFoundHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		templates.Temp.ExecuteTemplate(w, "404", nil)
	})
}
