package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bahern/go-sudoku-solver/sudoku"
)

func main() {
	filepath := flag.String("i", "", "Filepath to a properly formatted puzzle file")
	flag.Parse()

	if filepath == nil || *filepath == "" {
		flag.Usage()
		return
	}

	file, err := os.Open(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := sudoku.Puzzle{}
	p.Load(file)

	fmt.Println("Initial puzzle:")
	p.Print()
	fmt.Println()

	solved := p.Solve()
	if solved {
		fmt.Println("Solved puzzle:")
	} else {
		fmt.Println("The puzzle was not solvable.  Here is the current state:")
	}
	p.Print()
}
