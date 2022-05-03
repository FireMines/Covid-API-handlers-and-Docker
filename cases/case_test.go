package cases_test

import (
	"covidAss2/cases"
	outsideapi "covidAss2/outsideApi"
	consts "covidAss2/variables"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
 *	Tester for the HandlerCases function.
 *	Mocks any outgoing request during testing.
 */
func TestHandlerCases(t *testing.T) {
	// Mock outgoing functions so they return hardcoded results
	getGraphqlResponse_original := outsideapi.GetGraphqlResponse
	outsideapi.GetGraphqlResponse = func(graphqlBody, url string) (map[string]interface{}, error) {
		graphqlResponseMocked, err := outsideapi.ReadJSONToken("../test.json")
		if err != nil {
			return map[string]interface{}{}, err
		}

		return graphqlResponseMocked, nil
	}

	testurl := consts.COVIDCASES + "Norway"

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest(http.MethodGet, testurl, nil)
	req.URL.Path = consts.COVIDCASES + "Norway"
	if err != nil {
		t.Error("Error when creating new request", err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	cases.CovidInfoHandler(rr, req)
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

	expected := `{"name":"Norway","date":"2022-05-01","confirmed":1426128,"recovered":2932,"deaths":0,"growthRate":0.00008064442610270734}`

	// Check if expected result was the same as the actualy response
	data_str := string(data)
	if data_str != expected {
		t.Errorf("Handler did not return expected value: got '%v' want '%v'", data_str, expected)
	}

	// Revert mocking
	outsideapi.GetGraphqlResponse = getGraphqlResponse_original
}
