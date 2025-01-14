package services

import (
	"encoding/json"
	"fmt"
	"io"
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
	url := "https://mhw-db.com/monsters"
	method := "GET"
	_client := &http.Client{Timeout: time.Second * 5,}
	req, err := http.NewRequest(method, url, nil)
	
	if err != nil {
		return ListMonster{}, 500, fmt.Errorf("erreur Requête - Une erreur lors de la préparation de la requête 'GetListMonster' : %v", err.Error())
	}

	res, err := _client.Do(req)
	if err != nil {
		return ListMonster{}, 500, fmt.Errorf("erreur Requête - Une erreur s'est produite lors de l'exécution de la requête 'GetListMonster' : %v", err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
	  fmt.Println(err)
	  return ListMonster{}, 500, fmt.Errorf("erreur Body - Une erreur s'est produite lors du body body body : %v", err.Error())
	}
	fmt.Println(string(body))

	if res.StatusCode != http.StatusOK {
		return ListMonster{}, res.StatusCode, fmt.Errorf("erreur réponse - code %s", res.Status)
	}

	var list ListMonster
	decodeErr := json.NewDecoder(res.Body).Decode(&list)
	if decodeErr != nil {
		return ListMonster{}, 500, fmt.Errorf("erreur Décodage Des Données : Une erreur s'est produite lors de la lecture des données 'GetListMonster' : %v", decodeErr.Error())
	}
		
	return list, res.StatusCode, nil

}

func GetMonsterById(id int) (Monsters, int, error) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	req, reqErr := http.NewRequest(http.MethodGet, fmt.Sprintf("https://mhw-db.com/monsters/%v", id), nil)
	if reqErr != nil {
		return Monsters{}, 500, fmt.Errorf("erreur Requête - Une erreur lors de la préparation de la requête 'GetMonsterById' : %v", reqErr.Error())
	}

	res, resErr := client.Do(req)
	if resErr != nil {
		return Monsters{}, 500, fmt.Errorf("erreur Requête - Une erreur s'est produite lors de l'exécution de la requête 'GetMonsterById' : %v", resErr.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return Monsters{}, res.StatusCode, fmt.Errorf("erreur réponse - code %s", res.Status)
	}

	var monster Monsters
	decodeErr := json.NewDecoder(res.Body).Decode(&monster)
	if decodeErr != nil {
		return Monsters{}, 500, fmt.Errorf("erreur Décodage Des Données : Une erreur s'est produite lors de la lecture des données 'GetMonsterById' : %v", decodeErr.Error())
	}

	return monster, 200, nil
}

func PasMainMonster() {

	url := "https://mhw-db.com/monsters"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
