package main

//GameBoard is a two-dimensional slice of booleans
type GameBoard []([]bool)

// here we will provide our functions for the Game of Life

//PlayGoL takes an initial game board and a number of generations. It returns a slice of game boards of length numGens+1 corresponding to playing Game of Life numGens generations, starting with the initial board.
func PlayGoL(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1])
	}

	return boards
}

//UpdateBoard takes a GameBoard and returns the board resulting from playing the Game of Life for one generation.
func UpdateBoard(currBoard GameBoard) GameBoard {
	// first, create new board corresponding to the next generation.
	// let's have all cells dead to begin.
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	newBoard := InitializeBoard(numRows, numCols)

	//now, update values of newBoard
	//range through all cells of currBoard and update each one into newBoard.
	for r := range currBoard {
		// r will range over rows of board
		// current row is currBoard[r]
		// range over values in currBoard[r]
		for c := range currBoard[r] {
			//curr value is currBoard[r][c]
			newBoard[r][c] = UpdateCell(currBoard, r, c)
		}
	}

	// return newBoard
	return newBoard
}

//UpdateCell takes a GameBoard and row/col indices r and c, and it returns the state of the board at these row/col indices is in the next generation.
func UpdateCell(board GameBoard, r, c int) bool {
	numNeighbors := CountLiveNbrs(board, r, c)

	// now it's just a matter of consulting the rules.
	if board[r][c] == true { // I'm alive
		if numNeighbors == 2 || numNeighbors == 3 {
			return true // stayin alive
		} else {
			return false // dyin out
		}
	} else { //I'm dead now
		if numNeighbors == 3 {
			//zombie!
			return true
		} else { // RIP
			return false
		}
	}
}

//CountLiveNbrs takes a GameBoard board along with row and column indices r, c and it counts the live neighbors of board[r][c].
//It won't consider cells that are off the board.
func CountLiveNbrs(board GameBoard, r, c int) int {
	count := 0

	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i != r || j != c) && InField(board, i, j) {
				if board[i][j] == true {
					// we found a live nbr
					count++
				}
			}
		}
	}

	return count
}

//InField takes a GameBoard board as well as row and col indices (i, j) and returns true if board[i][j] is in the board and false otherwise.
func InField(board GameBoard, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= CountRows(board) || j >= CountCols(board) {
		return false
	}
	// if we survive to here, then we are on the board
	return true
}

func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}

//InitializeBoard takes a number of rows and columns as inputs and returns a gameboard with appropriate number of rows and colums, where all values = false.
func InitializeBoard(numRows, numCols int) GameBoard {
	// make a 2-D slice (default values = false)
	var board GameBoard
	board = make(GameBoard, numRows)
	// now we need to make the rows too
	for r := range board {
		board[r] = make([]bool, numCols)
	}

	return board
}
