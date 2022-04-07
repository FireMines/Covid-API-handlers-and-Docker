package handler

import (
	consts "covidAss2/variables"
	"net/http"
)

/*
Empty handler as default handler
*/
func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	//http.Error(w, "No functionality on root level. Please use paths "+LOCATION_PATH+" or "+COLLECTION_PATH+".", http.StatusOK)
	w.Write([]byte(`<a href="` + (consts.RESOURCE_ROOT_PATH + consts.COVIDCASES) + `">` + consts.COVIDCASES + `<br>`))
	w.Write([]byte(`<a href="` + (consts.RESOURCE_ROOT_PATH + consts.COVIDPOLICY) + `">` + consts.COVIDPOLICY + `<br>`))
	w.Write([]byte(`<a href="` + (consts.RESOURCE_ROOT_PATH + consts.COVIDSTATUS) + `">` + consts.COVIDSTATUS + `<br>`))
	w.Write([]byte(`<a href="` + (consts.RESOURCE_ROOT_PATH + consts.COVIDNOTIFICATIONS) + `">` + consts.COVIDNOTIFICATIONS + `<br>`))
}
