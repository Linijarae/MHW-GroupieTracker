package controllers

import (
	"fmt"
	"mhw/src/services"
	temp "mhw/src/templates"
	"net/http"
	"strconv"
)

func PageListMonster(w http.ResponseWriter, r *http.Request) {
    listMonster, listMonsterCode, listMonsterErr := services.GetListMonster()

    if listMonsterErr != nil {
        fmt.Println(listMonsterErr.Error())
        fmt.Println(listMonsterCode)
        http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=Erreur lors de la récupération des monstres", listMonsterCode), http.StatusPermanentRedirect)
        return
    }
    
    // Correction ici : utiliser listMonster directement
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

func PasMain(w http.ResponseWriter, r *http.Request) {
	/*listMonster, listMonsterCode, listMonsterErr := */ services.PasMainMonster()

	/*if listMonsterErr != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?code=%d&message=Erreur lors de la récupération des monstres", listMonsterCode), http.StatusPermanentRedirect)
		return
	}
	temp.Temp.ExecuteTemplate(w, "listMonster", listMonster)*/
}
