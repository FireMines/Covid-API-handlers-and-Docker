package notifications

import (
	"net/http"
)

/*
Entry point handler for Location information
*/
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		//handleMapPostRequest(w, r)
	case http.MethodGet:
		notificationGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET and POST are supported.", http.StatusNotImplemented)
		return
	}
}

func notificationGetRequest(w http.ResponseWriter, r *http.Request) {

}
