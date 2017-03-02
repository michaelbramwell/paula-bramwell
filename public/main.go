package paulabramwell

import (
	"net/http"
	ttemplate "text/template"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/faq", faq)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var htmlTempl = ttemplate.Must(ttemplate.ParseFiles("index.html"))
	htmlTempl.Execute(w, "")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var htmlTempl = ttemplate.Must(ttemplate.ParseFiles("faq.html"))
	htmlTempl.Execute(w, "")
}
