# Maze Generation Boilerplate Code
_Generate and show a maze, using the simple depth-first search algorithm._

This boilerplate contains code to display a maze but it is without a maze generation algorithm such as using a [depth-first maze generation algorithm]](https://en.wikipedia.org/wiki/Maze_generation_algorithm#Randomized_depth-first_search). This code can be implemented in the `generate()` function in the [Maze](maze/Maze.go) package. This repository is intended to provide the code you need to implement to render a maze, leaving the maze generation algorithm to be implemented later.

The current code will generate a PNG with an un-generated, un-solvable maze and a randomly-positioned green 'goal' (which actually will be the start of the maze generation but when generated will be the goal). The task is to use something like a depth-first algorithm to generate a maze. 

![Start State Maze](/images/maze.png)

A maze is represented as a `Maze` Go `struct` which is a slice of `Cell` slices:
```
type Maze struct {
	Cells [][]Cell
}
```

A `Cell` is an individual part of a maze with a route north, east, and a visited flag. A "route", representing lack of a wall, is used instead of something like `wall=true` as
the default type for a `bool` is false so by default there is no route from that cell.
```
type Cell struct {
	NorthRoute bool
	WestRoute  bool
	Visited    bool
}
```
To remove a wall between the current cell and the westerly cell you only need set `cell.WestRoute = true`. If you wish to make a route to the west then you need to find the adjacent cell (something like using its coordinates `xPosition - 1`) and setting that route, `cellAdjacent.WestRoute = true`. This saves needing to create routes in both cells.

NB: When rendering the image it will be flipped which makes west become east and north become south. This is because the x/y coordinate system has 0,0 at the top left and width,length to the bottom right.

## Running
```
go run . -width=n -height=n > output.png
```
...where `n` are integers

## Building and Running Binary

### With Go installed [https://golang.org/doc/install]:
```
go build -o maze .
./maze -width=n -height=n > output.png
```

## Running Tests
```
go test .
```
