package notifications

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	consts "covidAss2/variables"
)

// Webhook DB
var Webhooks = []consts.WebhookRegistration{}

const collection = "webhooks"

/*
 *	Entry point handler for Location information
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

/**
 *	Handles the post request for notifications
 */
func notificationPostRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// Expects incoming body in terms of WebhookRegistration struct
	webhook := consts.WebhookRegistration{}
	err := json.NewDecoder(r.Body).Decode(&webhook)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), http.StatusBadRequest)
	}

	id, _, err := consts.Client.Collection(collection).Add(consts.Ctx,
		map[string]interface{}{
			"url":     webhook.Url,
			"country": webhook.Country,
			"calls":   webhook.Calls,
		})
	webhook.Weebhook_ID = id.ID
	Webhooks = append(Webhooks, webhook)

	fmt.Println("Webhook_id for POST just done: ", webhook.Weebhook_ID)
	fmt.Println("Webhook " + webhook.Url + " has been registered.")
	http.Error(w, strconv.Itoa(len(Webhooks)-1), http.StatusCreated)

	if err != nil {
		// Error handling
		http.Error(w, "Error when adding message "+webhook.Weebhook_ID+", Error: "+err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Println("Entry added to collection. Identifier of returned document: " + id.ID)
		// Returns document ID in body
		http.Error(w, id.ID, http.StatusCreated)
		return
	}
}

/**
 *	Handles the get request for notifications
 */
func notificationGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	// Gets the url data and checks if user have entered a value
	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	if urlLastVal == "notifications" {
		err := json.NewEncoder(w).Encode(Webhooks)
		if err != nil {
			http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	// Goes through all webhooks and checks if they match a stated value in the url
	// If it matches one found, then print out its data
	for i := range Webhooks {
		if Webhooks[i].Weebhook_ID == urlLastVal {
			err := json.NewEncoder(w).Encode(Webhooks[i])
			if err != nil {
				http.Error(w, "Something went wrong: "+err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	// If none of the statements above happens, return status not found
	http.Error(w, "No webhook with that ID exists!", http.StatusNotFound)

}

/**
 *	Handles the delete request for notifications
 */
func notificationDeleteRequest(w http.ResponseWriter, r *http.Request) {
	urlLastVal := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	fmt.Println(urlLastVal)
	r.Header.Add("content-type", "application/json")
	if urlLastVal == "notifications" {
		http.Error(w, "Looks like you forgot to add a webhook_id! Please do so next time ;)", http.StatusBadRequest)
		return
	}
	for i := range Webhooks {
		if Webhooks[i].Weebhook_ID == urlLastVal {
			_, err := consts.Client.Collection(collection).Doc(urlLastVal).Delete(consts.Ctx)
			if err != nil {
				// Handle any errors in an appropriate way, such as returning them.
				log.Printf("An error has occurred: %s", err)
			}
			Webhooks = append(Webhooks[:i], Webhooks[i+1:]...)
			fmt.Println("Deleting webhook success", http.StatusOK)

			return
		}
	}

	// If none of the statements above happens, return status not found
	http.Error(w, "No webhook with that ID exists!", http.StatusNotFound)

}
