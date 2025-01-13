package main

import (
	routes "exemple/routes"
	temp "exemple/templates"
)

func main() {
	temp.InitTemplates()
	routes.InitServe()
}
