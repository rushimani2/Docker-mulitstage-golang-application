package main

import (
	"fmt"
	"math/rand"
	"time"
	"os/exec"
)

const (
	width    = 40
	height   = 10
	numStars = 20
)

func clearScreen() {
	cmd := exec.Command("clear") // For Linux/Unix
	if err := cmd.Run(); err != nil {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
		cmd.Run()
	}
}

func generateStars() [numStars][2]int {
	stars := [numStars][2]int{}
	for i := 0; i < numStars; i++ {
		stars[i][0] = rand.Intn(width)
		stars[i][1] = rand.Intn(height)
	}
	return stars
}

func printStars(stars [numStars][2]int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			found := false
			for _, star := range stars {
				if star[0] == x && star[1] == y {
					fmt.Print("*")
					found = true
					break
				}
			}
			if !found {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		clearScreen()
		stars := generateStars()
		printStars(stars)
		time.Sleep(200 * time.Millisecond) // Adjust speed here
	}
}

