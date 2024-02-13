package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

type cell struct {
	isBomb       bool
	adjBombCount int
	visible      bool
}

var showMines = true

var grid [15][15]cell

var clear map[string]func()
var redColor = color.New(color.FgRed)
var whiteColor = color.New(color.FgWhite)
var greenColor = color.New(color.FgGreen)

//var bombs = make(map[[2]int]bool)

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {
	CallClear()
	gameOver := false
	var input1 int
	var input2 int
	addBombsToGrid()
	printGrid()
	for !gameOver {
		fmt.Scan(&input1)
		fmt.Scan(&input2)
		if !grid[input1][input2].isBomb{
			revealArea(input1, input2)
			CallClear()
		} else {
			redColor.Println("GAME OVER")
			gameOver = true
		}

		if !gameOver {
			printGrid()
		}
	}
}

func revealArea(r int, c int) {
	//TO-DO
}

func randBool() bool {
	return rand.Int()%11 == 1
}

func addBombsToGrid() {
	for row := range grid {
		for c := range grid[row] {
			if randBool() {
				grid[row][c].isBomb = true
				//	temp := [2]int{row, c}
				//bombs[temp] = true
				addAdjbombCount(row, c)
			}
		}
	}
	fmt.Println()
}

func addAdjbombCount(r int, c int) {
	//for k := range bombs {
	//	r := k[0]
	//	c := k[1]
	if c+1 < 15 && !grid[r][c+1].isBomb{
		grid[r][c+1].adjBombCount++
	}
	if c-1 >= 0 && !grid[r][c-1].isBomb{
		grid[r][c-1].adjBombCount++
	}
	if r+1 < 15 && !grid[r+1][c].isBomb{
		grid[r+1][c].adjBombCount++
	}
	if r-1 >= 0 && !grid[r-1][c].isBomb{
		grid[r-1][c].adjBombCount++
	}
	if c+1 < 15 && r+1 < 15 && !grid[r+1][c+1].isBomb{
		grid[r+1][c+1].adjBombCount++
	}
	if r+1 < 15 && c-1 >= 0 && !grid[r+1][c-1].isBomb{
		grid[r+1][c-1].adjBombCount++
	}
	if c+1 < 15 && r-1 >= 0 && !grid[r-1][c+1].isBomb{
		grid[r-1][c+1].adjBombCount++
	}
	if c-1 >= 0 && r-1 >= 0 && !grid[r-1][c-1].isBomb{
		grid[r-1][c-1].adjBombCount++
	}
	// }
}

func printGrid() {
	for row := range grid {
		for _,colValue := range grid[row] {
			if colValue.visible {
				if colValue.adjBombCount > 0 {
					greenColor.Print(colValue.adjBombCount, " ")
				} else {
					whiteColor.Print(0, " ")
				}
			} else {
				whiteColor.Print(0, " ")
			}
		//	if grid[row][c].isBomb {
		//		redColor.Print(0, " ")
		//	} else if grid[row][c].adjBombCount>0  {
		//		greenColor.Print(colValue.adjBombCount, " ")
		//	} else {
		//		whiteColor.Print(0, " ")
		//	}
		}
		fmt.Println()

	}
}
