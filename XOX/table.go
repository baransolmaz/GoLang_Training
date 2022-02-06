package main

import (
	"fmt"
)

type table [3][3]byte

func newTable() table {
	newT := [3][3]byte{{}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			newT[i][j] = ' '
		}
	}
	return newT
}
func (t table) print() {
	print("  0 1 2\n")
	ch := [3]string{"a ", "b ", "c "}
	for i := 0; i < 3; i++ {
		fmt.Print(ch[i])
		for j := 0; j < 3; j++ {
			fmt.Printf("%c ", t[i][j])
		}
		println()
	}
}
func (t *table) play(loc string, player byte) bool {
	if len(loc) != 2 {
		fmt.Println("Wrong Input")
		return false
	}
	x := byte(loc[0]) - byte('a')
	y := byte(loc[1]) - byte('0')
	if x >= 3 || y >= 3 {
		fmt.Println("Index out of bounds")
		return false
	}
	if t[x][y] != ' ' {
		fmt.Println("Wrong Location")
		return false
	}
	t[x][y] = player
	return true
}
func (t table) checkWin() bool {
	res, r := checkRows(t)
	if res {
		t.print()
		fmt.Printf("%c Wins\n", t[r][0])
		return true
	}
	res, r = checkCols(t)
	if res {
		t.print()
		fmt.Printf("%c Wins\n", t[0][r])
		return true
	}
	res, r = checkDiag(t)
	if res {
		t.print()
		fmt.Printf("%c Wins\n", t[r][0])
		return true
	}
	res = checkDraw(t)
	if res {
		t.print()
		fmt.Printf("Draw\n")
		return true
	}
	return false
}
func checkRows(t table) (bool, int) {
	if (t[0][0] == t[0][1]) && (t[0][0] == t[0][2]) && t[0][0] != ' ' {
		return true, 0
	} else if (t[1][0] == t[1][1]) && (t[1][0] == t[1][2]) && t[1][0] != ' ' {
		return true, 1
	} else if (t[2][0] == t[2][1]) && (t[2][0] == t[2][2]) && t[2][0] != ' ' {
		return true, 2
	}
	return false, -1
}
func checkCols(t table) (bool, int) {
	if (t[0][0] == t[1][0]) && (t[0][0] == t[2][0]) && t[0][0] != ' ' {
		return true, 0
	} else if (t[0][1] == t[1][1]) && (t[0][1] == t[2][1]) && t[0][1] != ' ' {
		return true, 1
	} else if (t[0][2] == t[1][2]) && (t[0][2] == t[2][2]) && t[0][2] != ' ' {
		return true, 2
	}
	return false, -1
}
func checkDiag(t table) (bool, int) {
	if (t[0][0] == t[1][1]) && (t[0][0] == t[2][2]) && t[0][0] != ' ' {
		return true, 0
	} else if (t[2][0] == t[1][1]) && (t[1][1] == t[0][2]) && t[2][0] != ' ' {
		return true, 2
	}
	return false, -1
}
func checkDraw(t table) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if t[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}
