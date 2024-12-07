package day6

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Guard struct {
	row       int
	col       int
	direction string
}

type State struct {
	guard    Guard
	matrix   [][]string
	wallHit  bool
	looped   bool
	visited  []VisitedNode
	obstacle []int
}

type VisitedNode struct {
	row       int
	col       int
	direction string
}

func Part2() {
	matrix, err := ReadFile()

	if err != nil {
		panic("Couldnt read file")
	}

	var sum int32
	sum = 0

	wg := sync.WaitGroup{}

	for i := range matrix {
		for j := range matrix[i] {
			wg.Add(1)

			go func(i, j int, wg *sync.WaitGroup, sum *int32) {
				defer wg.Done()
				fmt.Println(i, j)

				copied := make([][]string, len(matrix))

				for i := range matrix {
					copied[i] = make([]string, len(matrix[i]))
					copy(copied[i], matrix[i])
				}

				guard := findGuardPosition(copied)

				var visited []VisitedNode
				var obstacle []int
				obstacle = append(obstacle, i)
				obstacle = append(obstacle, j)

				state := State{guard, copied, false, false, visited, obstacle}

				if computeObstaclePosition(state) {
					atomic.AddInt32(sum, 1)
				}
			}(i, j, &wg, &sum)
		}
	}

	wg.Wait()

	print(sum)
	return
}

func computeObstaclePosition(state State) bool {
	if state.matrix[state.obstacle[0]][state.obstacle[1]] == "." {
		state.matrix[state.obstacle[0]][state.obstacle[1]] = "#"

		return checkLoop(state)
	}

	return false
}

func checkLoop(state State) bool {

	for !state.wallHit && !state.looped {
		state = nextStep(state)
	}

	return state.looped
}

func nextStep(state State) State {
	if wasWallHit(state) {
		state.wallHit = true
		return state
	}

	canMove := canMoveForward(state)

	if canMove {
		return moveForward(state)
	} else {
		return turnRight(state)
	}
}

func canMoveForward(state State) bool {
	if state.guard.direction == "right" {
		return state.matrix[state.guard.row][state.guard.col+1] == "."
	}

	if state.guard.direction == "down" {
		return state.matrix[state.guard.row+1][state.guard.col] == "."
	}

	if state.guard.direction == "left" {
		return state.matrix[state.guard.row][state.guard.col-1] == "."
	}

	if state.guard.direction == "up" {
		return state.matrix[state.guard.row-1][state.guard.col] == "."
	}

	return false
}

func wasWallHit(state State) bool {
	if state.guard.direction == "right" {
		return state.guard.col+1 >= len(state.matrix[0])
	}

	if state.guard.direction == "down" {
		return state.guard.row+1 >= len(state.matrix)
	}

	if state.guard.direction == "left" {
		return state.guard.col-1 < 0
	}

	if state.guard.direction == "up" {
		return state.guard.row-1 < 0
	}

	return false
}

func moveForward(state State) State {

	if state.guard.direction == "right" {
		state.guard.col++
	}

	if state.guard.direction == "down" {
		state.guard.row++
	}

	if state.guard.direction == "left" {
		state.guard.col--
	}

	if state.guard.direction == "up" {
		state.guard.row--
	}

	state = visit(state, state.guard.row, state.guard.col)

	return state
}

func visit(state State, row int, col int) State {
	for _, visit := range state.visited {
		if visit.row == row && visit.col == col && visit.direction == state.guard.direction {
			state.looped = true
			return state
		}
	}

	state.visited = append(state.visited, VisitedNode{row, col, state.guard.direction})

	return state
}

var TURN_MAP = map[string]string{
	"right": "down",
	"down":  "left",
	"left":  "up",
	"up":    "right",
}

func turnRight(state State) State {
	state.guard.direction = TURN_MAP[state.guard.direction]

	return state
}

func findGuardPosition(matrix [][]string) Guard {
	for i, row := range matrix {
		for j, cell := range row {
			if cell == "^" {
				matrix[i][j] = "."
				return Guard{i, j, "up"}
			}
		}
	}

	panic("Guard not found")
}
