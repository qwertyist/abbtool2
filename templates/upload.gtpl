<html>
<head>
  <meta charset="UTF-8">
  <title>UTX - Ladda upp och konvertera förkortningslistor</title>
</head>
<body>
  Format som fungerar är Protype, <strike>Text-On-Top och IllumiType</strike>
  <form enctype="multipart/form-data" action="http://127.0.0.1:3434" method="post">
    <input type="file" name="file" />
    <input type="hidden" name="token" value="{{.}}"/>
    <input type="submit" value="upload" />
  </form>
</body>
</html>
