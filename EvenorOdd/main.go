package main

func main() {
	arr := []int{1, 2, 3, 5, 4, 6, 7, 8, 9, 10}

	for _, v := range arr {
		if v%2 == 0 {
			println(v, " is even")
		} else {
			println(v, " is odd")
		}
	}
}
