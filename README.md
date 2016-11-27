# go-sudoku-solver

A simple Sudoku solver written in Go.

Setting up your dev environment
-----------
### Build & Unit Test
The following will setup your dev environment to build the project and run unit tests.

- Download Go 1.7 from [here](https://golang.org/dl/) and run through the [installation steps](https://golang.org/doc/install)
- Pull down the source code into your $GOPATH/src dir
```bash
git clone git@github.com:bahern/go-sudoku-solver.git $GOPATH/src/github.com/bahern/go-sudoku-solver
```
- Get [Godeps](https://github.com/tools/godep) which handles managing package dependencies in Go
- Verify build and unit tests by running `make test`

Usage
-----------
Basic CLI looks like this `sudoku-solver -i /path/to/some/puzzle_file`

The puzzle file should look like the following:
```
x,3,9,x,x,5,x,8,4
5,8,7,4,x,x,x,x,6
x,x,4,3,x,9,x,x,x
x,x,1,x,x,x,6,x,2
x,5,x,x,7,x,x,9,x
2,x,6,x,x,x,4,x,x
x,x,x,6,x,2,7,x,x
9,x,x,x,x,7,8,2,5
7,4,x,5,x,x,3,6,x
```