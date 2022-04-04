package handler

import (
	"net/http"
)

/*
Entry point handler for Location information
*/
func DiagHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		//diagHandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}
