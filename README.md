# My GoGames API
Just a simple crud API for games.

## How to use:

Run go build and execute the restapi file

After you can visit:
```
http://localhost:4444
```
- `/games` to get list of games (json).

## Todo:
- Refresh list of games from db everytime a GET request is made. **DONE**

- One to many relationship between game and genre (genre has many games). **DONE**

- Handle request for just 1 game.

- Handle request to store a new game.

- Handle request to update a game.

- Handle request to delete a game.
