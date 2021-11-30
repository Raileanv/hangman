package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

type ConsoleInterface struct {
	Game    *Game
	Figures []string
}

func NewConsoleInterface(game *Game) *ConsoleInterface {
	return &ConsoleInterface{Game: game, Figures: getFigures()}
}

func getFigures() []string {
	files, err := ioutil.ReadDir("./figures")

	var figures []string

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		file, err := ioutil.ReadFile("./figures/" + file.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		figures = append(figures, string(file))
	}

	return figures
}

func (c *ConsoleInterface) PrintOut() {
	color.Yellow("Слово: %s", c.WordToShow())
	color.Red(c.figure())
	color.Magenta("Ошибки (%v): %v", len(c.Game.Errors()), c.Game.ErrorsToString())
	color.Magenta("У вас осталось ошибок: %v", c.Game.errorsAllowed())

	if c.Game.Won() {
		color.Green("Поздравляем, вы выиграли!")
		return
	} else if c.Game.Lost() {
		color.Red("Вы проиграли, загаданное слово: %v\n", c.Game.Word)
		return
	}
}

func (c *ConsoleInterface) GetLetter() string {
	fmt.Println("Введите букву:")
	letter := ""
	fmt.Scanln(&letter)
	return letter
}

func (c *ConsoleInterface) WordToShow() string {
	var wordToShow []string
	for _, letter := range c.Game.LettersToGuess() {
		if letter == "-" {
			wordToShow = append(wordToShow, "*")
		} else {
			wordToShow = append(wordToShow, letter)
		}
	}

	return strings.Join(wordToShow, " ")
}

func (c *ConsoleInterface) figure() string {
	return c.Figures[len(c.Game.Errors())]
}
