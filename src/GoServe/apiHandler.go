

package main


import (
	"fmt"
	"net/http"
)


func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>apiHandler</h1><div>%s</div>", "This is the API Handler.")
}
