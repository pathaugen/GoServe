

package main


import (
	//"bufio"
	"fmt"
	"os"
	//"strings"
	"net/http"
	
	//"github.com/fatih/color"
	//"external/color-master"
	//"GoServe/external/color-master"
	//"GoServe/external/go-colorable-master"
	
	"GoServe/external/ansicolor-master"
	
	//"GoServe/handlers"
)


// ========== ========== ========== ========== ==========
// Variables spanning across all .go files and functions in application
var exit		bool	= false
var version		string	= "1.0.0"
var monitoring	bool	= false		// Console based traffic monitoring switch (configured via flag or console)
var port		string	= "8080"	// Port to serve web requests on (configured via flag or console)
// ========== ========== ========== ========== ==========


func main() {
	
	/*
	fmt.Print("\n\n")
	fmt.Print(color.RedString("This is a test of Color"))
	fmt.Print("\n\n")
	*/
	
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
	
	//logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	//logrus.SetOutput(colorable.NewColorableStdout())
	
	// ========== ========== ========== ========== ==========
	// Splash Screen
	// http://patorjk.com/software/taag/#p=display&v=1&f=Big&t=GoServe!
	//fmt.Print(color.RedString("   _____       _____                     _"), "\n")
	//logrus.Info("   _____       _____                     _")
	//fmt.Print("\n")
	//fmt.Print("\x1b[31m", "   _____       _____                     _", "\n")
	//fmt.Print(uint16(0x0004 | 0x0000), "31", "   _____       _____                     _", "\n")
	//fmt.Print("^[[31m", "   _____       _____                     _", "\n")
	fmt.Fprintf(consoleColor, "\x1b[36m\x1b[1m") // Cyan, bold
	fmt.Print("   _____       _____                     _", "\n")
	fmt.Print("  / ____|     / ____|                   | |", "\n")
	fmt.Print(" | |  __  ___| (___   ___ _ ____   _____| |", "\n")
	fmt.Print(" | | |_ |/ _ \\\\___ \\ / _ \\ '__\\ \\ / / _ \\ |", "\n")
	fmt.Print(" | |__| | (_) |___) |  __/ |   \\ V /  __/_|", "\n")
	fmt.Print("  \\_____|\\___/_____/ \\___|_|    \\_/ \\___(_)")
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
	
	
	// ========== ========== ========== ========== ==========
	// Multiple webservers can be created and started
	// > webserver list -> List all running webservers
	// > webserver start 8080 -> Start a webserver on the specified port
	// webservers.cfg -> configuration file with webservers that will be run on startup
	fmt.Print("Starting Webserver on port "+port+".. please wait..")
	fmt.Print("\n")
	//http.HandleFunc("/", handler)
    //http.ListenAndServe(":8080", nil)
    
    // Can have multiple servers on different ports by duplicating this code
    server1 := http.NewServeMux()
    server1.HandleFunc("/", handlerWeb)
    server1.HandleFunc("/api/", handlerApi)
    go func() {
		http.ListenAndServe(":"+port, server1)
    }()
    
    fmt.Fprintf(consoleColor, "\x1b[32m\x1b[1m") // Green, bold
	fmt.Print("SUCCESS: Webserver STARTED on port "+port)
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	fmt.Print("\n")
	fmt.Print("Visit: http://localhost:"+port)
	fmt.Print("\n\n")
	// ========== ========== ========== ========== ==========
	
	
	// Every 10 seconds save log files
	// xxx
	
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

