package templates

import (
	"html/template"
)

var Table = template.Must(template.New("table").Parse(`
{{range $r := .Records}}
<tr>
	<td>{{$r.CreatedAt}}</td>
	<td>{{$r.Type}}</td>
	<td>{{$r.Value}}</td>
</tr>
{{end}}
`))

var Index = template.Must(template.New("index").Parse(`
`))
