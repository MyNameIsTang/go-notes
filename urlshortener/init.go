package urlshortener

import (
	"context"
	"fmt"
	"net/http"
	"text/template"

	"google.golang.org/api/urlshortener/v1"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)

	http.ListenAndServe("localhost:8080", nil)
}

// the template used to show the forms and the results web page to the user
var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`))

func root(w http.ResponseWriter, r *http.Request) {
	rootHtmlTmpl.Execute(w, nil)
}

func short(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("longUrl")
	ctx := context.Background()
	urlshortenerSvc, _ := urlshortener.NewService(ctx)
	url, _ := urlshortenerSvc.Url.Insert(&urlshortener.Url{LongUrl: longUrl}).Do()
	fmt.Println(longUrl)
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Shortened version of %s is : %s",
		longUrl, url.Id))
}

func long(w http.ResponseWriter, r *http.Request) {
	shortUrl := "http://goo.gl" + r.FormValue("shortUrl")
	urlshortenerSvc, _ := urlshortener.NewService(nil)
	url, err := urlshortenerSvc.Url.Get(shortUrl).Do()
	if err != nil {
		fmt.Println("error: %v", err)
		return
	}
	rootHtmlTmpl.Execute(w, fmt.Sprintf("Longer version of %s is : %s",
		shortUrl, url.LongUrl))
}
