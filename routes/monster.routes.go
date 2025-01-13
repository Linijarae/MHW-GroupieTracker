package routes

import (
	"exemple/controllers"
	"net/http"
)

func MonsterRoutes() {
	http.HandleFunc("/", controllers.PageListMonster)
	http.HandleFunc("/monster", controllers.PageDetailsMonster)
}
