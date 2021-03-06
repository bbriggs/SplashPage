package components

// ToHTMLPage consumes the template and struct to produce a html object
func ToHTMLPage(s string, style string) string {

	return `
	<!doctype html>
	<html>
	<head>
		<style>
		body {
			background-size: 100% 100%;
		}` +
		style +
		`</style>
	</head>
		<body>` +
		s +
		`</body>
	</html>`

}
