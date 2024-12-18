package main

// Capstone project. Animated Retro Digital Clock

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// A function found online to command the terminal to clear
func clearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Prints a given [8][5]Array, which is how our clock is built
func prntClk(c [8][5]string) {
	fmt.Println("")
	// Height loop
	for j := 0; j < 5; j++ {
		// Width loop
		for i := 0; i < 8; i++ {
			fmt.Print(c[i][j])
			// Print space between numbers/placeholders
			fmt.Print("  ")
		}
		// New line for next index
		fmt.Println("")
	}
	// Print line to bottom buffer
	fmt.Println()

}

func main() {

	// The clock multidimensional array will hold each of the 8 numbers
	// of the current time in a per-second snapshot. Populated initially
	// with solid colon, and switching back and forth
	clock := [8][5]string{
		2: colon,
		5: colon,
	}

	// The boolean to control whether colon is opaque or translucent
	colonSwitch := false

	// Our main infinite for loop
	for {
		// Clear console function
		clearConsole()

		// If opaque colons, switch to translucent
		if colonSwitch {
			clock[2], clock[5] = colon2, colon2
			colonSwitch = false
		} else {
			clock[2], clock[5] = colon, colon
			colonSwitch = true
		}

		// Get current time t
		t := time.Now()

		// Get hour/minute/second
		h := t.Hour()
		m := t.Minute()
		s := t.Second()

		// Hour place
		clock[0] = digits[h/10]
		clock[1] = digits[h%10]

		// Minute place
		clock[3] = digits[m/10]
		clock[4] = digits[m%10]

		// Second place
		clock[6] = digits[s/10]
		clock[7] = digits[s%10]

		prntClk(clock)

		time.Sleep(1 * time.Second)

	}

}
