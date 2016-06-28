

package main


import (
	//"fmt"
)


func templatePageEdit() (string) {
	template := `
<style>
	fieldset {
		background-color:	lightblue;
		margin:				2%;
		padding:			2%;
	}
	legend {
		font-size:			1.4em;
		font-weight:		bold;
	}
	label {
		font-size:			1.2em;
		font-weight:		bold;
	}
	input[type=text], textarea {
		width:				100%;
		margin-bottom:		2%;
		padding:			1%;
	}
	input[type=text] {
		font-size:			1.2em;
	}
</style>
<script>
	/* alert('You triggered an alert!'); */
</script>
<div>
	<h1>loadPage EDIT</h1>
	<div>
		<form action="/{{.Url}}.save" method="POST">
			<fieldset>
				<legend>Edit Page: /{{.Url}}</legend>
				<div>
					<label for="title">Title (not yet implemented)</label>
					<input type="text" id="title" />
			    </div>
			    <div style="margin-bottom:2%;">
					<div><input type="radio" name="type" value="html" checked /> HTML Page</div>
					<div><input type="radio" name="type" value="text" /> Text Page</div>
					<div><input type="radio" name="type" value="markdown" /> Markdown Page</div>
			    </div>
				<div>
					<label for="css">CSS</label>
					<textarea name="css" rows="20" cols="80">{{printf "%s" .Css}}</textarea>
				</div>
				<div>
					<label for="js">Javascript</label>
					<textarea name="js" rows="20" cols="80">{{printf "%s" .Js}}</textarea>
				</div>
				<div>
					<label for="body">Body</label>
					<textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea>
				</div>
				<div>
					<input type="submit" value="Save" />
					POST to: /{{.Url}}.save
				</div>
			</fieldset>
		</form>
	</div>
</div>`
	return template
}

