// domainr is a wrapper for the domai.nr API written in Go.
//
//  // Get the url to register a domain
//  domainr.Register("myfancydomain.com", "")
//  // Get the url of a specific registrar to register a domain
//  domainr.Register("myfancydomain.com", "gandi.net")
//  // Get information about a domain
//  domainr.Json(domainr.METHOD_INFO, "github.io", "")
//  // Get information about the domains matching your query
//  domainr.Json(domainr.METHOD_SEARCH, "myfancydomain", "")
//  // The optional third parameter at domainr.Json is a callback as defined in
//  // the API documentation https://domai.nr/api/docs/json
package domainr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URI      = "https://domai.nr/api"
	METHOD_SEARCH = "search"
	METHOD_INFO   = "info"
)

func Register(domain, registrar string) (interface{}, error) {
	params := "domain=" + domain
	if registrar != "" {
		params += "&registrar=" + registrar
	}
	return apiCall("register", params)
}

func Json(method, query, callback string) (interface{}, error) {
	if method == METHOD_INFO || method == METHOD_SEARCH {
		params := "q=" + query
		if callback != "" {
			params += "&callback=" + callback
		}
		return apiCall("json/"+method, params)
	} else {
		return nil, errors.New("Invalid method provided.")
	}
}

func apiCall(apiMethod, params string) (interface{}, error) {
	callUrl := BASE_URI + "/" + apiMethod + "?"
	callUrl += params

	resp, err := http.Get(callUrl)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.Request.URL.String() != callUrl {
		return resp.Request.URL.String(), nil
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsed interface{}
	err = json.Unmarshal(response, &parsed)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}
