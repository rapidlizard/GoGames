package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Game struct
type Game struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

//Genre struct
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getGames() []Game {
	var games []Game

	db, err := sql.Open("mysql", "root:kane@(127.0.0.1:3306)/gogames")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("success!!")

	results, err := db.Query("SELECT * FROM games")
	if err != nil {
		panic(err)
	}
	for results.Next() {
		var game Game

		results.Scan(&game.ID, &game.Title)
		games = append(games, game)
	}

	return games
}

func returnGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getGames())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/games", returnGames).Methods("GET")

	log.Fatal(http.ListenAndServe(":4444", r))
}
