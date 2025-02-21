package routes

import (
	"mhw/src/controllers"
	"net/http"
)

func MonsterRoutes() {
	http.HandleFunc("/", controllers.PageListMonster)
	http.HandleFunc("/monster", controllers.PageDetailsMonster)
	http.HandleFunc("/main", controllers.PasMain)
}
