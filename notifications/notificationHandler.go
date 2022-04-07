package notifications

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	consts "covidAss2/variables"
)

// Webhook DB
var Webhooks = []consts.WebhookRegistration{}

/*
Entry point handler for Location information
*/
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		notificationPostRequest(w, r)
	case http.MethodGet:
		notificationGetRequest(w, r)
	case http.MethodDelete:
		notificationDeleteRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET and POST are supported.", http.StatusNotImplemented)
		return
	}
}

func notificationPostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// Expects incoming body in terms of WebhookRegistration struct
	webhook := consts.WebhookRegistration{}
	err := json.NewDecoder(r.Body).Decode(&webhook)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
	}
	randString := RandomString(consts.GenLength)
	hash := hmac.New(sha512.New, []byte(randString))
	webhook.Weebhook_ID = hex.EncodeToString(hash.Sum(nil))

	Webhooks = append(Webhooks, webhook)
	// Note: Approach does not guarantee persistence or permanence of resource id (for CRUD)
	//fmt.Fprintln(w, len(webhooks)-1)

	fmt.Println("Webhook " + webhook.Url + " has been registered.")
	http.Error(w, strconv.Itoa(len(Webhooks)-1), http.StatusCreated)
}

func notificationGetRequest(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("content-type", "application/json")

	err := json.NewEncoder(w).Encode(Webhooks)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
	}
}

func notificationDeleteRequest(w http.ResponseWriter, r *http.Request) {
	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	fmt.Println(urlLastVal)
	r.Header.Add("content-type", "application/json")
	if urlLastVal == "notifications" {
		http.Error(w, "Looks like you forgot to add a webhook_id! Place do so next time ;)", http.StatusBadRequest)
		return
	}
	for i := range Webhooks {
		if Webhooks[i].Weebhook_ID == urlLastVal {
			Webhooks = append(Webhooks[:i], Webhooks[i+1:]...)
			return
		} /* else {
			fmt.Println("No webhook with this id: %s", urlLastVal)
		} */
	}
	/* 	form := url.Values{}
	   	form.Add("webhook_id", strconv.Itoa(5))
	   	body := strings.NewReader(form.Encode())
	   	req, err := http.NewRequest("DELETE", consts.RESOURCE_ROOT_PATH+consts.COVIDNOTIFICATIONS, body)
	   	if err != nil {
	   		http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
	   	}
	   	client := &http.Client{}
	   	resp, err := client.Do(req)
	   	if err != nil {
	   		http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
	   	}
	   	defer resp.Body.Close()
	*/
}
