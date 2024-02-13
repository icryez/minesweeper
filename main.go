package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var showMines = true

var grid [15][15]int

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
		if grid[input1][input2] != 1 && grid[input1][input2] != 2 {
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

}

func randBool() bool {
	return rand.Int()%11 == 1
}

func addBombsToGrid() {
	for row := range grid {
		for c := range grid[row] {
			if randBool() {
				grid[row][c] = 9
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
	if c+1 < 15 && grid[r][c+1] != 9 {
		grid[r][c+1]++
	}
	if c-1 >= 0 && grid[r][c-1] != 9 {
		grid[r][c-1]++
	}
	if r+1 < 15 && grid[r+1][c] != 9 {
		grid[r+1][c]++
	}
	if r-1 >= 0 && grid[r-1][c] != 9 {
		grid[r-1][c]++
	}
	if c+1 < 15 && r+1 < 15 && grid[r+1][c+1] != 9 {
		grid[r+1][c+1]++
	}
	if r+1 < 15 && c-1 >= 0 && grid[r+1][c-1] != 9 {
		grid[r+1][c-1]++
	}
	if c+1 < 15 && r-1 >= 0 && grid[r-1][c+1] != 9 {
		grid[r-1][c+1]++
	}
	if c-1 >= 0 && r-1 >= 0 && grid[r-1][c-1] != 9 {
		grid[r-1][c-1]++
	}
	// }
}
func printGrid() {
	for row := range grid {
		for c, colValue := range grid[row] {
			if grid[row][c] == 9 {
				redColor.Print(colValue, " ")
			} else if grid[row][c] > 0 && grid[row][c] < 9 {
				greenColor.Print(colValue, " ")
			} else {
				whiteColor.Print(colValue, " ")
			}
		}
		fmt.Println()

	}
}
