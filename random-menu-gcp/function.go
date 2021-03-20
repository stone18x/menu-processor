package function

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	firebase "firebase.google.com/go/v4"
)

func RandomMenu(w http.ResponseWriter, r *http.Request) {
	days, err := strconv.Atoi(r.URL.Query().Get("days"))
	if err != nil || days < 1 || days > 5 {
		days = 5
	}

	log.Printf("Fetch random menu for %d days", days)

	dishes, err := fetchRandomDishes(days)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode("Error!")
	} else {
		json.NewEncoder(w).Encode(dishes)
	}
}

func fetchRandomDishes(days int) ([]Dish, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: os.Getenv("FIRESTORE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer client.Close()

	docs, err := client.Collection("dishes").Documents(ctx).GetAll()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	// not really the most efficient way to fetch all documents and the choose some random ones but the easiest one and as long as the number of dishes is limited quite okay.
	// https://stackoverflow.com/questions/46798981/firestore-how-to-get-random-documents-in-a-collection
	// we need to think about a better way to do it anyway, e.g. consider not 3 times spaghtetti a week or the menu of the previous days
	dishes := []Dish{}
	for _, doc := range docs {
		dish := Dish{}
		if err := doc.DataTo(&dish); err != nil {
			log.Fatalf("Failed to convert: %v", err)
			continue
		}

		dish.ID = doc.Ref.ID
		dishes = append(dishes, dish)
	}

	temp := make(map[string]Dish)
	rand.Seed(time.Now().UnixNano())
	for len(temp) < days {
		randomNumber := rand.Intn(len(dishes))
		dish := dishes[randomNumber]
		temp[dish.ID] = dish
	}

	menu := []Dish{}
	for _, value := range temp {
		menu = append(menu, value)
	}
	return menu, nil
}
