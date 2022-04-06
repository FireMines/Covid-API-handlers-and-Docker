package notifications

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	consts "covidAss2"
)

// Webhook DB
var webhooks = []consts.WebhookRegistration{}

/*
Entry point handler for Location information
*/
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Expects incoming body in terms of WebhookRegistration struct
		webhook := consts.WebhookRegistration{}
		err := json.NewDecoder(r.Body).Decode(&webhook)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
		}
		randString := RandomString(consts.GenLength)
		hash := hmac.New(sha512.New, []byte(randString))
		webhook.Weebhook_ID = hex.EncodeToString(hash.Sum(nil))

		webhooks = append(webhooks, webhook)
		// Note: Approach does not guarantee persistence or permanence of resource id (for CRUD)
		//fmt.Fprintln(w, len(webhooks)-1)
		fmt.Println("Webhook " + webhook.Url + " has been registered.")
		http.Error(w, strconv.Itoa(len(webhooks)-1), http.StatusCreated)

		notificationPostRequest(w, r)

	case http.MethodGet:
		notificationGetRequest(w, r)
		err := json.NewEncoder(w).Encode(webhooks)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
		}

	case http.MethodDelete:

	default:
		http.Error(w, "Method not supported. Currently only GET and POST are supported.", http.StatusNotImplemented)
		return
	}
}

func notificationPostRequest(w http.ResponseWriter, r *http.Request) {

}

func notificationGetRequest(w http.ResponseWriter, r *http.Request) {

}

func RandomString(n int) string {
	var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
