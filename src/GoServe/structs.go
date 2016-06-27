

package main


import (
	"os"
	"io/ioutil"
)



type Domains struct {
	Domain		string
	Title		string
	Template	string
}


type Templates struct {
	Domain		string
	Title		string
	Body		[]byte
}


type Pages struct {
	//Domain		string
	Url			string
	//Title		string
	Body		[]byte
	Css			[]byte
}
func (p *Pages) save() error {
	//filename := p.Title + ".txt" // .html, .txt, or .md options
	//filename := p.Url
	filename := convUrlToFilename(p.Url)+".txt"
	if !checkPathExists("resources/") {
		// Announce that the resources/ folder does not exist
		os.MkdirAll("resources/", os.ModePerm)
		// Announce success that resources folder was created
	}
	return ioutil.WriteFile("resources/"+filename, p.Body, 0600)
}
func loadPage(url string) (*Pages, error) {
	//filename := title + ".txt" // .html, .txt, or .md options
	filename := url
	body, err := ioutil.ReadFile("resources/"+filename)
	if err != nil {
	    return nil, err
	}
	return &Pages{Url: url, Body: body}, nil
}

