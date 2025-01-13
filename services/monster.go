package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ListMonster struct {
    Results []Monsters `json:"results"`
}

type Monsters struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Species  string   `json:"species"`
	Type     string   `json:"type"`
	Elements []string `json:"elements"`
}

func GetListMonster() (ListMonster, int, error) {
	fmt.Println("Hello World")
	_client := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "https://mhw-db.com/monsters", nil)
	if reqErr != nil {
		return ListMonster{}, 500, fmt.Errorf("Erreur Requête - Une erreur lors de la préparation de la requête 'GetListMonster' : %v", reqErr.Error())
	
	}

	res, resErr := _client.Do(req)
	if resErr != nil {
		fmt.Println("Hello World 2")
		return ListMonster{}, 500, fmt.Errorf("Erreur Requête - Une erreur s'est produite lors de l'exécution de la requête 'GetListMonster' : %v", resErr.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Hello World 3")
		return ListMonster{}, res.StatusCode, fmt.Errorf("Erreur réponse - code %s", res.Status)
	}

	var list ListMonster
	decodeErr := json.NewDecoder(res.Body).Decode(&list)
	if decodeErr != nil {
		fmt.Println("Hello World 4")
		fmt.Println(ListMonster{})
		fmt.Println(list)
		return ListMonster{}, 500, fmt.Errorf("Erreur Décodage Des Données : Une erreur s'est produite lors de la lecture des données 'GetListMonster' : %v", decodeErr.Error())
	}
	fmt.Println("Hello World 5")
	return list, res.StatusCode, nil
}



func GetMonsterById(id int) (Monsters, int, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://mhw-db.com/monsters/%v", id), nil)
	if reqErr != nil {
		return Monsters{}, 500, fmt.Errorf("Erreur Requête - Une erreur lors de la préparation de la requête 'GetMonsterById' : %v", reqErr.Error())
	}

	res, resErr := client.Do(req)
	if resErr != nil {
		return Monsters{}, 500, fmt.Errorf("Erreur Requête - Une erreur s'est produite lors de l'exécution de la requête 'GetMonsterById' : %v", resErr.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Monsters{}, res.StatusCode, fmt.Errorf("Erreur réponse - code %s", res.Status)
	}

	var monster Monsters
	decodeErr := json.NewDecoder(res.Body).Decode(&monster)
	if decodeErr != nil {
		return Monsters{}, 500, fmt.Errorf("Erreur Décodage Des Données : Une erreur s'est produite lors de la lecture des données 'GetMonsterById' : %v", decodeErr.Error())
	}

	return monster, 200, nil
}
