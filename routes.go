package main

import "github.com/gorilla/mux"

func Routes() *mux.Router  {
	r := mux.NewRouter()
	r.Path("/api/v1/vm").HandlerFunc(GetVm).Methods("GET")
	r.Path("/api/v1/vm/{id}").HandlerFunc(GetVmID).Methods("GET")
	r.Path("/api/v1/vm").HandlerFunc(AddVm).Methods("POST")
	r.Path("/api/v1/vm/{id}").HandlerFunc(EditVm).Methods("PUT")
	r.Path("/api/v1/vm/{id}").HandlerFunc(DeleteVm).Methods("DELETE")
	return r
}
