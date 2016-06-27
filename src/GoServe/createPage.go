

package main


import (
	//"fmt"
	"strings"
)


func createPage(pageRequested string) (string) {
	output := ""
	
	// Strip Extension
	extensionArr := strings.Split(pageRequested, ".")
	pageRequested = extensionArr[0]
	
    //fmt.Fprintf(w, "<h1>createPage</h1><div>Page (%s) does not exist!</div>", urlToFilename(pageRequested))
    output += "<h1>createPage</h1><div>Page ("+convUrlToFilename(pageRequested)+") does not exist!</div>"
    
    //fmt.Fprintf(w, "<div>Attempting to create page: %s</div>", urlToFilename(pageRequested))
    output += "<div>Attempting to create page: "+convUrlToFilename(pageRequested)+"</div>"
    
	pageCreated := &Pages{Url: pageRequested, Body: []byte("This is a sample page for: "+pageRequested)}
    pageCreated.save()
    
    //fmt.Fprintf(w, "<div>Page (%s) Created!</div>", urlToFilename(pageRequested))
    output += "<div>Page ("+convUrlToFilename(pageRequested)+") Created!</div>"
    
	return output
}

