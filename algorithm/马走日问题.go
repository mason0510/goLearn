package main

//给定任意一个数，要求走n-1步，在123456789*0#的拨号键盘上,找出所有的路径。 举个例子。 只有从1开始走2步 ,走一步，路径就是1,6 1,8 ;再走一步就是1,6,7  1,6,0 1,8,3。

import (
	"fmt"
)

var (
	result [][]int
	// keyboard is the dial pad
	keyboard = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {-1, 0, -1}}
)

func findPaths(currX, currY, maxN int, path []int) bool {
	if len(path) == maxN {
		// copy the path to the result
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)
		return false
	}
	for _, nextMove := range getNextMoves(currX, currY) {
		path = append(path, keyboard[nextMove[0]][nextMove[1]])
		if findPaths(nextMove[0], nextMove[1], maxN, path) {
			return true
		}
		path = path[:len(path)-1]
	}
	return false
}

func getNextMoves(currX, currY int) [][]int {
	var nextMoves [][]int
	// check up
	if currX > 0 && keyboard[currX-1][currY] != -1 {
		nextMoves = append(nextMoves, []int{currX - 1, currY})
	}
	// check down
	if currX < len(keyboard)-1 && keyboard[currX+1][currY] != -1 {
		nextMoves = append(nextMoves, []int{currX + 1, currY})
	}
	// check left
	if currY > 0 && keyboard[currX][currY-1] != -1 {
		nextMoves = append(nextMoves, []int{currX, currY - 1})
	}
	// check right
	if currY < len(keyboard[0])-1 && keyboard[currX][currY+1] != -1 {
		nextMoves = append(nextMoves, []int{currX, currY + 1})
	}
	return nextMoves
}

func main() {
	//定义走的步数
	var stepMax = 5
	result = make([][]int, 0)
	path := []int{keyboard[0][0]}
	findPaths(0, 0, stepMax, path)
	fmt.Println(result)
}
