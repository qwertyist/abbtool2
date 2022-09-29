{{ define "title" }}Startsida{{ end }}
{{ define "content" }}
  <div class="container banner">
    <img src="/images/banner_main.png" />
  </div>
  <div class="container">
    <img src="/images/banner_converter.png" />
    <h2>Ladda upp och konvertera förkortningslistor till Shortform-format</h2>
    <p>Format som fungerar är Protype, Text-On-Top och IllumiType.</p>
    <form enctype="multipart/form-data" action="/convert" method="post">
      <input lang="sv-SE" type="file" name="file" value="Välj fil" accept=".zip,.json" />
      <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="upload" />
    </form>
  </div>
{{ end }}
