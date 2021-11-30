package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Привет!")
	file, err := ioutil.ReadFile("./words")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	words := strings.Split(string(file), "\n")

	rand.Seed(time.Now().Unix())
	randIndex := rand.Intn(len(words))

	word := words[randIndex]

	game := NewGame(word)
	consoleInterface := NewConsoleInterface(game)

	for !game.Over() {
		consoleInterface.PrintOut()

		letter := consoleInterface.GetLetter()

		game.Play(letter)
	}

	consoleInterface.PrintOut()
}
