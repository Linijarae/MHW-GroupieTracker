package routes

import (
	"fmt"
	"mhw/src/controllers"
	"net/http"
)

func InitServe() {
	
	http.HandleFunc("/error", controllers.ErrorController)
	http.HandleFunc("/", controllers.PageListMonster)
	http.HandleFunc("/monster", controllers.PageDetailsMonster)
	http.HandleFunc("/favoris", controllers.Favoris)
	http.HandleFunc("/about", controllers.About)

	fmt.Println("Le serveur est opérationel : http://localhost:8000/")
	http.ListenAndServe("localhost:8000", nil)
}
