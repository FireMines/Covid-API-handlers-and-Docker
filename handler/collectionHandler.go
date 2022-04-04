package handler

import (
	"net/http"
)

/*
Entry point handler for collection information
*/
func CollectionHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		//handleMapPostRequest(w, r)
	case http.MethodGet:
		//handleMapGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only POST and GET are supported.", http.StatusNotImplemented)
		return
	}

}
