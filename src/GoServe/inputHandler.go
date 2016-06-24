

package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"GoServe/external/ansicolor-master"
)


func inputHandler() {
	
	consolereader := bufio.NewReader(os.Stdin)
	
	consoleColor := ansicolor.NewAnsiColorWriter(os.Stdout) // Initialize Windows 10 console coloring
	
	// Command prompt
	//fmt.Print("Enter a command ('exit' to exit) > ")
	fmt.Fprintf(consoleColor, "\x1b[33m\x1b[1m") // Yellow, bold
	fmt.Print("GoServe! [DATETIME] > ")
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	
	input, err := consolereader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	input = strings.TrimSpace(input)
	// TODO: Strip '/', '-', '- ' off the start of the string: /help, -help, etc.
	
	if input == "" {
		//fmt.Print("\n")
		fmt.Print("Command not found. 'help' for a list of commands.")
		fmt.Print("\n\n")
	} else if input == "exit" {
		exit = true
	} else if input == "help" || input == "?" {
		//fmt.Print("\n")
		fmt.Print("GoServe! Help Pages")
		fmt.Print("\n\n")
	} else if input == "monitor" {
		fmt.Print("GoServe! Web Traffic Monitoring STARTING..")
		fmt.Print("\n")
		fmt.Print("(Press [ENTER] at any time to stop monitoring)")
		fmt.Print("\n\n")
	} else {
		//fmt.Print("\n")
		fmt.Print("Command entered : ")
		fmt.Print(input)
		fmt.Print("\n\n")
	}
}

