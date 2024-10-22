# Game of Life

Game of Life is a cellular automaton created by British mathematician John Horton Conway.

This is a zero-player game, meaning the process of the game depends entirely on its initial state.

  

## Rules

- Any live cell with fewer than two live neighbors dies, as if by underpopulation.
- Any live cell with two or three live neighbors lives on to the next generation.
- Any live cell with more than three live neighbors dies, as if by overpopulation.
- Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

  

## Project Overview

In my version of the game, a separate PNG file is created for each step of the evolution and stored in a new directory called `assets`. From these files, you can create an animation.

- The initial state of the grid is randomized for each run of the program.
- The grid is represented by a 2D array of cells.
- For each step of evolution, a separate PNG image is generated and saved in the `assets` directory.


## How to Run the Program

Make sure Golang is installed on your computer.

To run the program, execute the following command:

  

```bash
go run main.go array.go image.go
```
## Creating an Animation

To  create  an  animation  from  the  PNG  files,  you  can  use  `ffmpeg`.  Here  is  an  example  of  how  to  do  this:
```bash
ffmpeg -framerate 30 -i assets/frame_%d.png -c:v libx264 -pix_fmt yuv420p game-of-life.mp4
```
