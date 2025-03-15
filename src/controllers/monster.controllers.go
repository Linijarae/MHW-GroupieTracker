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
    // Paramètres de pagination
    page := r.URL.Query().Get("page")
    if page == "" {
        page = "1" // Valeur par défaut à la page 1 si aucun paramètre n'est fourni
    }

    // Convertir la page en entier
    currentPage, err := strconv.Atoi(page)
    if err != nil {
        currentPage = 1 // Si la conversion échoue, revenir à la page 1
    }

    // Nombre d'éléments par page
    itemsPerPage := 10

    listMonster, listMonsterCode, listMonsterErr := services.GetListMonster()
    if listMonsterErr != nil {
        fmt.Println(listMonsterErr.Error())
        fmt.Println(listMonsterCode)
        http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=Erreur lors de la récupération des monstres", listMonsterCode), http.StatusPermanentRedirect)
        return
    }

    // Calculer le nombre total de pages
    totalMonsters := len(listMonster)
    totalPages := (totalMonsters + itemsPerPage - 1) / itemsPerPage // Calcul pour avoir la dernière page

    // S'assurer que la page ne dépasse pas le nombre total de pages
    if currentPage > totalPages {
        currentPage = totalPages
    }

    // Calculer les monstres à afficher pour cette page
    startIndex := (currentPage - 1) * itemsPerPage
    endIndex := startIndex + itemsPerPage
    if endIndex > totalMonsters {
        endIndex = totalMonsters
    }
    monstersForPage := listMonster[startIndex:endIndex]

    // Passer toutes les données nécessaires au template
    temp.Temp.ExecuteTemplate(w, "listMonster", map[string]interface{}{
        "Monsters":    monstersForPage,
        "CurrentPage": currentPage,
        "TotalPages":  totalPages,
        "HasPrevious": currentPage > 1,  // Indicateur pour savoir s'il y a une page précédente
        "HasNext":     currentPage < totalPages, // Indicateur pour savoir s'il y a une page suivante
        "PreviousPage": currentPage - 1,  // Page précédente
        "NextPage":     currentPage + 1,  // Page suivante
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

	// Filtrer uniquement les monstres qui commencent par la recherche
	for _, monster := range monsters {
		if search == "" || strings.HasPrefix(strings.ToLower(monster.Name), search) {
			filteredMonsters = append(filteredMonsters, monster)
		}
	}

	fmt.Println("Nombre de monstres trouvés :", len(filteredMonsters))

	// Préparer les données pour le template
	data := struct {
		Monsters []services.Monsters
	}{
		Monsters: filteredMonsters,
	}

	// Exécuter le template "recherche"
	err = tmpl.ExecuteTemplate(w, "recherche", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
