# TicTacGo

TicTacGo is a implementation of the classic paper-and-pencil game Tic-tac-toe using [Ebitengine](https://github.com/hajimehoshi/ebiten). The player who succeeds in placing three of their marks in a horizontal, vertical, or diagonal row is the winner.

## How to play
Use the arrow keys to select a tile and press space to mark a tile.

## Features
- Single player only
- Computer doesn't play perfect on purpose to avoid all games ending in a draw

## How to build and run
To build and run TicTacGo, you will need Go 1.21 installed on your system. Once you have Go installed, follow these steps:

1. Clone the repository: `git clone https://github.com/bgaechter/tictacgo`
2. Change into the repository directory: `cd tictacgo`
3. Run the game: `go run .`
3a. Build and run the game: `go build -o ttg && ./ttg`

