package main

import (
	control "koh-lanta/controller"
	root "koh-lanta/root"
	templates "koh-lanta/templates"
	"net/http"
)

func main() {
	control.PrintColorResult("blue", " http://localhost:8080/connexion \n")
	templates.InitTemplate()
	root.InitServer()
	http.ListenAndServe("localhost:8080", nil)
}
