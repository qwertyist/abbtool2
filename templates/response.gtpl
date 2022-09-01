<html>
<head>
<meta charset="UTF-8">
  <title>UTX - Konverterade listor</title>
  <link rel="stylesheet" href="/styles/style.css">
</head>
<body>
  <div class="converter">
  <h2>Konverterade förkortningslistor</h2>
  <p>Om allt fungerar korrekt så går dessa att importera i undertextningsprogrammet.</p>
  {{ range $key, $value := . }}
    <h3>{{ $key }}</h3>
    <textarea rows="8" cols="40">{{ range $i, $v := . }}{{ $v }}
{{end}}</textarea>
  {{ end }}
  </div>
</body>
</html>
