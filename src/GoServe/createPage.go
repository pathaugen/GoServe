

package main


import (
	//"fmt"
)


func createPage(pageRequested string) (string) {
	output := ""
	
    //fmt.Fprintf(w, "<h1>createPage</h1><div>Page (%s) does not exist!</div>", urlToFilename(pageRequested))
    output += "<h1>createPage</h1><div>Page ("+urlToFilename(pageRequested)+") does not exist!</div>"
    
    //fmt.Fprintf(w, "<div>Attempting to create page: %s</div>", urlToFilename(pageRequested))
    output += "<div>Attempting to create page: "+urlToFilename(pageRequested)+"</div>"
    
	pageCreated := &Pages{Url: pageRequested, Body: []byte("This is a sample page for: "+pageRequested)}
    pageCreated.save()
    
    //fmt.Fprintf(w, "<div>Page (%s) Created!</div>", urlToFilename(pageRequested))
    output += "<div>Page ("+urlToFilename(pageRequested)+") Created!</div>"
    
	return output
}

