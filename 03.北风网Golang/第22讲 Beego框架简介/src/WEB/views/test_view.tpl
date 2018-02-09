<!DOCTYPE html>
<html>
    <title>{{.title}}</title>
<body>
    {{if .IsDisplay}}
        <em>{{.Content1}}</em>
    {{else}}
        <em>{{.Content2}}</em>
    {{end}}
    <br/>
    {{range .Users}}
        {{.Username}}{{$.len}}
    {{end}}
</body>
</html>
