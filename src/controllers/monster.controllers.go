package controllers

import (
	"fmt"
	"html/template"
	"mhw/src/services"
	temp "mhw/src/templates"
	"net/http"
	"strconv"
	"strings"
)

func PageListMonster(w http.ResponseWriter, r *http.Request) {
	listMonster, listMonsterCode, listMonsterErr := services.GetListMonster()

	if listMonsterErr != nil {
		fmt.Println(listMonsterErr.Error())
		fmt.Println(listMonsterCode)
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=Erreur lors de la récupération des monstres", listMonsterCode), http.StatusPermanentRedirect)
		return
	}
	temp.Temp.ExecuteTemplate(w, "listMonster", map[string]interface{}{
		"Monsters": listMonster,
	})
}

func PageDetailsMonster(w http.ResponseWriter, r *http.Request) {
	queryId := r.FormValue("id")
	idMonster, idMonsterErr := strconv.Atoi(queryId)
	if queryId == "" || idMonsterErr != nil || idMonster <= 0 {
		http.Redirect(w, r, fmt.Sprintf("/error?code%d&message=%s", 400, "erreur données manquantes ou invalides : id du monstre"), http.StatusPermanentRedirect)
		return
	}
	Monster, MonsterCode, MonsterErr := services.GetMonsterById(idMonster)
	if MonsterErr != nil {
		if MonsterCode == 404 {
			http.Redirect(w, r, fmt.Sprintf("/error?code%d&message=%s%d", 404, "erreur monstre intouvable id :", idMonster), http.StatusPermanentRedirect)
		} else {
			http.Redirect(w, r, fmt.Sprintf("/error?code%d&message=%s", 400, "erreur lors de la récupération des détails du monstre"), http.StatusPermanentRedirect)
		}
		return
	}
	temp.Temp.ExecuteTemplate(w, "Monster", Monster)
}

func About(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "About", nil)
}

func Favoris(w http.ResponseWriter, r *http.Request) {
	temp.Temp.ExecuteTemplate(w, "Favoris", nil)
}
var tmpl = template.Must(template.ParseFiles("./src/templates/recherche.html"))

// Handler pour la recherche de monstres
func SearchMonsters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🚀 Handler SearchMonsters appelé !")
	// Récupération des monstres depuis l'API
	monsters, statusCode, err := services.GetListMonster()
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	// Récupération du terme de recherche
	search := strings.ToLower(r.URL.Query().Get("search"))
	fmt.Println("🔎 Terme recherché :", search)

	// Initialiser un tableau pour stocker les monstres filtrés
	var filteredMonsters []services.Monsters

	// Si une recherche est effectuée, filtrer les résultats
	if search != "" {
		fmt.Println("⚠️ Aucun terme de recherche fourni")
		for _, monster := range monsters {
			if strings.Contains(strings.ToLower(monster.Name), search) {
				filteredMonsters = append(filteredMonsters, monster)
			}
		}
	}
	if len(filteredMonsters) == 0 {
		filteredMonsters = []services.Monsters{}
	}
	
	fmt.Println("Nombre de monstres trouvés :", len(filteredMonsters)) 

	data := struct {
		Monsters []services.Monsters
	}{
		Monsters: filteredMonsters,
	}
	err = tmpl.ExecuteTemplate(w, "recherche", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}