package outsideapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/machinebox/graphql"
)

var GetGraphqlResponse = GetGraphqlResponsef
var GetHttpResponse = GetHttpResponsef

/**
 *	Same code as before; gets the graphql response
 *
 *	@param graphqlBody - The body of the graphql request, as a string.
 *	@param url - The url of the graphql API.
 *
 *	@return A map containing the response from the GraphQl API, and an error if something went wrong.
 */
func GetGraphqlResponsef(graphqlBody string, url string) (map[string]interface{}, error) {
	graphqlClient := graphql.NewClient(url)
	graphqlRequest := graphql.NewRequest(graphqlBody)

	var graphqlResponse map[string]interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		return map[string]interface{}{}, err
	}

	/* 	err := SaveToJSONFile("test.json", graphqlResponse)
	   	if err != nil {
	   		return graphqlResponse, err
	   	} */

	return graphqlResponse, nil
}

/**
 *	Gets a http respons
 *
 *	@param url - The url of the website
 *
 *	@return	A map containing the response from the API, and an error if something went wrong.
 */
func GetHttpResponsef(url string) (map[string]interface{}, error) {
	// Gets covid tracker and country from url
	resp, err := http.Get(url)
	if err != nil {
		return map[string]interface{}{}, err
	}

	defer resp.Body.Close()

	// Decoding
	dataRaw := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&dataRaw)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return dataRaw, nil
}

/**
 *	Reads cached data from json file
 *
 *	@param fileName - The filename from which we get the json data from
 *
 *	@return	A map containing the data stored in the cache, and an error if something went wrong.
 */
func ReadJSONToken(fileName string) (map[string]interface{}, error) {
	// Read the mock file's contents
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var filteredData map[string]interface{}
	json.Unmarshal([]byte(content), &filteredData)

	/* 	err = json.NewDecoder(file).Decode(&filteredData)
	   	if err != nil {
	   		return map[string]interface{}{}, err
	   	} */

	fmt.Println("filtered data:", filteredData)
	return filteredData, nil
}

/**
 *	Saves cached data to json file
 *
 *	@param fileName - The filename from which we save the json data to
 *
 *	@return	Error if something went wrong.
 */
func SaveToJSONFile(filename string, graphqlResponse map[string]interface{}) error {
	file, err := os.OpenFile(filename, os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(graphqlResponse)
	if err != nil {
		return err
	}
	return nil
}
