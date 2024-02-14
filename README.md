# Minesweeper Game

This project implements a simple version of the classic Minesweeper game using Go programming language. The game is played in the terminal/console environment. Players reveal cells on a grid, trying to avoid hidden mines while uncovering safe areas.

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

![image](https://github.com/icryez/minesweeper/assets/35337801/f69cb822-04cc-43c7-8c1a-3a48f728eee9)

![image](https://github.com/icryez/minesweeper/assets/35337801/15a0b6d2-a5bf-4cdf-bd55-2bfdb8035711)

![image](https://github.com/icryez/minesweeper/assets/35337801/65581be1-b4ee-4e80-a6b2-2e7c8a394d8e)


