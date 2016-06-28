

package main


import (
	//"fmt"
)


func templatePageEdit() (string) {
	template := `
<div>
	<h1>loadPage EDIT</h1>
	POST to: /{{.Url}}.save
	<div>
	    <form action="/{{.Url}}.save" method="POST">
			<div>
				<textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea>
			</div>
			<div><input type="submit" value="Save"></div>
		</form>
	</div>
</div>`
	return template
}

