

package main


import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)


var exit bool = false


func main() {
	
	
	for !exit {
		waitForInput()
	}
	
	
	if exit {
		fmt.Print("\n\n")
		fmt.Print("Application is exiting. Press [enter] to continue.")
		fmt.Print("\n\n")
		fmt.Print("[exiting] > ")
		
		consolereader := bufio.NewReader(os.Stdin)
		_, err := consolereader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	
}

