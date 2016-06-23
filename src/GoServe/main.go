

package main


import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
)


var exit bool = false
var version string = "1.0.0"


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
	
	
	
	for !exit {
		waitForInput()
	}
	
	
	if exit {
		exitHandler()
	}
	
}

