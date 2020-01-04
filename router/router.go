package router

import (
	"github.com/YAWAL/ETLUNL/handlers"
	"github.com/gorilla/mux"
)

func New() (r *mux.Router) {
	r = mux.NewRouter()
	api := r.PathPrefix("/etl").Subrouter()
	api.HandleFunc("/", handlers.Index)
	api.HandleFunc("/extract", handlers.ExtractPage)
	api.HandleFunc("/download", handlers.Download)
	//api.HandleFunc("/download/{lastGame}", handlers.Download)
	api.HandleFunc("/transform", handlers.UploadAndTransform)
	api.HandleFunc("/pg_configs", handlers.PGConfigs)
	api.HandleFunc("/mongo_configs", handlers.MongoConfigs)
	return r
}
