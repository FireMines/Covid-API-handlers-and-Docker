package handler

import "net/http"

/*
Empty handler as default handler
*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "No functionality on root level. Please use paths "+LOCATION_PATH+" or "+COLLECTION_PATH+".", http.StatusOK)
	w.Write([]byte(`<a href="` + (RESOURCE_ROOT_PATH + COVIDCASES) + `">` + COVIDCASES + `<br>`))
	w.Write([]byte(`<a href="` + (RESOURCE_ROOT_PATH + COVIDPOLICY) + `">` + COVIDPOLICY + `<br>`))
	w.Write([]byte(`<a href="` + (RESOURCE_ROOT_PATH + COVIDSTATUS) + `">` + COVIDSTATUS + `<br>`))
	w.Write([]byte(`<a href="` + (RESOURCE_ROOT_PATH + COVIDNOTIFICATIONS) + `">` + COVIDNOTIFICATIONS + `<br>`))

}
