

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
	Url			string // e.g. /path/to/content (sans extension)
	//Title		string
	Type		string // html, text, markdown
	Css			[]byte // loaded from /path/to/content.css
	Js			[]byte // loaded from /path/to/content.js
	Body		[]byte
}
func (p *Pages) save() error {
	//filename := p.Title + ".txt" // .html, .txt, or .md options
	//filename := p.Url
	filename := convUrlToFilename(p.Url)+".txt"
	if !checkPathExists(resourcesLocation) {
		// Announce that the resources/ folder does not exist
		os.MkdirAll(resourcesLocation, os.ModePerm)
		// Announce success that resources folder was created
	}
	return ioutil.WriteFile(resourcesLocation+filename, p.Body, 0600)
}
func loadPage(url string) (*Pages, error) {
	//filename := title + ".txt" // .html, .txt, or .md options
	filename := convUrlToFilename(url)
	body, err := ioutil.ReadFile(resourcesLocation+filename+".txt")
	if err != nil {
		return nil, err
	}
	return &Pages{Url: url, Body: body}, nil
}

