

package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"GoServe/external/ansicolor-master"
)


func handlerInput() {
	
	consolereader := bufio.NewReader(os.Stdin)
	
	consoleColor := ansicolor.NewAnsiColorWriter(os.Stdout) // Initialize Windows 10 console coloring
	
	// Command prompt
	if !monitoringRequests {
		//fmt.Print("Enter a command ('exit' to exit) > ")
		fmt.Fprintf(consoleColor, "\x1b[33m\x1b[1m") // Yellow, bold
		fmt.Print("GoServe! [DATETIME] > ")
		fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	}
	
	input, err := consolereader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	input = strings.TrimSpace(input)
	// TODO: Strip '/', '-', '- ' off the start of the string: /help, -help, etc.
	
	if input == "" {
		if !monitoringRequests {
			//fmt.Print("\n")
			fmt.Fprintf(consoleColor, "\x1b[31m\x1b[1m") // Red, bold
			fmt.Print("Command not found.")
			fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
			fmt.Print(" 'help' for a list of commands.")
			fmt.Print("\n\n")
		} else {
			fmt.Print("GoServe! Web Traffic Monitoring STOPPING..")
			fmt.Print("\n")
			
			monitoringRequests = false
			
			fmt.Fprintf(consoleColor, "\x1b[31m\x1b[1m") // Red, bold
			fmt.Print("GoServe! Web Traffic Monitoring STOPPED.")
			fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
			fmt.Print("\n\n")
		}
	} else if input == "exit" {
		exit = true
	} else if input == "help" || input == "?" {
		//fmt.Print("\n")
		fmt.Print("GoServe! Help Pages")
		fmt.Print("\n\n")
	} else if input == "monitor" {
		fmt.Print("GoServe! Web Traffic Monitoring STARTING..")
		fmt.Print("\n")
		
		monitoringRequests = true
		
		fmt.Fprintf(consoleColor, "\x1b[32m\x1b[1m") // Green, bold
		fmt.Print("GoServe! Web Traffic Monitoring STARTED.")
		fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
		fmt.Print(" (Press [ENTER] at any time to stop monitoring)")
		fmt.Print("\n\n")
	} else if input == "active" {
		fmt.Print("GoServe! Active User Monitoring STARTING..")
		fmt.Print("\n")
		
		monitoringActive = true
		
		fmt.Fprintf(consoleColor, "\x1b[32m\x1b[1m") // Green, bold
		fmt.Print("GoServe! Active User Monitoring STARTED.")
		fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
		fmt.Print(" (Press [ENTER] at any time to stop monitoring)")
		fmt.Print("\n\n")
	} else {
		//fmt.Print("\n")
		fmt.Print("Command entered : ")
		fmt.Print(input)
		fmt.Print("\n\n")
	}
}

