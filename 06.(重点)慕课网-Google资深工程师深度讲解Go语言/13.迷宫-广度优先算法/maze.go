package main

import (
	"fmt"
	"os"
)

// 返回的是一个二维数组
func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var row, col int
	// 获取文件行和列的数据,Fscanf为取内容
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 二维数组行的高度
	maze := make([][]int, row)
	// 二维数组宽度
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

// 四个方向
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 判断当前的步子是否存在不合法的情况，合法就返回可以，然后进入队列
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

//走路的操作
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	queue := []point{start}
	// 判断队列长度
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		// 如果探寻的点为终点，我们就不进行走路了，直接进行返回
		if current == end {
			break
		}
		for _, dir := range dirs {
			next := current.add(dir)
			// 需要判断下一个节点的有效值为0才能进行探索，不然就是一堵墙，是不能走的
			// 往回走时next != start
			// 当前步和下一步都需要是0

			// 判断当前节点是否可以是走下去
			if value, ok := next.at(maze); !ok || value == 1 {
				// 如果不Ok，说明可能撞墙了
				// 我们继续下一个点
				continue
			}
			// 判断当前节点是否可以是走下去
			if value, ok := next.at(steps); !ok || value != 0 {
				// 如果不Ok，说明可能撞墙了
				// 我们继续下一个点
				continue
			}
			if next == start {
				continue
			}

			// 上面的排除了不相关的节点，我们就可以进行逻辑保存于操作 -- 探索
			// 获取当前的步骤数
			currentSteps, _ := current.at(steps)
			steps[next.i][next.j] = currentSteps + 1
			queue = append(queue, next)
		}
	}
	return steps
}

// 根据传入的点数，计算出新的坐标单位
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func main() {
	maze := readMaze("13.迷宫-广度优先算法\\migong.ini")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, value := range row {
			// 这个位置获取的就是我们二维数组保存的值
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
