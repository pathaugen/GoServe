

package main


import (
	"fmt"
	"net/http"
	//"regexp"
	"strings"
	"html/template"
)


func handlerWeb(w http.ResponseWriter, r *http.Request) {
	//pageRequested := r.URL.Path[1:]
	
	// Handle root requests by mapping to "index"
	//if pageRequested == "" {pageRequested = "index"}
	
	// Strip and analyze the dot notation (.md, .html, .text, .edit)
	//extensionArr := strings.Split(pageRequested, ".")
	pageArr := strings.Split(r.URL.Path[1:], ".")
	pageRequested := "index"
	//extension := ""
	extensionRequested := ""
	if len(pageArr) == 2 {
		pageRequested = pageArr[0]
		extensionRequested = pageArr[1]
	} else {
		pageRequested = pageArr[0]
	}
	
	if pageRequested != "favicon.ico" { // Browsers often send a second request to favicon.ico while pulling a URL
		
		// Display feedback if monitoring flag is true
		if monitoringRequests {
			//fmt.Print("\n")
			fmt.Print("function webHandler triggered: ", pageRequested) // http://localhost:8080/parentfolder/subfolder/file.ext
			fmt.Print("\n")
		}
		
	    //fmt.Fprintf(w, "<h1>URL Requested</h1><p>URL Requested: %s!</p>", pageRequested)
	    //fmt.Fprintf(w, "<h1>URL To Filename</h1><p>%s</p>", convUrlToFilename(pageRequested))
	    
	    // Attempt to load page
	    //pageLoaded, _ := loadPage(convUrlToFilename(pageRequested)) // +".txt"
	    pageLoaded, _ := loadPage(pageRequested) // +".txt"
	    if pageLoaded == nil { // If page does not exist (can not be loaded)
		    if extensionRequested == "create" {
		    	// If user is an administrator, ask if they want to create or generate the page
		    	// xxx
		    	
		    	// Create Page
		    	fmt.Fprintf(w, createPage(pageRequested))
		    	
		    	// Display page or edit page
		    	
		    	// RE-attempt load page after creation
			    //pageLoaded, _ := loadPage(convUrlToFilename(pageRequested)) // +".txt"
			    pageLoaded, _ := loadPage(pageRequested) // +".txt"
		    	fmt.Fprintf(w, "<h1>loadPage</h1><div>%s</div>", string(pageLoaded.Body))
		    } else if extensionRequested != "" {
			    // if another extension is on the URL, strip it
			    
			    // Redirect to URL without extension
				http.Redirect(w, r, "/"+pageRequested, http.StatusFound)
		    } else {
		    	// 404 delivery on non-existant page
		    	pageLoaded := &Pages{Url: pageRequested}
		    	t, _ := template.New("edit").Parse(templatePage404())
				t.Execute(w, pageLoaded)
		    }
		    
	    } else {
			// Page exists, handle the extension request being made to the page
			
		    if extensionRequested == "create" {
		    	// if the page exists, no reason to create it again
		    	
		    	// Redirect to URL without extension
				http.Redirect(w, r, "/"+pageRequested, http.StatusFound)
				
			} else if extensionRequested == "edit" {
				// Edit the page content
				
			    //fmt.Fprintf(w, templatePageEdit(), string(pageLoaded.Body))
			    
			    // New template style using "html/template" package
		    	t, _ := template.New("edit").Parse(templatePageEdit())
				t.Execute(w, pageLoaded)
					
		    } else if extensionRequested == "save" {
		    	// If post to this extension, save this page else redirect
		    	
		    	if r.FormValue("body") == "" {
		    		// Redirect to URL without extension
					http.Redirect(w, r, "/"+pageRequested, http.StatusFound)
		    	} else {
			    	newBody := r.FormValue("body")
			    	
					//p := &Page{Title: title, Body: []byte(body)}
					pageLoaded.Body = []byte(newBody)
					pageLoaded.save()
					//http.Redirect(w, r, "/view/"+title, http.StatusFound)
					
			    	fmt.Fprintf(w, "<h1>loadPage SAVE</h1><div>%s</div>", string(pageLoaded.Body))
		    	}
		    	
		    } else if extensionRequested == "text" { // TODO
		    	// Pure text version of page content (rendered as HTML)
		    	fmt.Fprintf(w, "<h1>loadPage TEXT</h1><div>%s</div>", string(pageLoaded.Body))
		    	
		    } else if extensionRequested == "css" { // TODO
		    	// CSS content for HTML pages
		    	fmt.Fprintf(w, "<h1>loadPage CSS</h1><div>%s</div>", string(pageLoaded.Body))
		    	
		    } else if extensionRequested == "js" { // TODO
		    	// JS content for HTML pages
		    	fmt.Fprintf(w, "<h1>loadPage JS</h1><div>%s</div>", string(pageLoaded.Body))
		    	
		    } else if extensionRequested == "template" { // TODO
		    	// Content of template without page content loaded
		    	fmt.Fprintf(w, "<h1>loadPage TEMPLATE</h1><div>%s</div>", string(pageLoaded.Body))
		    	
		    } else if extensionRequested != "" {
		    	// Strip extensions not supported
				http.Redirect(w, r, "/"+pageRequested, http.StatusFound) // Redirect to URL without extension
				
	    	} else {
	    		// Display page normally
	    		
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

