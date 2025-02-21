package routes

import (
	"mhw/src/controllers"
	"net/http"
)

func ErrorRoutes() {
	http.HandleFunc("/error", controllers.ErrorController)
}
