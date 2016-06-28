

package main


import (
	//"fmt"
)


func templatePage404() (string) {
	template := `
<div>
	<h1>404 Page Not Found</h1>
	<div>
		<h2>URL Requested</h2>
		{{.Url}}
	</div>
	<div>
	    The page you requested does not exist.
	    <br />Visit: <a href="/{{.Url}}.create">/{{.Url}}.create</a> to create this page.
	</div>
</div>`
	return template
}

