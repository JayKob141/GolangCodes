package main

import ( 
    "math/rand" 
    "fmt"
    "time"
)

type Mine struct {
    minesAround int
    isMine bool 
    clicked bool
}

func printBoard(nRows,nCols int, board [][]Mine ) {
    var caracter = ""

    /* Print numbers of columns */
    fmt.Print("  ")
    for k := 0; k < nCols; k++ {
        fmt.Printf("%d ", k + 1)
    }
    fmt.Println()

    for i := 0; i < nRows; i++ {

        /* Printing number of row */
        fmt.Printf("%d|", i + 1)

        for j := 0; j < nCols; j++ {
            mine := board[i][j]

            if mine.clicked {
                if board[i][j].isMine {
                    fmt.Printf("X ")
                } else {
                    if board[i][j].minesAround == 0 {
                        fmt.Printf("_ ")
                    } else {
                        fmt.Printf("%d ", board[i][j].minesAround)
                    }
                }
            } else {
                if false {
                    // cheat mode
                    if board[i][j].isMine {
                        fmt.Printf("X ")
                    } else {
                        fmt.Printf("%d ", board[i][j].minesAround)
                    } 
                } else {
                    fmt.Printf("%s ", caracter)
                }
            }

        }
        fmt.Println()
    }
}

func generateTwoRandomNumbers(nI, nJ int) (int, int) {
    i := rand.Intn(nI)	
    j := rand.Intn(nJ)	
    return i, j
}

func createBoard(nRows,nCols int) [][]Mine {
    board := make([][]Mine, nRows)

    /* Initialize array */
    for i := 0 ; i < nRows ; i++ {
        for j := 0 ; j < nCols ; j++ {
            arrayOfMines := make([]Mine, nCols)
            board[i] = arrayOfMines
        }
    }

    return board
}

func recursivelyUpdate(i, j, nRows, nCols int, board [][]Mine) {
    if i < 0 || i >= nRows || j < 0 || j >= nCols  || board[i][j].clicked {
        return
    }

    board[i][j].clicked = true

    if board[i][j].minesAround > 0 {
        return
    }

    recursivelyUpdate(i-1, j,   nRows, nCols, board)
    recursivelyUpdate(i,   j-1, nRows, nCols, board)
    recursivelyUpdate(i-1, j-1, nRows, nCols, board)

    recursivelyUpdate(i+1, j,   nRows, nCols, board)
    recursivelyUpdate(i,   j+1, nRows, nCols, board)
    recursivelyUpdate(i+1, j+1, nRows, nCols, board)

    recursivelyUpdate(i+1, j-1, nRows, nCols, board)
    recursivelyUpdate(i-1, j+1, nRows, nCols, board)
}

func simulateClick(i, j, nRows, nCols int, board [][]Mine) int {

    if i < 0 || i >= nRows || j < 0 || j >= nCols  {
        fmt.Println("Invalid position")
        return 0
    }
    if board[i][j].clicked {
        fmt.Printf("Position at (%d, %d) Already clicked\n", i, j)
        return 0
    }

    if board[i][j].isMine {
        fmt.Println("You lose. The game should end here!")
        return 1
    }

    recursivelyUpdate(i, j, nRows, nCols, board)

    if !anyCellAvailableForClickFound(nRows,nCols, board) {
        fmt.Println("You win!")
        return 2
    }

    return 0
}

func anyCellAvailableForClickFound(nRows, nCols int, board [][]Mine) bool {

    for i := 0 ; i < nRows ; i++ {
        for j := 0 ; j < nCols ; j++ {
            if !board[i][j].clicked && !board[i][j].isMine {
                return true
            }
        }
    }

    return false
}


func validIndexes(i,j,nRows,nCols int) bool{
    return !(i < 0 || i >= nRows || j < 0 || j >= nCols)
}


func resetBoard(nRows,nCols int, board [][]Mine) {
    numberOfMines := 10

    /* Reset initial values */
    for i := 0 ; i < nRows ; i++ {
        for j := 0 ; j < nCols ; j++ {
            board[i][j] = Mine{ minesAround: 0, clicked: false, isMine: false }
        }
    }

    /* Assign random mines */
    k := 0
    var selectedMineIsMine bool = false
    var i, j int
    for k < numberOfMines {
        i, j = generateTwoRandomNumbers(nRows, nCols)
        selectedMineIsMine = board[i][j].isMine

        for selectedMineIsMine {
            i, j = generateTwoRandomNumbers(nRows, nCols)
            selectedMineIsMine = board[i][j].isMine
        }

        board[i][j].isMine = true
        k += 1
    }

    /*  Compute mines around */
    for i := 0 ; i < nRows; i++ {
        for j := 0 ; j < nCols; j++ {

            x, y := (i - 1), j
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }

            x, y = i, j - 1
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }

            x, y = i - 1, j + 1
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }

            x, y = i + 1, j  -1
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }


            x, y = i - 1, j - 1
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }


            x, y = i + 1, j + 1
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }


            x, y = i + 1, j
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }


            x, y = i, j + 1
            if validIndexes(x,y,nRows,nCols) && board[x][y].isMine {
                board[i][j].minesAround += 1
            }
        }
    }
}

// ###########################
//
// 		MAIN FUNCTION
//
// ###########################

func main() {
    nRows, nCols := 10, 10

    board := createBoard(nRows, nCols)

    rand.Seed(time.Now().UnixNano())

    resetBoard(nRows, nCols, board)

    var i, j int
    gameOver := false

    for !gameOver {

        fmt.Println("###########################")
        printBoard(nRows, nCols, board)
        fmt.Println("###########################")
        fmt.Print("Type row and column: ")
        fmt.Scanf("%d %d", &i, &j)
        i = i - 1
        j = j - 1
        returnCode := simulateClick(i, j, nRows, nCols, board)

        if returnCode != 0 {
            gameOver = true
        }
    }


}