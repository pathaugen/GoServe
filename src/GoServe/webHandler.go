

package main


import (
	"fmt"
	"net/http"
	//"regexp"
	//"strings"
)


func webHandler(w http.ResponseWriter, r *http.Request) {
	pageRequested := r.URL.Path[1:]
	if pageRequested == "" {pageRequested = "index"}
	if pageRequested != "favicon.ico" {
		if monitoring {
			//fmt.Print("\n")
			fmt.Print("function webHandler triggered: ", pageRequested) // http://localhost:8080/parentfolder/subfolder/file.ext
			fmt.Print("\n")
		}
		
	    fmt.Fprintf(w, "<h1>URL Requested</h1><p>URL Requested: %s!</p>", pageRequested)
	    fmt.Fprintf(w, "<h1>URL To Filename</h1><p>%s</p>", urlToFilename(pageRequested))
	    
	    
	    // Attempt to load page
	    pageLoaded, _ := loadPage(urlToFilename(pageRequested)+".txt")
	    if pageLoaded == nil {
	    	// Create Page
	    	fmt.Fprintf(w, createPage(pageRequested))
	    } else {
		    //fmt.Println(string(p2.Body))
		    fmt.Fprintf(w, "<h1>loadPage</h1><div>%s</div>", string(pageLoaded.Body))
	    }
	    
	    
	}
}

