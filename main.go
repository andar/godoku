package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andar/godoku/board"
)

func main() {
	fileName := os.Args[1]
	puzzle, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Couldn't read file")
		return
	}
	board, err := board.Parse(puzzle)
	if err != nil {
		fmt.Println("couldn't parse board")
		return
	}
	i, j := board.BestCell()
	lastMove := 0
	for !board.Solved() {
		ps := board.PossibleValues(i, j)
		if lastMove != 0 {
			// backtracked
			ps = trimAt(ps, lastMove)
			lastMove = 0
		}
		if len(ps) == 0 {
			// no valid moves for board state. backtrack!
			i, j, lastMove = board.PopMove()
			continue
		} else {
			success := board.PushMove(i, j, ps[0])
			if !success {
				fmt.Printf("Something went terribly wrong for move (%d, %d, %d)", i, j, ps[0])
				fmt.Print(board.String())
				return
			}
			i, j = board.BestCell()
		}

	}
	fmt.Print(board.String())

}

func trimAt(vs []int, val int) []int {
	for x, v := range vs {
		if v == val {
			return vs[x+1:]
		}
	}
	return vs
}
