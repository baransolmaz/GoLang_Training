package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cardType := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Quenn", "King"}
	newD := deck{}
	for _, typ := range cardType {
		for _, val := range cardValue {
			newD = append(newD, val+" of "+typ)
		}
	}
	return newD
}
func (d deck) print() {
	for i, card := range d {
		println(i+1, card)
	}
}
func deal(d deck, i int) (deck, deck) {
	return d[:i], d[i:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")

}
func (d deck) saveToFile(name string) error {
	return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}
func readFromFile(name string) deck {
	byteVer, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	strVer := string(byteVer)
	return deck(strings.Split(strVer, "\n"))

}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		pos := r.Intn(len(d) - 1)
		d[i], d[pos] = d[pos], d[i]

	}
}
