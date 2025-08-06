# Conway's Game of Life

This is a simple implementation of Conway's Game of Life written in Go.

## Usage

### Run Game 
`make`: builds the program as a `wasm` and deploys it to a local server on `localhost:9090`

`make test`: runs the program directly with `go run`

### Controls
`click`: toggle a square alive or dead

`p`: toggles the paused state. The game starts paused.

`1-5`: set the game speed 1 is the slowest and 5 is uncapped

# Specs 
The specs can be easily tweaked for different board / screen sizes. The board is set to 128 by 128 with a screen size of 1024px.

The game state is stored in a 2D array and is iterated over with each tick. This is fine for small grids but quite inefficient for very large board sizes. In the future I'm looking into storing the game state with a map that only stores alive cells or even better, a quad tree that allows the program to skip large chunks of empty space.