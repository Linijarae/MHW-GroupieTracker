package main

import (
	routes "mhw/src/routes"
	temp "mhw/src/templates"
	"net/http"
)

func main() {
	http.Handle("/src/assets/", http.StripPrefix("/src/assets/", http.FileServer(http.Dir("src/assets"))))
	temp.InitTemplates()
	routes.InitServe()
}
