package main

import (
	"fmt"
	"math/rand"
	"time"
)

var number = randomNumber(1, 100)

func randomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Second * 1)

	return rand.Intn(max-min) + min
}

func playerID(i int) int {
	if i%2 == 0 {
		return 2
	} else {
		return 1
	}
}

type Game struct {
	Guess   int
	Low     int
	Big     int
	Guessed bool
}

func (g *Game) Menu() (i int) {
	fmt.Println("1 - Two players")
	fmt.Println("2 - Against ai")
	fmt.Println("3 - Exit")

	for i < 1 || 3 < i {
		fmt.Printf("\nOption: ")
		fmt.Scanf("%d", &i)
	}
	return
}

func (g *Game) Start() {
	i := g.Menu()

	switch i {
	case 1:
		g.TwoPlayers()

	case 2:
		g.VsAi()
	}
}

func (g *Game) GetGuess(counter int) {
	for g.Guess <= g.Low || g.Big <= g.Guess {
		fmt.Printf("\nEnter your guess player %d: ", playerID(counter))
		fmt.Scanf("%d", &g.Guess)
	}
}

func (g *Game) Eval() bool {
	if g.Guess < number {
		g.Low = g.Guess
		fmt.Println("Higher than that!")
	} else if number < g.Guess {
		g.Big = g.Guess
		fmt.Println("Lower than that!")
	} else {
		return true
	}
	return false
}

func (g *Game) TwoPlayers() {
	counter := 0
	for !g.Guessed {
		counter++

		g.GetGuess(counter)

		if g.Eval() {
			g.Guessed = true
		}
	}
}

func (g *Game) VsAi() {
	for !g.Guessed {
		g.GetGuess(1)

		if g.Eval() {
			g.Guessed = true
			fmt.Println("Congrats you guessed it")
			break
		}

		time.Sleep(time.Second * 1)

		// AI guess
		g.Guess = randomNumber(g.Low, g.Big)
		fmt.Println("\nThe AI guessed: ", g.Guess)

		if g.Eval() {
			g.Guessed = true
			fmt.Println("Sorry the ai guessed it")
			break
		}
	}

}

func NewGame() *Game {
	return &Game{0, 1, 100, false}
}

func main() {
	game := NewGame()
	game.Start()
}
