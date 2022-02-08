package main

import "fmt"

func main() {
	colors := map[string]int{
		"red":   1,
		"green": 2,
		"blue":  3,
	}
	fmt.Println(colors)
	colors["white"] = 4
	fmt.Println(colors)
	for key, val := range colors {
		fmt.Println(key, val)
	}
	delete(colors, "blue")
	fmt.Println(colors)
}
