# Minesweeper Game

The provided code implements a simple version of the classic Minesweeper game using Go programming language. The game is played in the terminal/console environment. Players reveal cells on a grid, trying to avoid hidden mines while uncovering safe areas.

## Features

1. **Grid Generation:** The game initializes a 15x15 grid where players will play the game.

2. **Random Bomb Placement:** Bombs are randomly placed on the grid, making each game session unique.

3. **Adjacent Bomb Count:** Each non-bomb cell stores information about how many neighboring cells contain bombs. This count aids players in identifying safe cells.

4. **User Input and Interaction:** Players input coordinates to reveal cells. If a cell contains a bomb, the game ends. Otherwise, adjacent cells without bombs are revealed recursively.

5. **Clearing the Screen:** The screen is cleared after each user input to provide a cleaner interface.

6. **Visual Indicators:** Cells are color-coded for better visibility. Bombs are marked in red, revealed safe cells show adjacent bomb counts in blue, and unrevealed cells are displayed in white.

## How to Play

1. Run the program.
2. Input row and column coordinates to reveal cells.
3. Avoid revealing cells with hidden bombs.
4. Use the adjacent bomb counts to deduce safe cells.
5. Continue until all safe cells are revealed or a bomb is encountered.

Enjoy playing Minesweeper in your terminal!
