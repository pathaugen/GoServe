

package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"GoServe/external/ansicolor-master"
	
	// Chromium GUI
	"GoServe/external/go-thrust-master/lib/dispatcher"
	"GoServe/external/go-thrust-master/lib/spawn"
	"GoServe/external/go-thrust-master/lib/bindings/window"
	"GoServe/external/go-thrust-master/lib/commands"
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
		} else { // TODO: Stop "active" and "monitor" status separatly
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
		
	} else if input == "listen" { // TODO: Add port number
		// Start listening on a port
		listenPort := "8080"
		startWebServer(listenPort)
		
		// Start Chrome based GUI
		spawn.Run()
		//thrustWindow	:= window.NewWindow("http://breach.cc/", nil)
		thrustUrl		:= "http://localhost:8080/"
		thrustSize		:= commands.SizeHW{Width:1024, Height: 768}
		thrustTitle		:= "GoServe!"
		thrustOptions	:= window.Options{
			RootUrl:	thrustUrl,
			Size:		thrustSize,
			Title:		thrustTitle,
			//IconPath:	""
			HasFrame:	true} // Size:"SizeHW{Width:1024,Height:768}"
		thrustWindow	:= window.NewWindow(thrustOptions) // Have to send options
		thrustWindow.Show()
		//thrustWindow.Maximize()
		thrustWindow.Focus()
	    go func() {
			dispatcher.RunLoop()
	    }()
		
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

