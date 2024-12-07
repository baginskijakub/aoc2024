package day6

// import "fmt"

// type Guard struct {
// 	row       int
// 	col       int
// 	direction string
// }

// type State struct {
// 	guard   Guard
// 	matrix  [][]string
// 	wallHit bool
// 	visited [][]int
// }

// func Part1() {
// 	matrix, err := ReadFile()

// 	if err != nil {
// 		panic("Couldnt read file")
// 	}

// 	guard := findGuardPosition(matrix)

// 	var visited [][]int

// 	state := State{guard, matrix, false, visited}
// 	state = visit(state, state.guard.row, state.guard.col)

// 	for !state.wallHit {
// 		state = nextStep(state)
// 		fmt.Println(state.guard)
// 	}

// 	fmt.Println(len(state.visited))
// }

// func nextStep(state State) State {
// 	if wasWallHit(state) {
// 		state.wallHit = true
// 		return state
// 	}

// 	canMove := canMoveForward(state)

// 	if canMove {
// 		return moveForward(state)
// 	} else {
// 		return turnRight(state)
// 	}
// }

// func canMoveForward(state State) bool {
// 	if state.guard.direction == "right" {
// 		return state.matrix[state.guard.row][state.guard.col+1] == "."
// 	}

// 	if state.guard.direction == "down" {
// 		return state.matrix[state.guard.row+1][state.guard.col] == "."
// 	}

// 	if state.guard.direction == "left" {
// 		return state.matrix[state.guard.row][state.guard.col-1] == "."
// 	}

// 	if state.guard.direction == "up" {
// 		return state.matrix[state.guard.row-1][state.guard.col] == "."
// 	}

// 	return false
// }

// func wasWallHit(state State) bool {
// 	if state.guard.direction == "right" {
// 		return state.guard.col+1 >= len(state.matrix[0])
// 	}

// 	if state.guard.direction == "down" {
// 		return state.guard.row+1 >= len(state.matrix)
// 	}

// 	if state.guard.direction == "left" {
// 		return state.guard.col-1 < 0
// 	}

// 	if state.guard.direction == "up" {
// 		return state.guard.row-1 < 0
// 	}

// 	return false
// }

// func moveForward(state State) State {

// 	if state.guard.direction == "right" {
// 		state.guard.col++
// 	}

// 	if state.guard.direction == "down" {
// 		state.guard.row++
// 	}

// 	if state.guard.direction == "left" {
// 		state.guard.col--
// 	}

// 	if state.guard.direction == "up" {
// 		state.guard.row--
// 	}

// 	state = visit(state, state.guard.row, state.guard.col)

// 	return state
// }

// var TURN_MAP = map[string]string{
// 	"right": "down",
// 	"down":  "left",
// 	"left":  "up",
// 	"up":    "right",
// }

// func turnRight(state State) State {
// 	print("Turning\n")
// 	state.guard.direction = TURN_MAP[state.guard.direction]

// 	return state
// }

// func findGuardPosition(matrix [][]string) Guard {
// 	for i, row := range matrix {
// 		for j, cell := range row {
// 			if cell == "^" {
// 				matrix[i][j] = "."
// 				return Guard{i, j, "up"}
// 			}
// 		}
// 	}

// 	panic("Guard not found")
// }

// func visit(state State, row int, col int) State {
// 	for _, visit := range state.visited {
// 		if visit[0] == row && visit[1] == col {
// 			return state
// 		}
// 	}

// 	state.visited = append(state.visited, []int{row, col})

// 	return state
// }
