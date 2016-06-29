

package main


import (
	//"bufio"
	"fmt"
	"os"
	//"strings"
	//"net/http"
	
	//"github.com/fatih/color"
	//"external/color-master"
	//"GoServe/external/color-master"
	//"GoServe/external/go-colorable-master"
	
	"GoServe/external/ansicolor-master"
	
	//"GoServe/handlers"
)


// ========== ========== ========== ========== ==========
// Variables spanning across all .go files and functions in application
var exit				bool	= false
var version				string	= "1.0.0"
var monitoringRequests	bool	= false		// Console based traffic monitoring switch (configured via flag or console)
var monitoringActive	bool	= false		// Console based active user monitoring switch (configured via flag or console)
//var port				string	= "8080"	// Port to serve web requests on (configured via flag or console)
var resourcesLocation	string	= ""		// e.g. goserve-resources/port-8080/localhost/
var adminPasswordHash	string	= ""		// MD5 Hashed admin password pulled from .cfg on load or prompts at runtime		
// ========== ========== ========== ========== ==========


func main() {
	
	consoleColor := ansicolor.NewAnsiColorWriter(os.Stdout) // Initialize Windows 10 console coloring
	/*
	text := "%sforeground %sbold%s %sbackground%s\n"
	fmt.Fprintf(w, text, "\x1b[31m", "\x1b[1m", "\x1b[21m", "\x1b[41;32m", "\x1b[0m") // Red
	fmt.Fprintf(w, text, "\x1b[32m", "\x1b[1m", "\x1b[21m", "\x1b[42;31m", "\x1b[0m") // Green
	fmt.Fprintf(w, text, "\x1b[33m", "\x1b[1m", "\x1b[21m", "\x1b[43;34m", "\x1b[0m") // Yellow
	fmt.Fprintf(w, text, "\x1b[34m", "\x1b[1m", "\x1b[21m", "\x1b[44;33m", "\x1b[0m") // Blue
	fmt.Fprintf(w, text, "\x1b[35m", "\x1b[1m", "\x1b[21m", "\x1b[45;36m", "\x1b[0m") // Purple
	fmt.Fprintf(w, text, "\x1b[36m", "\x1b[1m", "\x1b[21m", "\x1b[46;35m", "\x1b[0m") // Cyan
	fmt.Fprintf(w, text, "\x1b[37m", "\x1b[1m", "\x1b[21m", "\x1b[47;30m", "\x1b[0m") // White
	*/
	
	// ========== ========== ========== ========== ==========
	// Splash Screen
	// http://patorjk.com/software/taag/#p=display&v=1&f=Big&t=GoServe!
	fmt.Fprintf(consoleColor, "\x1b[36m\x1b[1m") // Cyan, bold
	asciiLogo := `
   _____       _____                     _
  / ____|     / ____|                   | |
 | |  __  ___| (___   ___ _ ____   _____| |
 | | |_ |/ _ \\___ \ / _ \ '__\ \ / / _ \ |
 | |__| | (_) |___) |  __/ |   \ V /  __/_|
  \_____|\___/_____/ \___|_|    \_/ \___(_)`
	fmt.Print(asciiLogo)
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	fmt.Print("\n\n")
	fmt.Print("Version : ", version)
	fmt.Print("\n")
	fmt.Print("https://github.com/pathaugen/GoServe")
	fmt.Print("\n\n")
	// ========== ========== ========== ========== ==========
	
	
	// ========== ========== ========== ========== ==========
	// Command List
	// TODO: Convert to JSON or another static command/description and format accoringly
	fmt.Print("GoServe! Console Commands:")
	fmt.Print("\n")
	fmt.Fprintf(consoleColor, "\x1b[36m\x1b[1m") // Cyan, bold
	fmt.Print("'monitor'")
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	fmt.Print("	Start web monitoring of visitors.")
	fmt.Print("\n")
	fmt.Fprintf(consoleColor, "\x1b[36m\x1b[1m") // Cyan, bold
	fmt.Print("'help'")
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	fmt.Print("		Open help documentation.")
	fmt.Print("\n")
	fmt.Fprintf(consoleColor, "\x1b[36m\x1b[1m") // Cyan, bold
	fmt.Print("'exit'")
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	fmt.Print("		Exit the application.")
	fmt.Print("\n\n")
	// ========== ========== ========== ========== ==========
	
	
	// Read .cfg file and parse, resume from where left off
	// xxx
	
	
	// Every 10 seconds save log files
	// goserve-activity.log -> load file into log buffer, dump log buffer
	// goserve-commands.log
	
	
	// Loop listening for input via inputHandler
	for !exit {
		handlerInput()
	}
	
	// If exit is set to true, gracefully exit via this code
	if exit {
		// Save command log file: goserve-commands.log
		// xxx
		
		// Save visitor log file: goserve-visitors.log
		// xxx
		
		handlerExit()
	}
}

