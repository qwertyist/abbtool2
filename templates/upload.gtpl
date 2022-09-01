<html>
<head>
  <meta charset="UTF-8">
  <link rel="stylesheet" href="/styles/style.css">
  <title>UTX - Ladda upp och konvertera förkortningslistor</title>
</head>
<body>
<div class="converter">
<h2>Ladda upp och konvertera förkortningslistor till Shortform-format</h2>
  <p>Format som fungerar är Protype, Text-On-Top och IllumiType.</p>
  <form enctype="multipart/form-data" action="/" method="post">
    <input lang="sv-SE" type="file" name="file" value="Välj fil" accept=".zip,.json" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
  </form>
</div>
</body>
</html>
