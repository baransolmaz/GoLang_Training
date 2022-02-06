package main

import "fmt"

func main() {
	table := newTable()
	chars := []byte{'x', 'o'}
	for i := 0; !(table.checkWin()); i++ {
		table.print()
		var loc string
		fmt.Printf("%d. Players turn: ( %c )\n", (i%2)+1, chars[i%2])
		fmt.Scanln(&loc)
		for !table.play(loc, chars[i%2]) {
			fmt.Printf("%d. Players turn: ( %c )\n", (i%2)+1, chars[i%2])
			fmt.Scanln(&loc)
		}
	}
}
