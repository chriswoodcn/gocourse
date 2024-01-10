package algorithm

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"
)

func readMiGong(filePath string) ([6][5]int, error) {
	file, err := os.Open(filePath)
	res := [6][5]int{{0}}
	if err != nil {
		return res, errors.New("open file error")
	}
	reader := bufio.NewReader(file)
	row := 0
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		col := 0
		for _, b := range s {
			if b == ' ' || b == '\n' {
				continue
			}
			res[row][col] = int(b - '0')
			fmt.Printf("%d  ", int(b-'0'))
			col++
		}
		fmt.Printf("\n  ")
		row++
	}
	return res, nil
}
func readMiGong2(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("open file error")
	}
	var row, col int
	fmt.Fscanf(file, "%d %d ", &row, &col)
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, col)
		for j := range res[i] {
			fmt.Fscanf(file, "%d", &res[i][j])
		}
	}

	return res, nil
}

type Pointer struct {
	i, j int
}

func (p Pointer) add(a Pointer) Pointer {
	return Pointer{p.i + a.i, p.j + a.j}
}
func (p Pointer) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

var dirs = [4]Pointer{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 迷宫算法
// mg迷宫 start起点 end终点
func walk(mg [][]int, start, end Pointer) [][]int {
	steps := make([][]int, len(mg))
	for i := range steps {
		steps[i] = make([]int, len(mg[i]))
	}

	Q := []Pointer{start}
	for len(Q) > 0 {
		cur := Q[0]
		if cur == end {
			break
		}
		Q = Q[1:]
		// 探索四个方向
		for _, dir := range dirs {

			next := cur.add(dir)

			val, ok := next.at(mg) //是否在迷宫内
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps) //是否已经走过 !=0表示已经走过了
			if !ok || val != 0 {
				continue
			}
			if next == start { // 是否到了终点
				continue
			}

			// step记录当前步数
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1
			// 探索到的下一个点放入队列中
			Q = append(Q, next)
		}
	}
	return steps
}

func TestLabyrinth(t *testing.T) {
	// 算法逻辑
	// 从第0步开始往外探索 1 2 ...
	//           2
	//       2  1   2
	//  2  1   1   1  2
	//      2   1  2
	//           2
	mg, _ := readMiGong2("./MiGong.in")
	fmt.Println(mg)
	steps := walk(mg, Pointer{0, 0}, Pointer{len(mg) - 1, len(mg[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
