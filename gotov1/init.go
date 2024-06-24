package gotov1

import (
	"html/template"
	"net/http"
)

var store = NewURLStore()

var AddForm = template.Must(template.New("AddForm").Parse(`
<html><body>
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
</body>
</html>
`))

var RedirectForm = template.Must(template.New("RedirectForm").Parse(`
<html><body>
<a target="_blank" href="{{. }}">{{. }}</a>
</body>
</html>
`))

func init() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8999", nil)
}

func Redirect(rw http.ResponseWriter, rq *http.Request) {
	path := rq.URL.Path
	key := path[1:]
	if url := store.Get(key); url != "" {
		http.Redirect(rw, rq, url, http.StatusFound)
		return
	}
	http.NotFound(rw, rq)
}

func Add(rw http.ResponseWriter, rq *http.Request) {
	url := rq.FormValue("url")
	if url == "" {
		AddForm.Execute(rw, nil)
		return
	}
	key := store.Put(url)
	RedirectForm.Execute(rw, "http://localhost:8999/"+key)
}
