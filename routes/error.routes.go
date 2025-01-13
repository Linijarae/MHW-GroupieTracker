package routes

import (
	"exemple/controllers"
	"net/http"
)

func ErrorRoutes() {
	http.HandleFunc("/error", controllers.ErrorController)
}
