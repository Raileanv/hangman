package main

import (
	"reflect"
	"sort"
	"strings"
)

const (
	TOTAL_ERRORS_ALLOWED = 7
)

type Game struct {
	Word       string
	userGueses []string
}

func NewGame(word string) *Game {
	return &Game{Word: word, userGueses: []string{}}
}

func (g *Game) Play(letter string) {
	if !g.Over() && !contains(g.userGueses, letter) {
		g.userGueses = append(g.userGueses, letter)
	}
}

func (g *Game) Over() bool {
	return g.Won() || g.Lost()
}

func (g *Game) Won() bool {
	letters := unique(g.letters())
	userGueses := g.userGueses

	var guessedWords []string

	for _, letter := range userGueses {
		if contains(letters, letter) {
			guessedWords = append(guessedWords, letter)
		}
	}

	sort.Strings(letters)
	sort.Strings(guessedWords)

	return reflect.DeepEqual(letters, guessedWords)
}

func (g *Game) Lost() bool {
	return g.errorsAllowed() == 0
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (g *Game) LettersToGuess() []string {
	var LettersToGuess []string

	for _, letter := range g.letters() {
		if contains(g.userGueses, letter) {
			LettersToGuess = append(LettersToGuess, letter)
		} else {
			LettersToGuess = append(LettersToGuess, "-")
		}
	}

	return LettersToGuess
}

func (g *Game) ErrorsToString() string {
	return strings.Join(g.Errors(), " ")
}

func (g *Game) letters() []string {
	return strings.Split(g.Word, "")
}

func (g *Game) Errors() []string {
	var diff []string

	letters := g.letters()
	guesses := g.userGueses

	for _, letter := range guesses {
		if !contains(letters, letter) {
			diff = append(diff, letter)
		}
	}

	return diff
}

func (g *Game) errorsAllowed() int {
	return TOTAL_ERRORS_ALLOWED - len(g.Errors())
}

func contains(s []string, searchterm string) bool {
	sort.Strings(s)

	i := sort.SearchStrings(s, searchterm)

	return i < len(s) && s[i] == searchterm
}
