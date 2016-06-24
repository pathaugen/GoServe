

package main


import (
	"io/ioutil"
)



type Domains struct {
	Title		string
	Url			string
	Template	string
}


type Templates struct {
	//Domain	string
	Title		string
	Body		[]byte
}


type Pages struct {
	//Domain	string
	Title		string
	Body		[]byte
}
func (p *Pages) save() error {
	filename := p.Title + ".txt" // .html, .txt, or .md options
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Pages, error) {
	filename := title + ".txt" // .html, .txt, or .md options
	body, err := ioutil.ReadFile(filename)
	if err != nil {
	    return nil, err
	}
	return &Pages{Title: title, Body: body}, nil
}

