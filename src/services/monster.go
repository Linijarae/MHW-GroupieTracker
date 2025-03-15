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
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Species     string   `json:"species"`
	Type        string   `json:"type"`
	Elements    []string `json:"elements"`
	Description string   `json:"description"`
}

func GetListMonster() ([]Monsters, int, error) {
	url := "https://mhw-db.com/monsters"
	client := &http.Client{Timeout: time.Second * 5}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 500, fmt.Errorf("erreur Requête - %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, 500, fmt.Errorf("erreur Exécution Requête - %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("erreur Réponse - code %s", res.Status)
	}

	var monsters []Monsters
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, 500, fmt.Errorf("erreur Lecture Body - %v", err)
	}

	err = json.Unmarshal(body, &monsters)
	if err != nil {
		return nil, 500, fmt.Errorf("erreur Décodage JSON - %v", err)
	}

	return monsters, res.StatusCode, nil
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
