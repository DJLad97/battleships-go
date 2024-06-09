package main

import (
	"fmt"
	// "bufio"
	// "os"
	//"strings"
)

func main() {
	printBoard()
	// reader := bufio.NewReader(os.Stdin)
	
    // for {
    //     text, _ := reader.ReadString('\n')
    //     fmt.Println(text)
    //
    // }
    fmt.Println("Hello, World!")
}

func printBoard() {
	var letters = [10]string{"A","B","C","D","E","F","G","H","I","J"}

	fmt.Print("  ")
	for rowNumbers := 0; rowNumbers < 10; rowNumbers++ {
		if rowNumbers != 9 {
			fmt.Print("  ")
		} else {
			fmt.Print(" ")
		}
		fmt.Print(rowNumbers + 1)
		fmt.Print(" ")
	}

	fmt.Println()

	for col := 0; col < 9; col++ {
		fmt.Print(letters[col] + " ")	
		fmt.Print("| · ")
		for row:= 0; row < 9; row++ {
			if row == 8 {
				fmt.Print("| · |")
			} else {
				fmt.Print("| · ")
			}
		}
		fmt.Println("")
	}
}
