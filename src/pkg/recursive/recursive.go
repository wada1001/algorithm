package recursive

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// count 1 in supplied number
func CountOneFor(number int) int {
	ret := 0
	for i := number; i > 0; i = i/10 {
		if i % 10 == 1 {
			ret++
		}
	}
	return ret
}

// count 1 in supplied number
func CountOneRecursive(number int) int {
	if number == 0 {
		return 0
	}

	if number % 10 != 1 {
		return 0 + CountOneRecursive(number / 10)
	}

	return 1 + CountOneRecursive(number / 10)
}

func Gcd(nums []int) int {
	if len(nums) == 2 { 
		return GcdIn(nums[0], nums[1])
	}

	return GcdIn(nums[0], Gcd(nums[1:]))
}

func GcdIn(a, b int) int {
	for i:= a; i > 0; i-- {
		if a % i != 0 || b % i != 0 {
			continue
		}
		return i
	}
	return 0
}


// Back Track(DFS)
func EightQueen(board [8][8]int, x, count int) bool {
	if count == 8 {
		PrintBoard(board)
		return true
	}

	for j := 0; j < 8; j++ {
		if (!CheckQueen(board, x, j)) {
			continue
		}

		board[j][x] = 1
		count++

		// using lot memory
		if EightQueen(board, x + 1, count) {
			return true
		} else {
			board[j][x] = 0
			count--
		}
	}
	
	return false
}

func CheckQueen(board [8][8]int, x, y int) bool {
	x1 := x
	for x1 > 0 {
		x1--
		if board[y][x1] != 0 {
			return false;
		}
	}

	x1 = x
	y1 := y
	for x1 > 0 && y1 < 7 {
		x1--
		y1++
		if board[y1][x1] != 0 {
			return false;
		}
	}

	x1 = x
	y1 = y
	for x1 > 0 && y1 > 0 {
		x1--
		y1--
		if board[y1][x1] != 0 {
			return false;
		}
	}
	return true
}

func PrintBoard(board [8][8]int) {
	for y := 0; y < 8; y++ {
		for x:= 0; x < 8; x++ {
			fmt.Print(board[x][y])
			fmt.Print(",")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

type Puzzle struct {
	x int
	y int
	board []string
	goal string
	hist []string
	queue []string	
}

func MakePuzzle(x, y int) *Puzzle{
	b := make([]string, y * x)
	g  := make([]string, len(b))
	for i := range b{
		if i == y * x -1 {
			b[i] = strconv.Itoa(0)
			continue;
		}
		b[i] = strconv.Itoa(i+1)
	}
	copy(g, b)  

	n := len(b)
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        b[i], b[j] = b[j], b[i]
    }

	p := &Puzzle{x: x, y: y, board: b, hist: []string{}, queue: []string{}}
	p.goal = p.MakeHash(g)
	return p
}

// BFS
func (p *Puzzle) Solve() {
	p.queue = append(p.queue, p.MakeHash(p.board))
	
	for len(p.queue) != 0 {
		// dequeue
		c := p.queue[0]
		p.queue = p.queue[1:]
		arr := strings.Split(c, ",")

		// TODO make history

		if c == p.goal {
			fmt.Println(p.goal)
			return;
		}

		index := -1
		for i := range arr {
			if arr[i] == "0" {
				index = i
			}
		}

		cy := index / 3
		cx := index % p.x

		if cy + 1 < p.y {
			arr[index] = arr[index + p.y]
			arr[index + p.y] = "0"
			p.queue = append(p.queue, p.MakeHash(arr))
			arr[index + p.y] = arr[index]
			arr[index] = "0"
		}

		if cy - 1 >= 0 {
			arr[index] = arr[index - p.y]
			arr[index - p.y] = "0"
			p.queue = append(p.queue, p.MakeHash(arr))
			arr[index - p.y] = arr[index]
			arr[index] = "0"
		}

		if cx + 1 < p.x {
			arr[index] = arr[index + 1]
			arr[index + 1] = "0"
			p.queue = append(p.queue, p.MakeHash(arr))
			arr[index + 1] = arr[index]
			arr[index] = "0"
		}

		if cx - 1 >= 0 {
			arr[index] = arr[index - 1]
			arr[index - 1] = "0"
			p.queue = append(p.queue, p.MakeHash(arr))
			arr[index - 1] = arr[index]
			arr[index] = "0"
		}
	}
}

func (p *Puzzle) MakeHash(board []string) string {
	// for i := 0; i < p.y; i++ {
	// 	for j:= 0; j < p.x; j++ {
	// 		fmt.Print(board[i % p.y * p.y + j])
	// 		fmt.Print(",")
	// 	}
	// 	fmt.Println("")
	// }
	// fmt.Println("")
	hash := ""
	for i := range board {
		if i != 0 {
			hash += "," + board[i];
		} else {
			hash += board[i];
		}
	}
	return hash
}
