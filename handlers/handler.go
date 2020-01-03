package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/YAWAL/ETLUNL/E"
	"github.com/YAWAL/ETLUNL/T"
	"github.com/YAWAL/ETLUNL/config"
	. "github.com/YAWAL/ETLUNL/logging"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) // remove this
		return
	}
	err = t.Execute(w, "Hello World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error())) // remove this
		return
	}

}

func Extract(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("static/extract.html"))
	if err := t.Execute(w, "Hello World!"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Log.Errorf("t.Execute: %v", err)
		w.Write([]byte(err.Error())) // remove this
		return
	}
}

func Transform(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded")
	if err != nil {
		Log.Errorf("error during uploading file: %v", err)
		return
	}

	preparedData, err := T.Transform(E.ProcessRawData(file))
	if err != nil {
		Log.Errorf("error during transformation of data: %v", err)
		return
	}

	// fmt.Println(preparedData)
	Log.Infof("total %d items has been transformed", len(preparedData))
	//w.Write([]byte("data has been transformed"))

	t := template.Must(template.ParseFiles("static/load.html"))
	if err := t.Execute(w, "Hello World!"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		Log.Errorf("t.Execute: %v", err)
		w.Write([]byte(err.Error())) // remove this
		return
	}
	

}

func PGConfigs(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/pg_configs.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	configs := config.Postgres{
		Dialect:      r.FormValue("dialect"),
		User:         r.FormValue("user"),
		DataBaseName: r.FormValue("db_name"),
		SSLMode:      r.FormValue("sslMode"),
		Password:     r.FormValue("pass"),
	}

	// do something with details
	fmt.Println(configs)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func MongoConfigs(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/mongo_configs.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	configs := config.Postgres{
		Dialect:      r.FormValue("dialect"),
		User:         r.FormValue("user"),
		DataBaseName: r.FormValue("db_name"),
		SSLMode:      r.FormValue("sslMode"),
		Password:     r.FormValue("pass"),
	}

	// do something with details
	fmt.Println(configs)

	tmpl.Execute(w, struct{ Success bool }{true})
}
