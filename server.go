package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		body, _ := ioutil.ReadAll(r.Body)
		w.Write([]byte("debut body::"))
		w.Write(body)
		w.Write([]byte("::fin body\n"))
		w.Write([]byte(token))
	})
	http.HandleFunc("/api2", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		w.Write([]byte("debut body::"))
		w.Write(body)
		w.Write([]byte("::fin body\n"))
		fmt.Fprint(w, r.Form)
		r.ParseForm()
		fmt.Fprint(w, r.Form)
	})
	http.HandleFunc("/api3", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		w.Write([]byte("debut body::"))
		w.Write(body)
		w.Write([]byte("::fin body\n"))
		fmt.Fprint(w, r.Form)
		fmt.Fprint(w, r.PostForm)
		r.ParseForm()
		token := r.PostFormValue("token")
		w.Write([]byte(token))
		w.Write([]byte{'\n'})

	})
	http.HandleFunc("/api4", func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		fmt.Fprintf(w, "TOKEN : |%s|\n", token)
	})
	http.HandleFunc("/api5", func(w http.ResponseWriter, r *http.Request) {
		token := r.PostFormValue("token")
		fmt.Fprintf(w, "TOKEN : |%s|\n", token)
	})
	http.HandleFunc("/api6", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "|BODY START|%s|BODY END|\n", string(body))

		token := r.PostFormValue("token")
		fmt.Fprintf(w, "TOKEN : |%s|\n", token)

		body2, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "|BODY START|%s|BODY END|\n", string(body2))
	})
	http.HandleFunc("/api7", func(w http.ResponseWriter, r *http.Request) {
		token := r.PostFormValue("token")
		fmt.Fprintf(w, "TOKEN : |%s|\n", token)

		body2, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "|BODY START|%s|BODY END|\n", string(body2))
	})
	http.ListenAndServe(":8090", nil)
}
