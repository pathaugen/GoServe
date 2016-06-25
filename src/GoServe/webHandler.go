

package main


import (
	"fmt"
	"net/http"
)


func webHandler(w http.ResponseWriter, r *http.Request) {
	pageRequested := r.URL.Path[1:]
	if pageRequested != "favicon.ico" {
		if monitoring {
			//fmt.Print("\n")
			fmt.Print("function webHandler triggered: ", pageRequested)
			fmt.Print("\n")
		}
		
	    fmt.Fprintf(w, "Hi there, I love %s!", pageRequested)
	    
	    p1 := &Pages{Title: "test", Body: []byte("This is a sample page for: "+pageRequested)}
	    p1.save()
	    p2, _ := loadPage("test")
	    //fmt.Println(string(p2.Body))
	    fmt.Fprintf(w, string(p2.Body))
	}
}