<html>
<head>
<meta charset="UTF-8">
<title>UTX - Konverterade listor</title>
</head>
<body>
<h1>Importerade listor</h1>
Listor: 
{{ range $key, $value := . }}
  <h3>{{ $key }}</h3>
  <textarea>{{ range $i, $v := . }}{{ $v }}
{{end}}</textarea>
{{ end }}
</body>
</html>
