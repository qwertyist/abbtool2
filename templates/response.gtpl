{{ define "content" }}
  <div class="container">
    <img src="images/banner_converter.png" />
  <p>Om allt fungerar korrekt så går dessa att importera i undertextningsprogrammet.</p>
  {{ range $key, $value := .Lists }}
    <h3>{{ $key }}</h3>
    <textarea rows="8" cols="40">{{ range $i, $v := . }}{{ $v }}
{{end}}</textarea>
  {{ end }}
  </div>
{{ end }}
