package main

//GameBoard is a two-dimensional slice of booleans
type GameBoard []([]bool)

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

func DrawGoLBoard(board GameBoard, cellWidth int) Canvas {
	//make the board, which we will assume is rectangular
	numRows := len(board)
	numCols := len(board[0])
	w := numCols * cellWidth
	h := numRows * cellWidth
	pic := CreateNewCanvas(w, h)
	// set the fill color as black and fill the board with it
	black := MakeColor(0, 0, 0)
	pic.SetFillColor(black)
	pic.Clear()
	// set white as the stroke color
	white := MakeColor(255, 255, 255)
	pic.SetStrokeColor(white)
	//draw grid lines
	DrawGridLines(pic, cellWidth)
	// set the fill color to white, then range over the board, filling in each alive square
	pic.SetFillColor(white)
	for row := range board {
		for col := range board[row] {
			if board[row][col] {
				DrawSquare(pic, row, col, cellWidth)
			}
		}
	}
	pic.Fill()
	return pic
}

func DrawGridLines(pic Canvas, cellWidth int) {
	w, h := pic.width, pic.height
	// first, draw vertical lines
	for i := 1; i < pic.width/cellWidth; i++ {
		x := i * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	// next, draw horizontal lines
	for j := 1; j < pic.height/cellWidth; j++ {
		y := j * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	pic.Stroke()
}

func DrawSquare(pic Canvas, row, col, cellWidth int) {
	x1 := col * cellWidth
	y1 := row * cellWidth
	x2 := (col + 1) * cellWidth
	y2 := (row + 1) * cellWidth
	pic.ClearRect(x1, y1, x2, y2)
}

func main() {
	rPentomino := InitializeBoard(10, 10) // sets board to all false
	rPentomino[4][5] = true
	rPentomino[4][6] = true
	rPentomino[5][4] = true
	rPentomino[5][5] = true
	rPentomino[6][5] = true
	cellWidth := 25
	pic := DrawGoLBoard(rPentomino, cellWidth)
	pic.SaveToPNG("GameOfLife.png")
}
