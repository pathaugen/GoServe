

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
	
	/*
	fmt.Print("\n\n")
	fmt.Print(color.RedString("This is a test of Color"))
	fmt.Print("\n\n")
	*/
	
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	text := "%sforeground %sbold%s %sbackground%s\n"
	fmt.Fprintf(w, text, "\x1b[31m", "\x1b[1m", "\x1b[21m", "\x1b[41;32m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[32m", "\x1b[1m", "\x1b[21m", "\x1b[42;31m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[33m", "\x1b[1m", "\x1b[21m", "\x1b[43;34m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[34m", "\x1b[1m", "\x1b[21m", "\x1b[44;33m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[35m", "\x1b[1m", "\x1b[21m", "\x1b[45;36m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[36m", "\x1b[1m", "\x1b[21m", "\x1b[46;35m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[37m", "\x1b[1m", "\x1b[21m", "\x1b[47;30m", "\x1b[0m")
	
	//logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	//logrus.SetOutput(colorable.NewColorableStdout())
	
	// http://patorjk.com/software/taag/#p=display&v=1&f=Big&t=GoServe!
	//fmt.Print(color.RedString("   _____       _____                     _"), "\n")
	//logrus.Info("   _____       _____                     _")
	//fmt.Print("\n")
	//fmt.Print("\x1b[31m", "   _____       _____                     _", "\n")
	//fmt.Print(uint16(0x0004 | 0x0000), "31", "   _____       _____                     _", "\n")
	//fmt.Print("^[[31m", "   _____       _____                     _", "\n")
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

