

package main


import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
	"net/http"
)


var exit bool = false
var version string = "1.0.0"



func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n\n")
	fmt.Print("function handler triggered")
	fmt.Print("\n\n")
	
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {
	
	// http://patorjk.com/software/taag/#p=display&v=1&f=Big&t=GoServe!
	fmt.Print("   _____       _____                     _", "\n")
	fmt.Print("  / ____|     / ____|                   | |", "\n")
	fmt.Print(" | |  __  ___| (___   ___ _ ____   _____| |", "\n")
	fmt.Print(" | | |_ |/ _ \\\\___ \\ / _ \\ '__\\ \\ / / _ \\ |", "\n")
	fmt.Print(" | |__| | (_) |___) |  __/ |   \\ V /  __/_|", "\n")
	fmt.Print("  \\_____|\\___/_____/ \\___|_|    \\_/ \\___(_)")
	fmt.Print("\n\n")
	fmt.Print("Version : ", version)
	fmt.Print("\n")
	fmt.Print("https://github.com/pathaugen/GoServe")
	fmt.Print("\n\n")
	
	
	
	fmt.Print("Starting Webserver on port 8080")
	//http.HandleFunc("/", handler)
    //http.ListenAndServe(":8080", nil)
    server8080 := http.NewServeMux()
    server8080.HandleFunc("/", handler)
    go func() {
		http.ListenAndServe(":8080", server8080)
    }()
	fmt.Print("\n\n")
	fmt.Print("Webserver STARTED on port 8080")
	fmt.Print("\n\n")
	
	
	
	// Every 10 seconds save log files
	// xxx
	
	
	for !exit {
		waitForInput()
	}
	
	
	if exit {
		// Save command log file: goserve-commands.log
		// xxx
		
		// Save visitor log file: goserve-visitors.log
		// xxx
		
		exitHandler()
	}
	
	
    
	
}

