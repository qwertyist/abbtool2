<!DOCTYPE html>
<head lang="sv">
  <meta charset="UTF-8">
  <title>UTX - Verktygslåda | {{ if .Title }}Hello{{else}}hehe{{ end }}</title>
  <link rel="stylesheet" href="styles/style.css">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  {{ block "head" .}}{{ end }}
</head>

<body>
	<div class="credits">
    <img src="images/stfhsk_logo.png" />
    Sidan drivs av Undertextningsutbildningen på Södertörns Folkhögskola
  </div>
  <div class="content">
    {{ template "content" . }}
  </div>
  {{ block "foot" . }}{{ end }}
</body>
</html>
