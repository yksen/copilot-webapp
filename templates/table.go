package templates

import "html/template"

var Table = template.Must(template.New("table").Parse(`
	<table>
		<tr>
			<th>Id</th>
			<th>CreatedAt</th>
			<th>RecordType</th>
			<th>Value</th>
		</tr>
		{{range $r := .Records}}
		<tr>
			<td>{{$r.Id}}</td>
			<td>{{$r.CreatedAt}}</td>
			<td>{{$r.RecordType}}</td>
			<td>{{$r.Value}}</td>
		</tr>
		{{end}}
	</table>
`))
