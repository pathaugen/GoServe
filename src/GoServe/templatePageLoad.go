

package main


import (
	//"fmt"
)


func templatePageLoad() (string) {
	template := `
<div>
	<h1>loadPage LOAD</h1>
	<div>
		<h2>URL</h2>
		{{.Url}}
	</div>
	<div>
	    {{printf "%s" .Body}}
	</div>
</div>`
	return template
}

