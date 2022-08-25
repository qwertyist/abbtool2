<html>
<head>
  <meta charset="UTF-8">
  <title>UTX - Ladda upp och konvertera förkortningslistor</title>
</head>
<body>
<h2>Ladda upp och konvertera förkortningslistor till Shortform-format</h2>
  <p>Format som fungerar är Protype, Text-On-Top och IllumiType.</p>
  <form enctype="multipart/form-data" action="http://127.0.0.1:3434" method="post">
    <input type="file" name="file" value="Välj fil" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
  </form>
</body>
</html>
