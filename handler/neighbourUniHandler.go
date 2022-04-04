package handler

import (
	"net/http"
)

/*
Entry point handler for Location information
*/
func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		//neighbourHandleGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}
