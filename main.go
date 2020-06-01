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
	GameID   int    `json:"game_id"`
	GameName string `json:"game_name"`
	Genre    Genre  `json:"Genre"`
}

//Genre struct
type Genre struct {
	GenreID   int    `json:"genre_id"`
	GenreName string `json:"genre_name"`
}

func getGames() []Game {
	var games []Game

	db, err := sql.Open("mysql", "root:kane@(127.0.0.1:3306)/gogames")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("success!!")

	rows, err := db.Query("SELECT game_id, game_name, genre_id, genre_name FROM games LEFT JOIN genres ON genre_id = genre_fk")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var game Game

		rows.Scan(&game.GameID, &game.GameName, &game.Genre.GenreID, &game.Genre.GenreName)

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
