

package main


import (
	"fmt"
	"net/http"
	//"regexp"
	"strings"
	"html/template"
)


func handlerWeb(w http.ResponseWriter, r *http.Request) {
	pageRequested := r.URL.Path[1:]
	
	// Handle root requests by mapping to "index"
	if pageRequested == "" {pageRequested = "index"}
	
	// Strip and analyze the dot notation (.md, .html, .text, .edit)
	extensionArr := strings.Split(pageRequested, ".")
	extension := ""
	if len(extensionArr) == 2 {
		extension = extensionArr[1]
	}
	
	if pageRequested != "favicon.ico" { // Browsers often send a second request to favicon.ico while pulling a URL
		
		// Display feedback if monitoring flag is true
		if monitoring {
			//fmt.Print("\n")
			fmt.Print("function webHandler triggered: ", pageRequested) // http://localhost:8080/parentfolder/subfolder/file.ext
			fmt.Print("\n")
		}
		
	    fmt.Fprintf(w, "<h1>URL Requested</h1><p>URL Requested: %s!</p>", pageRequested)
	    fmt.Fprintf(w, "<h1>URL To Filename</h1><p>%s</p>", convUrlToFilename(pageRequested))
	    
	    
	    // Attempt to load page
	    pageLoaded, _ := loadPage(convUrlToFilename(pageRequested)+".txt")
	    if pageLoaded == nil { // If page does not exist (can not be loaded)
	    	// Create Page
	    	fmt.Fprintf(w, createPage(pageRequested))
	    	
	    	// Display page or edit page
	    	
	    	// RE-attempt load page after creation
		    pageLoaded, _ := loadPage(convUrlToFilename(pageRequested)+".txt")
	    	fmt.Fprintf(w, "<h1>loadPage</h1><div>%s</div>", string(pageLoaded.Body))
	    } else {
	    	// Display page or edit page
	    	
		    //fmt.Println(string(p2.Body))
		    // Utilize the correct page template
		    if extension == "edit" {
		    	// Edit the page content
			    //fmt.Fprintf(w, templatePageEdit(), string(pageLoaded.Body))
			    
			    // New template style using "html/template" package
		    	t, _ := template.New("edit").Parse(templatePageEdit())
				t.Execute(w, pageLoaded)
		    } else if extension == "text" {
		    	// Pure text version of page content (rendered as HTML)
		    	fmt.Fprintf(w, "<h1>loadPage TEXT</h1><div>%s</div>", string(pageLoaded.Body))
		    } else if extension == "content" {
		    	// Just page content without template applied
		    	fmt.Fprintf(w, "<h1>loadPage CONTENT</h1><div>%s</div>", string(pageLoaded.Body))
		    } else if extension == "template" {
		    	// Content of template without page content loaded
		    	fmt.Fprintf(w, "<h1>loadPage TEMPLATE</h1><div>%s</div>", string(pageLoaded.Body))
		    } else {
		    	//fmt.Fprintf(w, "<h1>loadPage</h1><div>%s</div>", string(pageLoaded.Body))
		    	//fmt.Fprintf(w, templatePageLoad(), string(pageLoaded.Body))
		    	
		    	// New template style using "html/template" package
		    	t, _ := template.New("page").Parse(templatePageLoad())
				//_ = t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
				t.Execute(w, pageLoaded)
		    }
	    }
	    
	    
	}
}

