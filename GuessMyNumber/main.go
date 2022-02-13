package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Welcome!")
	fmt.Println("Which type do you wanna play?")
	fmt.Println("1- User guess")
	fmt.Println("2- Computer guess")
	var ty string
	fmt.Scanln(&ty)
	checkType(ty)
}
func checkType(ty string) {
	num, err := strconv.Atoi(ty)
	if err != nil {
		fmt.Println("Wrong Input")
		return
	}
	if num == 1 {
		userGuess()
	} else if num == 2 {
		computerGuess()
	} else {
		fmt.Println("Wrong Input")
		return
	}

}
func userGuess() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	myNumber := random.Intn(100)
	fmt.Println("My number is between 0 - 100")
	fmt.Println("Guess my number?")
	getInputs(myNumber)
}
func getInputs(myNumber int) {
	var counter int
	var guess int
	for !checkNumber(myNumber, guess) {
		fmt.Scanln(&guess)
		counter++
	}
	fmt.Println("Try: ", counter)
}
func checkNumber(myNumber int, guess int) bool {
	if myNumber == guess {
		fmt.Println("Thats Right!")
		return true
	} else if myNumber > guess {
		fmt.Println("Low!")
		return false
	} else {
		fmt.Println("High!")
		return false
	}
}
func computerGuess() {
	fmt.Println("Enter H/h for High, L/l for Low.")
	fmt.Println("Enter C/c if guess is right")
	computerNumber := []int{0, 50, 100}
	var letter string
	var counter int
	for !compare(computerNumber, letter) {
		fmt.Println("Computer Guess: ", computerNumber[1])
		fmt.Scanln(&letter)
		counter++
	}
	fmt.Println("Try: ", counter)
}
func compare(number []int, letter string) bool {
	if letter == "C" || letter == "c" {
		return true
	} else if letter == "L" || letter == "l" {
		number[0] = number[1]
	} else if letter == "H" || letter == "h" {
		number[2] = number[1]
	}
	number[1] = (number[0] + number[2]) / 2
	return false
}
