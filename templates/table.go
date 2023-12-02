package templates

import "html/template"

var Table = template.Must(template.New("table").Parse(`
	<table>
		<tr>
			<th>Id</th>
			<th>CreatedAt</th>
			<th>Type</th>
			<th>Value</th>
		</tr>
		{{range $r := .Records}}
		<tr>
			<td>{{$r.Id}}</td>
			<td>{{$r.CreatedAt}}</td>
			<td>{{$r.Type}}</td>
			<td>{{$r.Value}}</td>
		</tr>
		{{end}}
	</table>
`))
