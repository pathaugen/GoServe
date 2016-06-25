

package main


import (
	"fmt"
	"net/http"
	//"regexp"
	"strings"
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
		
	    fmt.Fprintf(w, "<p>Hi there, I love %s!</p>", pageRequested)
	    fmt.Fprintf(w, "<p>THIS IS A SECOND LINE</p>")
	    fmt.Fprintf(w, "<p>%s</p>", urlToFilename(pageRequested))
	    
	    
	    p1 := &Pages{Url: pageRequested, Body: []byte("This is a sample page for: "+pageRequested)}
	    p1.save()
	    
	    p2, _ := loadPage(urlToFilename(pageRequested)+".txt")
	    //fmt.Println(string(p2.Body))
	    fmt.Fprintf(w, string(p2.Body))
	    
	}
}


func urlToFilename(input string) (string) {
	/*
	//re, err := regexp.Compile(`img.*?src=\"(.*?)\.(.*?)\"`)
	re, err := regexp.Compile(`img.*?src=\"(.*?)\.(.*?)\"`) // parentfolder/subfolder/file.ext
	if err != nil {
		//return "", err
		return ""
	}
	indexes := re.FindAllStringSubmatchIndex(input, -1)
	
	output := input
	for _, match := range indexes {
		imgStart := match[2]
		imgEnd := match[3]
		newImgName := strings.Replace(input[imgStart:imgEnd], "m", "a", -1)
		output = output[:imgStart] + newImgName + input[imgEnd:]
	}
	return output //, nil
	*/
	
	longFilenameArr := strings.Split(input, ".")
	longFilename := strings.Replace(longFilenameArr[0], "/", "--", -1)
	
	/*
	arr := strings.Split(input, "/") // Array of all parts of input
	filename := arr[len(arr)-1] // Last part of array
	*/
	
	/*
	filearr := strings.Split(filename, ".")
	file := filearr[0]
	ext := filearr[len(arr)-1]
	*/
	
	return longFilename
}

