package function

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func RandomMenus(w http.ResponseWriter, r *http.Request) {
	days, err := strconv.Atoi(r.URL.Query().Get("days"))
	if err != nil || days < 1 || days > 5 {
		days = 5
	}

	log.Printf("Fetch random menus for %d days", days)

	menus := []Menu{}
	menus = append(menus, Menu{Title: "Big Mac"}, Menu{Title: "Gemüsepfanne"}, Menu{Title: "Gegrillter Lachs"},
		Menu{Title: "Spinat mit Kartoffeln"}, Menu{Title: "Flammkuchen"}, Menu{Title: "Vogerlsalat"}, Menu{"Backfisch"}, Menu{"Wurstsalat"},
		Menu{Title: "Couscous Salat"}, Menu{Title: "Gefüllte Zuchini"}, Menu{Title: "Kichererbsenpfanne"}, Menu{Title: "Süßkartoffelsalat"})

	randomMenus := make(map[string]Menu)

	for len(randomMenus) < days {
		rand.Seed(time.Now().UnixNano())

		menu := menus[rand.Intn(len(menus))]
		randomMenus[menu.Title] = menu
	}

	returnValues := []Menu{}
	for _, value := range randomMenus {
		returnValues = append(returnValues, value)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(returnValues)
}
