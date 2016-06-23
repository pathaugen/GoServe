

package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func waitForInput() {
	
	consolereader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter a command ('exit' to exit) > ")
	input, err := consolereader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	input = strings.TrimSpace(input)
	// TODO: Strip '/', '-', '- ' off the start of the string: /help, -help, etc.
	
	if input == "" {
		fmt.Print("\n")
		fmt.Print("Command not found. 'help' for a list of commands.")
		fmt.Print("\n\n")
	} else if input == "exit" {
		exit = true
	} else if input == "help" || input == "?" {
		fmt.Print("\n")
		fmt.Print("GoServe! Help Pages")
		fmt.Print("\n\n")
	} else {
		fmt.Print("\n")
		fmt.Print("Command entered : ")
		fmt.Print(input)
		fmt.Print("\n\n")
	}
}

