package paulabramwell

import (
	"net/http"
	ttemplate "text/template"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/faq", faq)
	http.HandleFunc("/my-approach", myApproach)
	http.HandleFunc("/therapy-for-adolescents", therapyForAdolescents)
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

func myApproach(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var htmlTempl = ttemplate.Must(ttemplate.ParseFiles("my-approach.html"))
	htmlTempl.Execute(w, "")
}

func therapyForAdolescents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var htmlTempl = ttemplate.Must(ttemplate.ParseFiles("therapy-for-adolescents.html"))
	htmlTempl.Execute(w, "")
}
