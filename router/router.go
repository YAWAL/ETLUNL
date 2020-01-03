package router

import (
	"github.com/YAWAL/ETLUNL/handlers"
	"github.com/gorilla/mux"
)

func InitRouter() (r *mux.Router) {
	r = mux.NewRouter()
	api := r.PathPrefix("/etl").Subrouter()
	api.HandleFunc("/", handlers.Index)
	api.HandleFunc("/extract", handlers.Extract)
	api.HandleFunc("/transform", handlers.Transform)

	api.HandleFunc("/pg_configs", handlers.PGConfigs)
	api.HandleFunc("/mongo_configs", handlers.MongoConfigs)

	return r
}
