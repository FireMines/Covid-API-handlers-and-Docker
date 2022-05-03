package notifications_test

import (
	"covidAss2/notifications"
	consts "covidAss2/variables"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
 *	Tester for the HandlerCases function.
 *	Mocks any outgoing request during testing.
 */
func TestHandlerNotifications(t *testing.T) {
	// Mock outgoing functions so they return hardcoded results
	Webhooks_original := notifications.Webhooks
	NotificationGetRequest_original := notifications.NotificationGetRequest
	notifications.NotificationGetRequest = func(w http.ResponseWriter, r *http.Request) {
		// Get the mocked webhooks and set the Webhooks var to be it
		content, err := ioutil.ReadFile("webhookTestData.json")
		if err != nil {
			panic("Could not load test data")
		}

		var filteredData []consts.WebhookRegistration
		json.Unmarshal([]byte(content), &filteredData)

		notifications.Webhooks = filteredData

		// ...Then continue as usual
		NotificationGetRequest_original(w, r)
	}

	/**
	 *	To shorten the code, create a temporary function which is repeated.
	 */
	tempFunc := func(testurl string, expected string) {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(http.MethodGet, testurl, nil)
		if err != nil {
			t.Error("Error when creating new request", err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		notifications.NotificationHandler(rr, req)
		result := rr.Result()
		defer result.Body.Close()

		// Format it into data
		data, err := ioutil.ReadAll(result.Body)
		if err != nil {
			t.Error("Error when creating new request", err)
		}

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check if expected result was the same as the actualy response
		data_str := string(data)
		data_str = data_str[:len(data_str)-1]
		if data_str != expected {
			t.Errorf("Handler did not return expected value: got '%v' want '%v'", data_str, expected)
		}
	}

	/**
	 *	First test; get all webhooks
	 */
	tempFunc(
		consts.COVIDNOTIFICATIONS,
		`[{"webhook_id":"JLF7HyWtBCU7L4jU7M4L","url":"https://webhook.site/7b87be25-4e67-46e3-9f70-096ed2eaf1b6","country":"Norway","calls":1},{"webhook_id":"OuyRmaYnAXFbpZrILY3F","url":"https://webhook.site/262a030c-458b-4b8c-b8cf-5a848165ef22","country":"Sweden","calls":2},{"webhook_id":"UfKDHNGbaDkj5qHZdwBa","url":"https://webhook.site/262a030c-458b-4b8c-b8cf-5a848165ef22","country":"Denmark","calls":5},{"webhook_id":"fhH6xuu9LAE4MJEGTBoE","url":"https://webhook.site/53d624f0-3f1f-4899-8cac-62c02014176b","country":"Spain","calls":2}]`,
	)

	/**
	 *	Second test; get a specific webhook
	 */
	tempFunc(
		consts.COVIDNOTIFICATIONS+"JLF7HyWtBCU7L4jU7M4L",
		`{"webhook_id":"JLF7HyWtBCU7L4jU7M4L","url":"https://webhook.site/7b87be25-4e67-46e3-9f70-096ed2eaf1b6","country":"Norway","calls":1}`,
	)

	// Revert mocking
	notifications.Webhooks = Webhooks_original
	notifications.NotificationGetRequest = NotificationGetRequest_original
}
