# My GoGames API
Just a simple crud API for games.

## Documentation

GET `/games` response:
```json
[{
    "game_id": 1,
    "game_name": "Skyrim",
    "Genre": {
        "genre_id": 3,
        "genre_name": "adventure"
    }
}, {
    "game_id": 2,
    "game_name": "CSGO",
    "Genre": {
        "genre_id": 1,
        "genre_name": "action"
    }
},]
```

POST /games needs to be in this format:
```json
{
    "game_name": "Skyrim",
    "Genre": {
        "genre_id": 3,
    }
}
```

## Todo:
- Refresh list of games from db everytime a GET request is made. **DONE**

- One to many relationship between game and genre (genre has many games). **DONE**

- Handle request for just 1 game.

- Handle request to store a new game. **DONE**

- Handle request to update a game.

- Handle request to delete a game.
