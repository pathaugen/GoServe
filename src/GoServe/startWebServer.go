

package main


import (
	"os"
	"fmt"
	"net/http"
	
	// Switch to Vendoring
	//"GoServe/external/ansicolor-master"
	"ansicolor-master"
)


func startWebServer(serverPort string) (bool) {
	
	consoleColor := ansicolor.NewAnsiColorWriter(os.Stdout) // Initialize Windows 10 console coloring
	
	// Save settings to .cfg file for auto-resume
	// xxx
	
	// ========== ========== ========== ========== ==========
	// Multiple webservers can be created and started
	// > webserver list -> List all running webservers
	// > webserver start 8080 -> Start a webserver on the specified port
	// webservers.cfg -> configuration file with webservers that will be run on startup
	fmt.Print("Starting Webserver on port "+serverPort+".. please wait..")
	fmt.Print("\n")
	//http.HandleFunc("/", handler)
    //http.ListenAndServe(":8080", nil)
    
    // Resources Location
    resourcesLocation = "goserve-resources/port-"+serverPort+"/localhost/" // Set the resourcesLocation for this server and domain
    // resourcesLocation = "goserve-resources/port-[PORT]/[DOMAIN]/" // Set the resourcesLocation for this server and domain
    
    // Can have multiple servers on different ports by duplicating this code
    serverX := http.NewServeMux()
    serverX.HandleFunc("/", handlerWeb)
    serverX.HandleFunc("/api/", handlerApi)
    go func() {
		http.ListenAndServe(":"+serverPort, serverX)
    }()
    fmt.Fprintf(consoleColor, "\x1b[32m\x1b[1m") // Green, bold
	fmt.Print("SUCCESS: Webserver STARTED on port "+serverPort)
	fmt.Fprintf(consoleColor, "\x1b[0m") // Reset colors, bold
	fmt.Print("\n")
	fmt.Print("Visit: http://localhost:"+serverPort)
	fmt.Print("\n\n")
	// ========== ========== ========== ========== ==========
	
	return true
}

